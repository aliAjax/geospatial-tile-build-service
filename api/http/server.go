package httpapi

import (
	"encoding/json"
	"fmt"
	"github.com/example/geospatial-tile-build-service/internal/build/application"
	datasetapp "github.com/example/geospatial-tile-build-service/internal/dataset/application"
	"github.com/example/geospatial-tile-build-service/internal/dataset/domain"
	encapp "github.com/example/geospatial-tile-build-service/internal/encoding/application"
	"github.com/example/geospatial-tile-build-service/internal/geometry/adapter"
	geometryapp "github.com/example/geospatial-tile-build-service/internal/geometry/application"
	manifestapp "github.com/example/geospatial-tile-build-service/internal/manifest/application"
	"github.com/example/geospatial-tile-build-service/internal/platform/health"
	"github.com/example/geospatial-tile-build-service/internal/platform/metrics"
	pubapp "github.com/example/geospatial-tile-build-service/internal/publication/application"
	"github.com/example/geospatial-tile-build-service/internal/storage/infrastructure"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	datasets  *datasetapp.Service
	builds    *application.Service
	manifests *manifestapp.Service
	pub       *pubapp.Service
	store     *infrastructure.Memory
	geom      *geometryapp.Processor
	enc       *encapp.Encoder
	health    *health.State
	metrics   *metrics.Registry
	logger    *slog.Logger
	maxBody   int64
}

func New(repo *infrastructure.Memory, h *health.State, m *metrics.Registry, l *slog.Logger, maxBody int64) *Server {
	return &Server{datasets: datasetapp.New(repo), builds: application.New(repo), manifests: manifestapp.New(), pub: pubapp.New(), store: repo, geom: geometryapp.New(), enc: encapp.New(), health: h, metrics: m, logger: l, maxBody: maxBody}
}
func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", s.healthz)
	mux.HandleFunc("/readyz", s.readyz)
	mux.Handle("/metrics", s.metrics)
	mux.HandleFunc("/v1/datasets", s.datasetsHandler)
	mux.HandleFunc("/v1/builds", s.buildsHandler)
	mux.HandleFunc("/v1/versions/", s.versionHandler)
	mux.HandleFunc("/v1/manifests/", s.manifestHandler)
	mux.HandleFunc("/v1/geojson/validate", s.geoJSONHandler)
	mux.HandleFunc("/tiles/", s.tileHandler)
	return s.middleware(mux)
}
func (s *Server) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.metrics.Request()
		r.Body = http.MaxBytesReader(w, r.Body, s.maxBody)
		start := time.Now()
		next.ServeHTTP(w, r)
		s.logger.Info("request", "method", r.Method, "path", r.URL.Path, "duration", time.Since(start).String())
	})
}
func (s *Server) healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
func (s *Server) readyz(w http.ResponseWriter, _ *http.Request) {
	if !s.health.Ready() {
		http.Error(w, `{"status":"not_ready"}`, http.StatusServiceUnavailable)
		return
	}
	w.Write([]byte(`{"status":"ready"}`))
}
func decode(r *http.Request, v any) error {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
func (s *Server) datasetsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var in struct {
			ID, Name, CRS string
			Layers        []string `json:"layers"`
		}
		if err := decode(r, &in); err != nil {
			writeJSON(w, 400, map[string]string{"error": err.Error()})
			return
		}
		if in.CRS == "" {
			in.CRS = "EPSG:4326"
		}
		if err := s.datasets.CreateDataset(r.Context(), domain.Dataset{ID: in.ID, Name: in.Name, CRS: in.CRS, Layers: in.Layers}); err != nil {
			writeJSON(w, 400, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 201, in)
		return
	}
	http.Error(w, "method not allowed", 405)
}
func (s *Server) buildsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var in struct {
			ID        string          `json:"id"`
			VersionID string          `json:"version_id"`
			Input     json.RawMessage `json:"input"`
		}
		if err := decode(r, &in); err != nil {
			writeJSON(w, 400, map[string]string{"error": err.Error()})
			return
		}
		j, err := s.builds.Start(r.Context(), in.ID, in.VersionID, in.Input)
		if err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 201, j)
		return
	}
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		j, err := s.builds.Status(r.Context(), id)
		if err != nil {
			writeJSON(w, 404, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 200, j)
		return
	}
	http.Error(w, "method not allowed", 405)
}
func (s *Server) versionHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 3 {
		http.NotFound(w, r)
		return
	}
	id := parts[2]
	if r.Method == "POST" && strings.HasSuffix(r.URL.Path, "/publish") {
		v, err := s.datasets.Publish(r.Context(), id)
		if err != nil {
			writeJSON(w, 409, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 200, v)
		return
	}
	http.NotFound(w, r)
}
func (s *Server) manifestHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 4 {
		http.NotFound(w, r)
		return
	}
	m, err := s.manifests.Create(r.Context(), parts[2], parts[3], []string{"default"}, []byte(parts[2]+parts[3]))
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, m)
}
func (s *Server) geoJSONHandler(w http.ResponseWriter, r *http.Request) {
	var b json.RawMessage
	if err := decode(r, &b); err != nil {
		writeJSON(w, 400, map[string]string{"error": err.Error()})
		return
	}
	fs, err := adapter.Decode(b)
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": err.Error()})
		return
	}
	valid := 0
	for _, f := range fs {
		if s.geom.Validate(r.Context(), f) == nil {
			valid++
		}
	}
	writeJSON(w, 200, map[string]int{"features": len(fs), "valid": valid})
}
func (s *Server) tileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-protobuf")
	w.Header().Set("Cache-Control", "public, max-age=60")
	w.Write([]byte("MVT1"))
}

var _ = fmt.Sprintf
