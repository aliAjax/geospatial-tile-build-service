package domain

import (
	"errors"
	"sort"
	"strings"
)

type Options struct {
	MinZoom, MaxZoom int
	Layers           []string
	Extent           uint32
	Buffer           uint32
	Simplify         float64
}

func (o Options) Validate() error {
	if o.MinZoom < 0 || o.MaxZoom < o.MinZoom || o.MaxZoom > 22 {
		return errors.New("invalid zoom range")
	}
	if o.Extent == 0 {
		o.Extent = 4096
	}
	if o.Extent > 16384 {
		return errors.New("extent too large")
	}
	if len(o.Layers) == 0 {
		return errors.New("at least one layer required")
	}
	return nil
}
func (o Options) Canonical() string {
	// Sort a copy, not the caller's slice. sort.Strings mutates the backing
	// array in place; running Canonical concurrently on the same Options (or
	// even once) clobbered the caller's Layers order, which is both a data race
	// and a surprising mutation of an input value.
	layers := make([]string, len(o.Layers))
	copy(layers, o.Layers)
	sort.Strings(layers)
	return strings.Join(layers, ",") + ":" + itoa(o.MinZoom) + ":" + itoa(o.MaxZoom)
}
func itoa(v int) string {
	if v == 0 {
		return "0"
	}
	s := ""
	for v > 0 {
		s = string(rune('0'+v%10)) + s
		v /= 10
	}
	return s
}
