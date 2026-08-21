package application

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/example/geospatial-tile-build-service/internal/encoding/domain"
)

type Encoder struct{}

func New() *Encoder { return &Encoder{} }
func (e *Encoder) Encode(ctx context.Context, l domain.Layer) ([]byte, string, error) {
	select {
	case <-ctx.Done():
		return nil, "", ctx.Err()
	default:
	}
	b := domain.Encode(l)
	s := sha256.Sum256(b)
	return b, hex.EncodeToString(s[:]), nil
}
