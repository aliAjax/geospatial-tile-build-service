package domain

import "math"

func Quantize(v, min, max float64, extent uint32) uint32 {
	if max <= min || v <= min {
		return 0
	}
	if v >= max {
		return extent
	}
	return uint32(math.Round((v - min) / (max - min) * float64(extent)))
}
func QuantizePoint(x, y, minX, minY, maxX, maxY float64, extent uint32) (uint32, uint32) {
	return Quantize(x, minX, maxX, extent), Quantize(y, minY, maxY, extent)
}
func SignedCommand(cmd uint32, delta int32) uint32 { return cmd | uint32((delta<<1)^(delta>>31))<<3 }
func MoveTo(x, y int32) []uint32                   { return []uint32{SignedCommand(1, x), SignedCommand(1, y)} }
func LineTo(x, y int32) []uint32                   { return []uint32{SignedCommand(2, x), SignedCommand(2, y)} }
