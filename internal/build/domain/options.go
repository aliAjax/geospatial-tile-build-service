package domain

import (
	"errors"
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
	layers := append([]string(nil), o.Layers...)
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
