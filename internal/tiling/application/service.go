package application

import (
	"context"
	"github.com/example/geospatial-tile-build-service/internal/tiling/domain"
)

type Service struct{}

func New() *Service { return &Service{} }
func (s *Service) Validate(ctx context.Context, t domain.Coordinate) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	return t.Validate()
}
func (s *Service) Enumerate(ctx context.Context, z int) ([]domain.Coordinate, error) {
	out := []domain.Coordinate{}
	n := 1 << z
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
			}
			out = append(out, domain.Coordinate{Z: z, X: x, Y: y})
		}
	}
	return out, nil
}
