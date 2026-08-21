package adapter

import (
	"encoding/json"
	"fmt"
	"github.com/example/geospatial-tile-build-service/internal/geometry/domain"
)

type GeoJSON struct {
	Type     string `json:"type"`
	Features []struct {
		Type     string `json:"type"`
		ID       any    `json:"id"`
		Geometry struct {
			Type        string          `json:"type"`
			Coordinates json.RawMessage `json:"coordinates"`
		} `json:"geometry"`
		Properties map[string]any `json:"properties"`
	} `json:"features"`
}

func Decode(data []byte) ([]domain.Feature, error) {
	var doc GeoJSON
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("decode geojson: %w", err)
	}
	out := make([]domain.Feature, 0, len(doc.Features))
	for i, f := range doc.Features {
		g, err := decodeGeometry(f.Geometry.Type, f.Geometry.Coordinates)
		if err != nil {
			return nil, err
		}
		id := fmt.Sprintf("%v", f.ID)
		if id == "<nil>" {
			id = fmt.Sprintf("feature-%d", i)
		}
		out = append(out, domain.Feature{ID: id, Geometry: g, Properties: f.Properties})
	}
	return out, nil
}
func decodeGeometry(t string, b []byte) (any, error) {
	switch t {
	case "Point":
		var c [2]float64
		if err := json.Unmarshal(b, &c); err != nil {
			return nil, err
		}
		return domain.Point{X: c[0], Y: c[1]}, nil
	case "LineString":
		var c [][2]float64
		if err := json.Unmarshal(b, &c); err != nil {
			return nil, err
		}
		l := make(domain.LineString, len(c))
		for i, p := range c {
			l[i] = domain.Point{X: p[0], Y: p[1]}
		}
		return l, nil
	case "Polygon":
		var c [][][2]float64
		if err := json.Unmarshal(b, &c); err != nil {
			return nil, err
		}
		p := domain.Polygon{Rings: make([][]domain.Point, len(c))}
		for i, r := range c {
			p.Rings[i] = make([]domain.Point, len(r))
			for j, v := range r {
				p.Rings[i][j] = domain.Point{X: v[0], Y: v[1]}
			}
		}
		return p, nil
	default:
		return nil, fmt.Errorf("unsupported geometry %s", t)
	}
}
