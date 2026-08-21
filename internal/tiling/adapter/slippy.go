package adapter

import (
	"github.com/example/geospatial-tile-build-service/internal/tiling/domain"
	"math"
)

func Bounds(t domain.Coordinate) (minLon, minLat, maxLon, maxLat float64) {
	n := math.Exp2(float64(t.Z))
	minLon = float64(t.X)/n*360 - 180
	maxLon = float64(t.X+1)/n*360 - 180
	maxLat = math.Atan(math.Sinh(math.Pi*(1-2*float64(t.Y)/n))) * 180 / math.Pi
	minLat = math.Atan(math.Sinh(math.Pi*(1-2*float64(t.Y+1)/n))) * 180 / math.Pi
	return
}
