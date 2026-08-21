package application

import (
	"context"
	"errors"
	"testing"
)

func TestCompressHonorsCancellation(t *testing.T) {
	ctx, c := context.WithCancel(context.Background())
	c()
	_, err := Compress(ctx, []byte("x"))
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("cancel lost: %v", err)
	}
}
