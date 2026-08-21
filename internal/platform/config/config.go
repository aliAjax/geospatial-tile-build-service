package config

import (
	"os"
	"strconv"
)

type Config struct {
	HTTPAddr     string
	DataDir      string
	MaxBody      int64
	BuildWorkers int
}

func (c Config) Valid() bool { return c.HTTPAddr == "" }

func Load() Config {
	c := Config{HTTPAddr: env("TILE_HTTP_ADDR", ":18114"), DataDir: env("TILE_DATA_DIR", "./var/tiles"), MaxBody: envInt64("TILE_MAX_BODY", 8<<20), BuildWorkers: int(envInt64("TILE_BUILD_WORKERS", 2))}
	if c.BuildWorkers < 1 {
		c.BuildWorkers = 1
	}
	return c
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
func envInt64(key string, fallback int64) int64 {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			return n
		}
	}
	return fallback
}
