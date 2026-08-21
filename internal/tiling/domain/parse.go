package domain

import (
	"fmt"
	"strconv"
	"strings"
)

func ParsePath(path string) (Coordinate, error) {
	p := strings.Split(strings.Trim(path, "/"), "/")
	if len(p) != 3 {
		return Coordinate{}, fmt.Errorf("tile path must have z/x/y")
	}
	z, e1 := strconv.Atoi(p[0])
	x, e2 := strconv.Atoi(p[1])
	y, e3 := strconv.Atoi(strings.TrimSuffix(p[2], ".pbf"))
	if e1 != nil || e2 != nil || e3 != nil {
		return Coordinate{}, fmt.Errorf("tile path contains non-numeric coordinate")
	}
	t := Coordinate{Z: z, X: x, Y: y}
	if err := t.Validate(); err != nil {
		return t, err
	}
	return t, nil
}
func Parent(t Coordinate) Coordinate {
	if t.Z == 0 {
		return t
	}
	return Coordinate{Z: t.Z - 1, X: t.X / 2, Y: t.Y / 2}
}
func Siblings(t Coordinate) []Coordinate {
	if t.Z == 0 {
		return []Coordinate{t}
	}
	p := Parent(t)
	out := make([]Coordinate, 0, 4)
	for dy := 0; dy < 2; dy++ {
		for dx := 0; dx < 2; dx++ {
			out = append(out, Coordinate{Z: t.Z, X: p.X*2 + dx, Y: p.Y*2 + dy})
		}
	}
	return out
}
