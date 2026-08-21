package application

import (
	"context"
	"github.com/example/geospatial-tile-build-service/internal/geometry/domain"
	"strings"
)

type Filter struct{}

func NewFilter() *Filter { return &Filter{} }
func (f *Filter) Select(ctx context.Context, features []domain.Feature, allowed map[string]struct{}) []domain.Feature {
	out := make([]domain.Feature, 0, len(features))
	for _, v := range features {
		select {
		case <-ctx.Done():
			return out
		default:
		}
		if len(allowed) == 0 {
			out = append(out, v.Clone())
			continue
		}
		clone := v.Clone()
		for k := range clone.Properties {
			if _, ok := allowed[k]; !ok {
				delete(clone.Properties, k)
			}
		}
		out = append(out, clone)
	}
	return out
}
func (f *Filter) NormalizeProperties(features []domain.Feature) []domain.Feature {
	out := make([]domain.Feature, len(features))
	for i, v := range features {
		clone := v.Clone()
		src := v.Properties
		p := make(map[string]any, len(src))
		for k, val := range src {
			p[strings.TrimSpace(strings.ToLower(k))] = val
		}
		clone.Properties = p
		out[i] = clone
	}
	return out
}
