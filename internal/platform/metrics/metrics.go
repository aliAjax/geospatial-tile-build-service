package metrics

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

type Registry struct {
	requests atomic.Uint64
	builds   atomic.Uint64
	failures atomic.Uint64
}

func (r *Registry) Request() { r.requests.Add(1) }
func (r *Registry) Build()   { r.builds.Add(1) }
func (r *Registry) Failure() { r.failures.Add(1) }
func (r *Registry) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; version=0.0.4")
	fmt.Fprintf(w, "tile_http_requests_total %d\ntile_builds_total %d\ntile_build_failures_total %d\n", r.requests.Load(), r.builds.Load(), r.failures.Load())
}
