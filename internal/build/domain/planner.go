package domain

import (
	"fmt"
	"sort"
	"time"
)

type TilePlan struct {
	Zoom           int
	Tiles          uint64
	EstimatedBytes uint64
	CreatedAt      time.Time
}

func Plan(o Options) []TilePlan {
	out := make([]TilePlan, 0, o.MaxZoom-o.MinZoom+1)
	for z := o.MinZoom; z <= o.MaxZoom; z++ {
		count := uint64(1) << (2 * z)
		out = append(out, TilePlan{Zoom: z, Tiles: count, EstimatedBytes: count * 256, CreatedAt: time.Now().UTC()})
	}
	return out
}
func SortPlans(items []TilePlan) []TilePlan {
	out := append([]TilePlan(nil), items...)
	sort.Slice(out, func(i, j int) bool { return out[i].Zoom < out[j].Zoom })
	return out
}
func Estimate(o Options) uint64 {
	var n uint64
	for _, p := range Plan(o) {
		n += p.EstimatedBytes
	}
	return n
}
func ValidatePlan(items []TilePlan) error {
	if len(items) == 0 {
		return fmt.Errorf("empty tile plan")
	}
	for i, p := range items {
		if p.Zoom < 0 || p.Tiles == 0 {
			return fmt.Errorf("invalid plan at %d", i)
		}
	}
	return nil
}
func Chunk(items []TilePlan, size int) [][]TilePlan {
	if size < 1 {
		size = 1
	}
	out := [][]TilePlan{}
	for len(items) > 0 {
		n := size
		if n > len(items) {
			n = len(items)
		}
		out = append(out, append([]TilePlan(nil), items[:n]...))
		items = items[n:]
	}
	return out
}
