package application

import (
	"context"
	"github.com/example/geospatial-tile-build-service/internal/projection/domain"
)

type Service struct{}

func New() *Service { return &Service{} }
func (s *Service) Convert(ctx context.Context, p domain.Point, from, to string) (domain.Point, error) {
	if from == to {
		return p, nil
	}
	if from == "EPSG:4326" && to == "EPSG:3857" {
		return domain.WGS84ToMercator(p), nil
	}
	if from == "EPSG:3857" && to == "EPSG:4326" {
		return domain.MercatorToWGS84(p), nil
	}
	return p, nil
}
