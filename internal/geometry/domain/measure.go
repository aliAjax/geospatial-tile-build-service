package domain

import (
	"math"
	"sort"
)

func Distance(a, b Point) float64 { return math.Hypot(a.X-b.X, a.Y-b.Y) }
func LineLength(l LineString) float64 {
	var n float64
	for i := 1; i < len(l); i++ {
		n += Distance(l[i-1], l[i])
	}
	return n
}
func RingArea(r []Point) float64 {
	var a float64
	for i := range r {
		j := (i + 1) % len(r)
		a += r[i].X*r[j].Y - r[j].X*r[i].Y
	}
	return a / 2
}
func PolygonArea(p Polygon) float64 {
	if len(p.Rings) == 0 {
		return 0
	}
	a := math.Abs(RingArea(p.Rings[0]))
	for _, r := range p.Rings[1:] {
		a -= math.Abs(RingArea(r))
	}
	if a < 0 {
		return 0
	}
	return a
}
func Centroid(points []Point) Point {
	if len(points) == 0 {
		return Point{}
	}
	xs := make([]float64, len(points))
	ys := make([]float64, len(points))
	for i, p := range points {
		xs[i] = p.X
		ys[i] = p.Y
	}
	sort.Float64s(xs)
	sort.Float64s(ys)
	return Point{X: xs[len(xs)/2], Y: ys[len(ys)/2]}
}
func Snap(p Point, grid float64) Point {
	if grid <= 0 {
		return p
	}
	return Point{X: math.Round(p.X/grid) * grid, Y: math.Round(p.Y/grid) * grid}
}
