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
	m.items[k] = append([]byte(nil), b...)
	return nil
}
func (m *Memory) Open(_ context.Context, k string) (io.ReadCloser, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	b, ok := m.items[k]
	if !ok {
		return nil, errors.New("object not found")
	}
	return io.NopCloser(bytes.NewReader(append([]byte(nil), b...))), nil
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
