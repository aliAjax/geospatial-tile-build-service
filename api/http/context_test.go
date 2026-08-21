package httpapi

import (
	"context"
	"github.com/example/geospatial-tile-build-service/internal/platform/health"
	"github.com/example/geospatial-tile-build-service/internal/platform/metrics"
	"github.com/example/geospatial-tile-build-service/internal/storage/infrastructure"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestManifestHTTPHonorsRequestCancellation(t *testing.T) {
	s := New(infrastructure.New(), &health.State{}, &metrics.Registry{}, slog.New(slog.NewTextHandler(io.Discard, nil)), 1024)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	r := httptest.NewRequest(http.MethodGet, "/v1/manifests/roads/v1", nil).WithContext(ctx)
	w := httptest.NewRecorder()
	s.Routes().ServeHTTP(w, r)
	if w.Code == http.StatusOK {
		t.Fatal("canceled manifest request succeeded")
	}
}
