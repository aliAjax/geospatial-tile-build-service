package application

import (
	"context"
	"github.com/example/geospatial-tile-build-service/internal/geometry/domain"
	"testing"
)

func TestNormalizePropertiesOwnsResult(t *testing.T) {
	src := []domain.Feature{{Properties: map[string]any{" NAME ": "road"}}}
	got := NewFilter().NormalizeProperties(src)
	got[0].Properties["name"] = "x"
	if _, ok := src[0].Properties[" NAME "]; !ok {
		t.Fatal("normalization mutated source")
	}
}
func TestSimplifyOwnsFeature(t *testing.T) {
	src := domain.Feature{Properties: map[string]any{"name": "road"}}
	got := New().Simplify(context.Background(), src, 1)
	got.Properties["name"] = "x"
	if src.Properties["name"] != "road" {
		t.Fatal("simplify aliases source")
	}
}
