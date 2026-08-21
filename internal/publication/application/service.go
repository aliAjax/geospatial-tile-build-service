package application

import (
	"context"
	"fmt"
	"github.com/example/geospatial-tile-build-service/internal/publication/domain"
	"sync"
	"time"
)

type Service struct {
	mu       sync.RWMutex
	releases map[string]domain.Release
}

func New() *Service { return &Service{releases: map[string]domain.Release{}} }
func releaseKey(dataset, channel string) string { return dataset + ":" + channel }
func (s *Service) Publish(ctx context.Context, dataset, version, channel string) (domain.Release, error) {
	select {
	case <-ctx.Done():
		return domain.Release{}, ctx.Err()
	default:
	}
	if dataset == "" || version == "" || channel == "" {
		return domain.Release{}, fmt.Errorf("dataset, version and channel required")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	key := releaseKey(dataset, channel)
	if prev, ok := s.releases[key]; ok && prev.Active {
		prev = prev.Retire()
		prev.UpdatedAt = time.Now().UTC()
		s.releases[key] = prev
	}
	r := domain.Release{DatasetID: dataset, VersionID: version, Channel: channel, UpdatedAt: time.Now().UTC()}.Activate()
	s.releases[key] = r
	return r, nil
}
func (s *Service) Current(_ context.Context, dataset, channel string) (domain.Release, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	r, ok := s.releases[releaseKey(dataset, channel)]
	if !ok {
		return r, fmt.Errorf("release not found")
	}
	return r, nil
}
