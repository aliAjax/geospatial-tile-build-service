package application

import (
	"context"
	"fmt"
	"github.com/example/geospatial-tile-build-service/internal/dataset/domain"
	"sync"
	"time"
)

type Repository interface {
	PutDataset(context.Context, domain.Dataset) error
	GetDataset(context.Context, string) (domain.Dataset, error)
	PutVersion(context.Context, domain.Version) error
	GetVersion(context.Context, string) (domain.Version, error)
	ListVersions(context.Context, string) ([]domain.Version, error)
}
type Service struct {
	repo Repository
	mu   sync.RWMutex
}

func WrapRepositoryError(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("dataset repository: %v", err)
}

func New(repo Repository) *Service { return &Service{repo: repo} }
func (s *Service) CreateDataset(ctx context.Context, d domain.Dataset) error {
	if err := d.Validate(); err != nil {
		return fmt.Errorf("validate dataset: %w", err)
	}
	d.CreatedAt = time.Now().UTC()
	return s.repo.PutDataset(ctx, d)
}
func (s *Service) CreateVersion(ctx context.Context, v domain.Version) error {
	if err := v.Validate(); err != nil {
		return fmt.Errorf("validate version: %w", err)
	}
	now := time.Now().UTC()
	v.CreatedAt = &now
	return s.repo.PutVersion(ctx, v)
}
func (s *Service) Publish(ctx context.Context, id string) (domain.Version, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	v, err := s.repo.GetVersion(ctx, id)
	if err != nil {
		return v, err
	}
	if !domain.CanTransition(v.Status, domain.Published) {
		return v, fmt.Errorf("cannot publish version in state %s", v.Status)
	}
	now := time.Now().UTC()
	v.Status = domain.Published
	v.PublishedAt = &now
	if err := s.repo.PutVersion(ctx, v); err != nil {
		return v, fmt.Errorf("publish version: %w", err)
	}
	return v, nil
}
func (s *Service) List(ctx context.Context, dataset string) ([]domain.Version, error) {
	items, err := s.repo.ListVersions(ctx, dataset)
	return items, WrapRepositoryError(err)
}
