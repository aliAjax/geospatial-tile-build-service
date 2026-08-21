package application

import (
	"github.com/example/geospatial-tile-build-service/internal/projection/domain"
	"testing"
)

func TestClonePointsOwnsBackingArray(t *testing.T) {
	src := []domain.Point{{X: 1, Y: 1}}
	got := ClonePoints(src)
	got[0].X = 9
	if src[0].X != 1 {
		t.Fatal("point slice aliases source")
	}
}
