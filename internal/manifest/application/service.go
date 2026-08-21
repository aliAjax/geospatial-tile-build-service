package application

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/example/geospatial-tile-build-service/internal/manifest/domain"
	"time"
)

type Service struct{}

func New() *Service { return &Service{} }
func (s *Service) Create(ctx context.Context, dataset, version string, layers []string, raw []byte) (domain.Manifest, error) {
	select {
	case <-ctx.Done():
		return domain.Manifest{}, ctx.Err()
	default:
	}
	if dataset == "" || version == "" {
		return domain.Manifest{}, fmt.Errorf("dataset and version required")
	}
	h := sha256.Sum256(raw)
	return domain.Manifest{DatasetID: dataset, VersionID: version, MinZoom: 0, MaxZoom: 14, Layers: layers, Tiles: 1, Digest: hex.EncodeToString(h[:]), CreatedAt: time.Now().UTC()}, nil
}
