package domain

import "math"

type Point struct{ X, Y float64 }

func WGS84ToMercator(p Point) Point {
	x := p.X * 20037508.34 / 180
	y := math.Log(math.Tan((90+p.Y)*math.Pi/360)) / (math.Pi / 180)
	y = y * 20037508.34 / 180
	return Point{X: x, Y: y}
}
func MercatorToWGS84(p Point) Point {
	x := p.X / 20037508.34 * 180
	y := p.Y / 20037508.34 * 180
	y = 180 / math.Pi * (2*math.Atan(math.Exp(y*math.Pi/180)) - math.Pi/2)
	return Point{X: x, Y: y}
}
