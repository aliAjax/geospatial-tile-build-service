package domain

import "math"

type Transform struct{ ScaleX, ScaleY, OffsetX, OffsetY float64 }

func NewTransform(src, dst Point) Transform {
	return Transform{ScaleX: 1, ScaleY: 1, OffsetX: dst.X - src.X, OffsetY: dst.Y - src.Y}
}
func (t Transform) Apply(p Point) Point {
	return Point{X: p.X*t.ScaleX + t.OffsetX, Y: p.Y*t.ScaleY + t.OffsetY}
}
func (t Transform) Inverse(p Point) Point {
	return Point{X: (p.X - t.OffsetX) / t.ScaleX, Y: (p.Y - t.OffsetY) / t.ScaleY}
}
func NormalizeLon(lon float64) float64 {
	for lon > 180 {
		lon -= 360
	}
	for lon < -180 {
		lon += 360
	}
	return lon
}
func NormalizeLat(lat float64) float64 { return math.Max(-90, math.Min(90, lat)) }
