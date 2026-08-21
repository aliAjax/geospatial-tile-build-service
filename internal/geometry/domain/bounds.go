package domain

import "math"

type Bounds struct{ MinX, MinY, MaxX, MaxY float64 }

func NewBounds() Bounds {
	return Bounds{MinX: math.Inf(1), MinY: math.Inf(1), MaxX: math.Inf(-1), MaxY: math.Inf(-1)}
}
func (b *Bounds) Add(p Point) {
	if p.X < b.MinX {
		b.MinX = p.X
	}
	if p.Y < b.MinY {
		b.MinY = p.Y
	}
	if p.X > b.MaxX {
		b.MaxX = p.X
	}
	if p.Y > b.MaxY {
		b.MaxY = p.Y
	}
}
func (b Bounds) Valid() bool     { return b.MinX <= b.MaxX && b.MinY <= b.MaxY }
func (b Bounds) Width() float64  { return b.MaxX - b.MinX }
func (b Bounds) Height() float64 { return b.MaxY - b.MinY }
func (b Bounds) Contains(p Point) bool {
	return p.X >= b.MinX && p.X <= b.MaxX && p.Y >= b.MinY && p.Y <= b.MaxY
}
func (b Bounds) Intersects(o Bounds) bool {
	return b.MinX <= o.MaxX && b.MaxX >= o.MinX && b.MinY <= o.MaxY && b.MaxY >= o.MinY
}
func (b Bounds) Expand(delta float64) Bounds {
	return Bounds{b.MinX - delta, b.MinY - delta, b.MaxX + delta, b.MaxY + delta}
}
func (b Bounds) ClampWorld() Bounds {
	if b.MinX < -180 {
		b.MinX = -180
	}
	if b.MaxX > 180 {
		b.MaxX = 180
	}
	if b.MinY < -90 {
		b.MinY = -90
	}
	if b.MaxY > 90 {
		b.MaxY = 90
	}
	return b
}
