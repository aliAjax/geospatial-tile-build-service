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

func (f Feature) Clone() Feature {
	c := f
	if f.Properties != nil {
		c.Properties = make(map[string]any, len(f.Properties))
		for k, v := range f.Properties {
			c.Properties[k] = v
		}
	}
	switch g := f.Geometry.(type) {
	case LineString:
		ng := make(LineString, len(g))
		copy(ng, g)
		c.Geometry = ng
	case Polygon:
		np := Polygon{Rings: make([][]Point, len(g.Rings))}
		for i, r := range g.Rings {
			nr := make([]Point, len(r))
			copy(nr, r)
			np.Rings[i] = nr
		}
		c.Geometry = np
	}
	return c
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
