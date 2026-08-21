package domain

import (
	"fmt"
	"math"
)

type Matrix struct {
	Zoom int
	Size int
}

func NewMatrix(z int) Matrix {
	if z < 0 {
		z = 0
	}
	if z > 22 {
		z = 22
	}
	return Matrix{Zoom: z, Size: 1 << z}
}
func (m Matrix) Valid() bool             { return m.Size == int(math.Exp2(float64(m.Zoom))) }

func ValidateMatrix(m *Matrix) error { if !m.Valid() { return fmt.Errorf("invalid tile matrix") }; return nil }
func (m Matrix) Key(t Coordinate) string { return fmt.Sprintf("z%d-x%d-y%d", t.Z, t.X, t.Y) }
func (m Matrix) Children(t Coordinate) []Coordinate {
	return []Coordinate{{t.Z + 1, t.X * 2, t.Y * 2}, {t.Z + 1, t.X*2 + 1, t.Y * 2}, {t.Z + 1, t.X * 2, t.Y*2 + 1}, {t.Z + 1, t.X*2 + 1, t.Y*2 + 1}}
}
