package application

import (
	"context"
	"github.com/example/geospatial-tile-build-service/internal/projection/domain"
)

func (s *Service) ConvertMany(ctx context.Context, points []domain.Point, from, to string) ([]domain.Point, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	out := make([]domain.Point, len(points))
	for i, p := range points {
		// Honor cancellation between conversions: a canceled batch must not
		// keep grinding through the whole point set.
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		v, err := s.Convert(ctx, p, from, to)
		if err != nil {
			return nil, err
		}
		out[i] = v
	}
	return out, nil
}
func Bounds(points []domain.Point) domain.Point {
	if len(points) == 0 {
		return domain.Point{}
	}
	var x, y float64
	for _, p := range points {
		x += p.X
		y += p.Y
	}
	return domain.Point{X: x / float64(len(points)), Y: y / float64(len(points))}
}
