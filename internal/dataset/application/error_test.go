package application

import (
	"context"
	"errors"
	"github.com/example/geospatial-tile-build-service/internal/dataset/domain"
	"testing"
)

type failingRepo struct{}

func (failingRepo) PutDataset(context.Context, domain.Dataset) error { return nil }
func (failingRepo) GetDataset(context.Context, string) (domain.Dataset, error) {
	return domain.Dataset{}, nil
}
func (failingRepo) PutVersion(context.Context, domain.Version) error { return nil }
func (failingRepo) GetVersion(context.Context, string) (domain.Version, error) {
	return domain.Version{}, nil
}
func (failingRepo) ListVersions(context.Context, string) ([]domain.Version, error) {
	return nil, domain.ErrInvalidLayer
}

func TestWrapRepositoryErrorPreservesCause(t *testing.T) {
	if !errors.Is(WrapRepositoryError(domain.ErrInvalidLayer), domain.ErrInvalidLayer) {
		t.Fatal("repository error cause was lost")
	}
	if _, err := New(failingRepo{}).List(context.Background(), "roads"); !errors.Is(err, domain.ErrInvalidLayer) {
		t.Fatalf("service error cause was lost: %v", err)
	}
}
