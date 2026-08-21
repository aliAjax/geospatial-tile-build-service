package application

import (
	"context"
	"fmt"
	storagedomain "github.com/example/geospatial-tile-build-service/internal/storage/domain"
	tilingdomain "github.com/example/geospatial-tile-build-service/internal/tiling/domain"
	"io"
)

type Service struct{ store storagedomain.Store }

func New(s storagedomain.Store) *Service { return &Service{store: s} }
func (s *Service) Put(ctx context.Context, t tilingdomain.Coordinate, b []byte) error {
	if err := t.Validate(); err != nil {
		return fmt.Errorf("tile coordinate: %w", err)
	}
	return s.store.Put(ctx, "tiles/"+t.Path()+".pbf", b)
}
func (s *Service) Open(ctx context.Context, t tilingdomain.Coordinate) (io.ReadCloser, error) {
	if err := t.Validate(); err != nil {
		return nil, err
	}
	return s.store.Open(ctx, "tiles/"+t.Path()+".pbf")
}
