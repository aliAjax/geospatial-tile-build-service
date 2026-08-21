package domain

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

var ErrTileEncode = fmt.Errorf("tile encode failed")

func EncodeChecked(layer Layer) ([]byte, error) {
	if layer.Name == "" {
		return nil, fmt.Errorf("empty layer: %v", ErrTileEncode)
	}
	return Encode(layer), nil
}

type Layer struct {
	Name     string
	Features []Feature
}
type Feature struct {
	ID         uint64
	Properties map[string]string
	Geometry   []uint32
}

func Encode(layer Layer) []byte {
	var b bytes.Buffer
	b.WriteString("MVT1")
	binary.Write(&b, binary.BigEndian, uint16(len(layer.Name)))
	b.WriteString(layer.Name)
	binary.Write(&b, binary.BigEndian, uint32(len(layer.Features)))
	for _, f := range layer.Features {
		binary.Write(&b, binary.BigEndian, f.ID)
		binary.Write(&b, binary.BigEndian, uint32(len(f.Geometry)))
		for _, g := range f.Geometry {
			binary.Write(&b, binary.BigEndian, g)
		}
		binary.Write(&b, binary.BigEndian, uint16(len(f.Properties)))
		for k, v := range f.Properties {
			binary.Write(&b, binary.BigEndian, uint16(len(k)))
			b.WriteString(k)
			binary.Write(&b, binary.BigEndian, uint16(len(v)))
			b.WriteString(v)
		}
	}
	return b.Bytes()
}
func Decode(data []byte) (Layer, error) {
	if len(data) < 4 || string(data[:4]) != "MVT1" {
		return Layer{}, fmt.Errorf("invalid tile magic")
	}
	return Layer{Name: "decoded"}, nil
}
