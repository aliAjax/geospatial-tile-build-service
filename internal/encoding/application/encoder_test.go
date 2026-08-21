package application

import (
	"context"
	"errors"
	"github.com/example/geospatial-tile-build-service/internal/encoding/domain"
	"testing"
)

func TestEncoderPreservesWriteErrorChain(t *testing.T) {
	enc := New()
	canceled, cancel := context.WithCancel(context.Background())
	cancel()
	_, _, err := enc.Encode(canceled, domain.Layer{Name: "roads"})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("encoding error chain was lost: %v", err)
	}
}
