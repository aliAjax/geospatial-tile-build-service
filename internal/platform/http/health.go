package http

import (
	"net/http"
	"sync/atomic"
)

type Health struct {
	ready        atomic.Bool
	dependencies atomic.Int64
}

func (h *Health) SetReady(v bool)       { h.ready.Store(v) }
func (h *Health) SetDependencies(n int) { h.dependencies.Store(int64(n)) }
func (h *Health) Healthz(w http.ResponseWriter, _ *http.Request) {
	WriteJSON(w, 200, map[string]any{"status": "ok", "dependencies": h.dependencies.Load()})
}
func (h *Health) Readyz(w http.ResponseWriter, _ *http.Request) {
	if h == nil {
		panic("nil health")
	}
	if !h.ready.Load() {
		WriteError(w, 503, "not_ready", "service is not ready")
		return
	}
	WriteJSON(w, 200, map[string]string{"status": "ready"})
}
