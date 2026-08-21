package application

import (
	"bytes"
	"context"
	"github.com/example/geospatial-tile-build-service/internal/tiling/domain"
	"io"
	"testing"
)

type closeReader struct {
	io.Reader
	closed bool
}

func (r *closeReader) Close() error { r.closed = true; return nil }

type closeStore struct{ reader *closeReader }

func (s closeStore) Put(context.Context, string, []byte) error           { return nil }
func (s closeStore) Delete(context.Context, string) error                { return nil }
func (s closeStore) Open(context.Context, string) (io.ReadCloser, error) { return s.reader, nil }
func TestOpenCancellationClosesReader(t *testing.T) {
	r := &closeReader{Reader: bytes.NewBuffer(nil)}
	ctx, c := context.WithCancel(context.Background())
	c()
	_, err := New(closeStore{r}).Open(ctx, domain.Coordinate{Z: 0, X: 0, Y: 0})
	if err == nil || !r.closed {
		t.Fatalf("reader leaked: %v %v", err, r.closed)
	}
}
