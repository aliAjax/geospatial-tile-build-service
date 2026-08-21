package config

import (
	buildapp "github.com/example/geospatial-tile-build-service/internal/build/application"
	"os"
	"testing"
)

func TestDefaultDatasetConfigDoesNotPanic(t *testing.T) {
	for _, key := range []string{"TILE_HTTP_ADDR", "TILE_MAX_BODY", "TILE_BUILD_WORKERS"} {
		_ = os.Unsetenv(key)
	}
	cfg := Load()
	if cfg.HTTPAddr == "" || cfg.MaxBody <= 0 || cfg.BuildWorkers <= 0 {
		t.Fatalf("invalid defaults: %+v", cfg)
	}
	c := buildapp.NewCache(cfg.BuildWorkers)
	c.Put("default", []byte("ready"))
	if got, ok := c.Get("default"); !ok || string(got) != "ready" {
		t.Fatalf("default cache is unusable: %q %v", got, ok)
	}
}
