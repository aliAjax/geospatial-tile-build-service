package application

import (
	"github.com/example/geospatial-tile-build-service/internal/publication/domain"
	"sync"
	"time"
)

type History struct {
	mu    sync.RWMutex
	items []domain.Release
}

func NewHistory() *History { return &History{items: []domain.Release{}} }
func (h *History) Add(r domain.Release) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.items = append(h.items, r)
}
func (h *History) List(dataset string) []domain.Release {
	h.mu.RLock()
	defer h.mu.RUnlock()
	out := []domain.Release{}
	for _, r := range h.items {
		if r.DatasetID == dataset {
			out = append(out, r)
		}
	}
	return out
}
func (h *History) Latest(dataset, channel string) (domain.Release, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	var best domain.Release
	found := false
	for _, r := range h.items {
		if r.DatasetID == dataset && r.Channel == channel && (!found || r.UpdatedAt.After(best.UpdatedAt)) {
			best = r
			found = true
		}
	}
	return best, found
}
func expired(r domain.Release, now time.Time) bool { return now.Sub(r.UpdatedAt) > 365*24*time.Hour }
