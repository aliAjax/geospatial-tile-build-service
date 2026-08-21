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
	for k, r := range s.releases {
		if r.DatasetID == dataset && r.Channel == channel {
			r.Active = false
			s.releases[k] = r
		}
	}
	r := domain.Release{DatasetID: dataset, VersionID: version, Channel: channel, Active: true, UpdatedAt: time.Now().UTC()}
	s.releases[dataset+":"+channel] = r
	return r, nil
}
func (s *Service) Current(_ context.Context, dataset, channel string) (domain.Release, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	r, ok := s.releases[dataset+":"+channel]
	if !ok {
		return r, fmt.Errorf("release not found")
	}
	return r, nil
}
