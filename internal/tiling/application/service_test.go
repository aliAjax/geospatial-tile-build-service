package application

import (
	"context"
	"github.com/example/geospatial-tile-build-service/internal/tiling/adapter"
	"github.com/example/geospatial-tile-build-service/internal/tiling/domain"
	"testing"
)

func TestDefaultMatrixEnumerationIsSafe(t *testing.T) {
	s := New()
	got, err := s.Enumerate(context.Background(), 0)
	if err != nil || len(got) != 1 || got[0].Z != 0 || got[0].X != 0 || got[0].Y != 0 {
		t.Fatalf("default matrix enumeration invalid: %#v %v", got, err)
	}
	m := domain.NewMatrix(-1)
	if !m.Valid() {
		t.Fatalf("negative zoom did not normalize: %+v", m)
	}
	minLon, minLat, maxLon, maxLat := adapter.Bounds(got[0])
	if minLon >= maxLon || minLat >= maxLat {
		t.Fatalf("invalid world bounds: %v %v %v %v", minLon, minLat, maxLon, maxLat)
	}
}
