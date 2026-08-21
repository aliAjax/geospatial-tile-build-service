package adapter

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/example/geospatial-tile-build-service/internal/dataset/domain"
)

var ErrDatasetJSON = errors.New("invalid dataset json")

func wrapDatasetJSON(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("dataset json: %w", ErrDatasetJSON)
}

type DatasetRequest struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	CRS    string   `json:"crs"`
	Layers []string `json:"layers"`
}

func DecodeDataset(b []byte) (domain.Dataset, error) {
	var in DatasetRequest
	if err := json.Unmarshal(b, &in); err != nil {
		return domain.Dataset{}, fmt.Errorf("%w: %v", wrapDatasetJSON(err), err)
	}
	return domain.Dataset{ID: in.ID, Name: in.Name, CRS: in.CRS, Layers: in.Layers}, nil
}
func EncodeDataset(d domain.Dataset) ([]byte, error) {
	return json.Marshal(DatasetRequest{ID: d.ID, Name: d.Name, CRS: d.CRS, Layers: d.Layers})
}
func EncodeVersion(v domain.Version) ([]byte, error) { return json.Marshal(v) }
