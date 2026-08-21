package domain

import (
	"bytes"
	"io"
	"testing"
)

type trackedReader struct {
	io.Reader
	closed bool
}

func (r *trackedReader) Close() error { r.closed = true; return nil }
func TestReadAllAndCloseReleasesReader(t *testing.T) {
	r := &trackedReader{Reader: bytes.NewBufferString("tile")}
	b, err := ReadAllAndClose(r)
	if err != nil || string(b) != "tile" || !r.closed {
		t.Fatalf("reader not closed: %q %v %v", b, err, r.closed)
	}
}
