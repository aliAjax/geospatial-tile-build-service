package adapter

import (
	"encoding/binary"
	"fmt"
	"github.com/example/geospatial-tile-build-service/internal/geometry/domain"
	"math"
)

func EncodePoint(p domain.Point) []byte {
	b := make([]byte, 17)
	b[0] = 1
	binary.LittleEndian.PutUint32(b[1:], 1)
	binary.LittleEndian.PutUint64(b[5:], mathbits(p.X))
	binary.LittleEndian.PutUint64(b[13:], mathbits(p.Y))
	return b
}
func DecodePoint(b []byte) (domain.Point, error) {
	if len(b) < 17 {
		return domain.Point{}, fmt.Errorf("point binary too short")
	}
	return domain.Point{X: unmathbits(binary.LittleEndian.Uint64(b[5:])), Y: unmathbits(binary.LittleEndian.Uint64(b[13:]))}, nil
}
func mathbits(v float64) uint64   { return math.Float64bits(v) }
func unmathbits(v uint64) float64 { return math.Float64frombits(v) }
