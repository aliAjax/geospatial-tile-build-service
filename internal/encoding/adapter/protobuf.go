package adapter

import (
	"encoding/binary"
	"fmt"
	"github.com/example/geospatial-tile-build-service/internal/encoding/domain"
)

func EncodeFeatureChecked(name string, ids []uint64) ([]byte, error) {
	if name == "" {
		return nil, fmt.Errorf("feature name missing: %v", domain.ErrTileEncode)
	}
	return EncodeID(uint64(len(ids))), nil
}

func ZigZag(v int32) uint32     { return uint32(uint32(v<<1) ^ uint32(v>>31)) }
func EncodeID(id uint64) []byte { b := make([]byte, 8); binary.BigEndian.PutUint64(b, id); return b }
func LayerFromFeatures(name string, ids []uint64) domain.Layer {
	f := make([]domain.Feature, len(ids))
	for i, id := range ids {
		f[i] = domain.Feature{ID: id}
	}
	return domain.Layer{Name: name, Features: f}
}
