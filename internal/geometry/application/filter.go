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
			out = append(out, v)
			continue
		}
		next := map[string]any{}
		for k, val := range v.Properties {
			if _, ok := allowed[k]; ok {
				next[k] = val
			}
		}
		v.Properties = next
		out = append(out, v)
	}
	return out
}
func (f *Filter) NormalizeProperties(features []domain.Feature) []domain.Feature {
	for i := range features {
		p := map[string]any{}
		for k, v := range features[i].Properties {
			p[strings.TrimSpace(strings.ToLower(k))] = v
		}
		features[i].Properties = p
	}
	return features
}
