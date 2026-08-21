package application

import (
	"context"
	"github.com/example/geospatial-tile-build-service/internal/geometry/domain"
	"testing"
)

func TestFilterDoesNotMutateOriginalFeatures(t *testing.T) {
	f := NewFilter()
	original := []domain.Feature{{ID: "7", Properties: map[string]any{"name": "road", "class": "primary"}}}
	out := f.Select(context.Background(), original, map[string]struct{}{"name": {}})
	out[0].Properties["name"] = "changed"
	if original[0].Properties["name"] != "road" {
		t.Fatalf("original feature was aliased: %#v", original[0].Properties)
	}
}
