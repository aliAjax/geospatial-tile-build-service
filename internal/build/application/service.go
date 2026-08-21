package application

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/example/geospatial-tile-build-service/internal/build/domain"
	storage "github.com/example/geospatial-tile-build-service/internal/storage/domain"
	"sync"
	"time"
)

type Service struct {
	store storage.Store
	mu    sync.RWMutex
	jobs  map[string]domain.Job
}

func New(s storage.Store) *Service { return &Service{store: s, jobs: map[string]domain.Job{}} }
func (s *Service) Start(ctx context.Context, id, version string, input []byte) (domain.Job, error) {
	j := domain.Job{ID: id, VersionID: version, Status: domain.BuildRunning, InputHash: hash(input)}
	if err := j.Validate(); err != nil {
		return j, err
	}
	now := time.Now().UTC()
	j.StartedAt = &now
	s.mu.Lock()
	s.jobs[id] = j
	s.mu.Unlock()
	if err := s.store.Put(ctx, "builds/"+id+"/input", input); err != nil {
		return s.fail(id, err)
	}
	out := sha256.Sum256(append([]byte("compiled:"), input...))
	j.OutputHash = hex.EncodeToString(out[:])
	j.Status = domain.BuildDone
	done := time.Now().UTC()
	j.FinishedAt = &done
	s.mu.Lock()
	s.jobs[id] = j
	s.mu.Unlock()
	return j, nil
}
func (s *Service) fail(id string, err error) (domain.Job, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	j := s.jobs[id]
	j.Status = domain.BuildFailed
	j.Error = err.Error()
	s.jobs[id] = j
	return j, fmt.Errorf("build: %w", err)
}
func (s *Service) Status(_ context.Context, id string) (domain.Job, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	j, ok := s.jobs[id]
	if !ok {
		return j, fmt.Errorf("build %s not found", id)
	}
	return j, nil
}
func hash(b []byte) string { v := sha256.Sum256(b); return hex.EncodeToString(v[:]) }
