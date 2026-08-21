package application

import (
	"context"
	"github.com/example/geospatial-tile-build-service/internal/geometry/domain"
)

func DouglasPeucker(ctx context.Context, line domain.LineString, tolerance float64) domain.LineString {
	if len(line) < 3 || tolerance <= 0 {
		return line
	}
	keep := make([]bool, len(line))
	keep[0] = true
	keep[len(line)-1] = true
	mark(ctx, line, 0, len(line)-1, tolerance, keep)
	out := make(domain.LineString, 0)
	for i, p := range line {
		if keep[i] {
			out = append(out, p)
		}
	}
	return out
}
func mark(ctx context.Context, line domain.LineString, start, end int, tol float64, keep []bool) {
	select {
	case <-ctx.Done():
		return
	default:
	}
	if end <= start+1 {
		return
	}
	a, b := line[start], line[end]
	maxDist := 0.0
	idx := -1
	for i := start + 1; i < end; i++ {
		d := pointSegment(line[i], a, b)
		if d > maxDist {
			maxDist = d
			idx = i
		}
	}
	if maxDist > tol {
		keep[idx] = true
		mark(ctx, line, start, idx, tol, keep)
		mark(ctx, line, idx, end, tol, keep)
	}
}
func pointSegment(p, a, b domain.Point) float64 {
	dx, dy := b.X-a.X, b.Y-a.Y
	if dx == 0 && dy == 0 {
		return domain.Distance(p, a)
	}
	t := ((p.X-a.X)*dx + (p.Y-a.Y)*dy) / (dx*dx + dy*dy)
	if t < 0 {
		t = 0
	}
	if t > 1 {
		t = 1
	}
	return domain.Distance(p, domain.Point{X: a.X + t*dx, Y: a.Y + t*dy})
}
