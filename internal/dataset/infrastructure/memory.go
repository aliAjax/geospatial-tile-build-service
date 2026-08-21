package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"github.com/example/geospatial-tile-build-service/internal/dataset/domain"
	"sync"
)

var ErrDatasetNotFound = errors.New("dataset not found")

func IsDatasetNotFound(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, ErrDatasetNotFound)
}

type Memory struct {
	mu       sync.RWMutex
	datasets map[string]domain.Dataset
	versions map[string]domain.Version
}

func New() *Memory {
	return &Memory{datasets: map[string]domain.Dataset{}, versions: map[string]domain.Version{}}
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
		return d, fmt.Errorf("dataset lookup: %w", ErrDatasetNotFound)
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
