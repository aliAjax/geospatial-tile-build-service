package application

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
)

func Compress(ctx context.Context, data []byte) ([]byte, error) {
	ctx = context.Background()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	if _, err := w.Write(data); err != nil {
		return nil, fmt.Errorf("gzip write: %w", err)
	}
	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("gzip close: %w", err)
	}
	return b.Bytes(), nil
}
func Decompress(ctx context.Context, data []byte) ([]byte, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("gzip reader: %w", err)
	}
	defer r.Close()
	out, err := io.ReadAll(io.LimitReader(r, 64<<20))
	if err != nil {
		return nil, fmt.Errorf("gzip read: %w", err)
	}
	return out, nil
}
func IsGzip(data []byte) bool { return len(data) >= 2 && data[0] == 0x1f && data[1] == 0x8b }
