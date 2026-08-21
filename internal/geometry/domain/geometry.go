package domain

import (
	"errors"
	"math"
)

type Point struct{ X, Y float64 }
type LineString []Point
type Polygon struct{ Rings [][]Point }
type Feature struct {
	ID         string
	Geometry   any
	Properties map[string]any
}

func ValidPoint(p Point) bool {
	return !math.IsNaN(p.X) && !math.IsNaN(p.Y) && !math.IsInf(p.X, 0) && !math.IsInf(p.Y, 0) && p.X >= -180 && p.X <= 180 && p.Y >= -90 && p.Y <= 90
}
func ValidateLine(l LineString) error {
	if len(l) < 2 {
		return errors.New("line requires two points")
	}
	for _, p := range l {
		if !ValidPoint(p) {
			return errors.New("invalid line point")
		}
	}
	return nil
}
func ValidatePolygon(p Polygon) error {
	if len(p.Rings) == 0 {
		return errors.New("polygon has no rings")
	}
	for _, r := range p.Rings {
		if len(r) < 4 {
			return errors.New("ring requires four points")
		}
		for _, pt := range r {
			if !ValidPoint(pt) {
				return errors.New("invalid polygon point")
			}
		}
	}
	return nil
}
