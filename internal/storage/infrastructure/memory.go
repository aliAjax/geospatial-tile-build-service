package infrastructure

import (
	"bytes"
	"context"
	"errors"
	"github.com/example/geospatial-tile-build-service/internal/dataset/domain"
	"io"
	"sync"
)

type Memory struct {
	mu       sync.RWMutex
	items    map[string][]byte
	datasets map[string]domain.Dataset
	versions map[string]domain.Version
}

func New() *Memory {
	return &Memory{items: map[string][]byte{}, datasets: map[string]domain.Dataset{}, versions: map[string]domain.Version{}}
}
func (m *Memory) Put(_ context.Context, k string, b []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	// Always take an owned, immutable snapshot of the caller's bytes. A previous
	// version reused the backing array on overwrite (copy into the existing
	// slice) and exposed that same array through Open's reader. Under concurrent
	// Put/Open on the same key, a reader observed a half-written buffer (data
	// race) and build results drifted on partial input. A fresh allocation per
	// write means each stored slice is never mutated again, so Open can safely
	// hand out a reader over it without copying on every read.
	snap := make([]byte, len(b))
	copy(snap, b)
	m.items[k] = snap
	return nil
}
func (m *Memory) Open(_ context.Context, k string) (io.ReadCloser, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	b, ok := m.items[k]
	if !ok {
		return nil, errors.New("object not found")
	}
	// b is an immutable snapshot (see Put), so it is safe to share with callers
	// even while a later Put replaces m.items[k] with a different slice header.
	return io.NopCloser(bytes.NewReader(b)), nil
}
func (m *Memory) Delete(_ context.Context, k string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.items, k)
	return nil
}
func (m *Memory) PutDataset(_ context.Context, d domain.Dataset) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.datasets[d.ID]; ok {
		return errors.New("dataset already exists")
	}
	m.datasets[d.ID] = d
	return nil
}
func (m *Memory) GetDataset(_ context.Context, id string) (domain.Dataset, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	d, ok := m.datasets[id]
	if !ok {
		return d, errors.New("dataset not found")
	}
	return d, nil
}
func (m *Memory) PutVersion(_ context.Context, v domain.Version) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.versions[v.ID] = v
	return nil
}
func (m *Memory) GetVersion(_ context.Context, id string) (domain.Version, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.versions[id]
	if !ok {
		return v, errors.New("version not found")
	}
	return v, nil
}
func (m *Memory) ListVersions(_ context.Context, did string) ([]domain.Version, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := []domain.Version{}
	for _, v := range m.versions {
		if v.DatasetID == did {
			out = append(out, v)
		}
	}
	return out, nil
}
