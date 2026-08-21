package adapter

import (
	"encoding/binary"
	"github.com/example/geospatial-tile-build-service/internal/encoding/domain"
)

func ZigZag(v int32) uint32     { return uint32(uint32(v<<1) ^ uint32(v>>31)) }
func EncodeID(id uint64) []byte { b := make([]byte, 8); binary.BigEndian.PutUint64(b, id); return b }
func LayerFromFeatures(name string, ids []uint64) domain.Layer {
	f := make([]domain.Feature, len(ids))
	for i, id := range ids {
		f[i] = domain.Feature{ID: id}
	}
	return domain.Layer{Name: name, Features: f}
}
