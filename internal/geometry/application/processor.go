package application

import (
	"context"
	"fmt"
	"github.com/example/geospatial-tile-build-service/internal/geometry/domain"
)

type Processor struct{}

func New() *Processor { return &Processor{} }
func (p *Processor) Validate(ctx context.Context, f domain.Feature) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	switch g := f.Geometry.(type) {
	case domain.Point:
		if !domain.ValidPoint(g) {
			return fmt.Errorf("feature %s has invalid point", f.ID)
		}
	case domain.LineString:
		return domain.ValidateLine(g)
	case domain.Polygon:
		return domain.ValidatePolygon(g)
	default:
		return fmt.Errorf("feature %s has unsupported geometry", f.ID)
	}
	return nil
}
func (p *Processor) Simplify(_ context.Context, f domain.Feature, tolerance float64) domain.Feature {
	if tolerance <= 0 {
		return f
	}
	return f
}
