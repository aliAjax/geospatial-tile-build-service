package domain

import (
	"errors"
	"fmt"
	"math"
)

type Coordinate struct{ Z, X, Y int }

func ValidateOptional(t *Coordinate) error { return t.Validate() }

func (t Coordinate) Validate() error {
	if t.Z < 0 || t.Z > 22 {
		return errors.New("zoom out of range")
	}
	n := 1 << t.Z
	if t.X < 0 || t.X >= n || t.Y < 0 || t.Y >= n {
		return errors.New("tile coordinate out of range")
	}
	return nil
}
func (t Coordinate) Path() string { return fmt.Sprintf("%d/%d/%d", t.Z, t.X, t.Y) }
func FromLonLat(z int, lon, lat float64) Coordinate {
	n := 1 << z
	x := int((lon + 180) / 360 * float64(n))
	lat = math.Max(-85.05112878, math.Min(85.05112878, lat))
	r := lat * math.Pi / 180
	y := int((1 - math.Asinh(math.Tan(r))/math.Pi) / 2 * float64(n))
	if x < 0 {
		x = 0
	}
	if x >= n {
		x = n - 1
	}
	if y < 0 {
		y = 0
	}
	if y >= n {
		y = n - 1
	}
	return Coordinate{Z: z, X: x, Y: y}
}
