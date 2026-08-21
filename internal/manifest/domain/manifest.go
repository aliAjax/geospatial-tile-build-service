package domain

import (
	"encoding/json"
	"time"
)

type Manifest struct {
	DatasetID, VersionID string
	MinZoom, MaxZoom     int
	Layers               []string
	Bounds               [4]float64
	Tiles                uint64
	Digest               string
	CreatedAt            time.Time
}

func (m Manifest) JSON() ([]byte, error) { return json.MarshalIndent(m, "", "  ") }
