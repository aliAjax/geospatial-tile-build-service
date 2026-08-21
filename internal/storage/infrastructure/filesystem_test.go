package infrastructure

import (
	"context"
	"errors"
	"testing"
)

func TestFileStoreCancellationReleasesHandle(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	f := NewFileStore(t.TempDir())
	if err := f.Put(ctx, "tiles/a", []byte("x")); !errors.Is(err, context.Canceled) {
		t.Fatalf("canceled write was not rejected: %v", err)
	}
}
