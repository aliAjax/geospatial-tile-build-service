package domain

import "fmt"

func (m Manifest) Validate() error {
	if m.DatasetID == "" || m.VersionID == "" {
		return fmt.Errorf("manifest identity required")
	}
	if m.MinZoom < 0 || m.MaxZoom < m.MinZoom {
		return fmt.Errorf("manifest zoom range invalid")
	}
	if len(m.Layers) == 0 {
		return fmt.Errorf("manifest layers required")
	}
	if m.Digest == "" {
		return fmt.Errorf("manifest digest required")
	}
	return nil
}
func (m Manifest) TileCount() uint64 {
	var n uint64
	for z := m.MinZoom; z <= m.MaxZoom; z++ {
		n += uint64(1) << (2 * z)
	}
	return n
}
func (m *Manifest) Finalize() { m.Tiles = m.TileCount() }
