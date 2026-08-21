package main

import (
	"context"
	"errors"
	"github.com/example/geospatial-tile-build-service/api/http"
	"github.com/example/geospatial-tile-build-service/internal/platform/config"
	"github.com/example/geospatial-tile-build-service/internal/platform/health"
	"github.com/example/geospatial-tile-build-service/internal/platform/logging"
	"github.com/example/geospatial-tile-build-service/internal/platform/metrics"
	"github.com/example/geospatial-tile-build-service/internal/storage/infrastructure"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.Load()
	logger := logging.New()
	repo := infrastructure.New()
	h := &health.State{}
	m := &metrics.Registry{}
	srv := httpapi.New(repo, h, m, logger, cfg.MaxBody)
	server := &http.Server{Addr: cfg.HTTPAddr, Handler: srv.Routes(), ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 15 * time.Second, WriteTimeout: 20 * time.Second, IdleTimeout: 60 * time.Second}
	h.SetReady(true)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("server stopped", "error", err)
		}
	}()
	<-ctx.Done()
	shutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(shutdown)
}
