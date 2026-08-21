package adapter

import (
	"encoding/json"
	"github.com/example/geospatial-tile-build-service/internal/manifest/domain"
)

func Encode(m domain.Manifest) ([]byte, error) { return json.Marshal(m) }
