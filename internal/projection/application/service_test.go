package application

import (
	"context"
	"github.com/example/geospatial-tile-build-service/internal/projection/domain"
	"testing"
)

func TestConvertManyHonorsCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := New().ConvertMany(ctx, []domain.Point{{X: 1, Y: 1}}, "EPSG:4326", "EPSG:3857"); err == nil {
		t.Fatal("canceled batch accepted")
	}
}
