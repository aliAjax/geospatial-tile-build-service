package http

import (
	"bytes"
	"io"
	"net/http/httptest"
	"testing"
)

type rangeCloser struct {
	io.Reader
	closed bool
}

func (r *rangeCloser) Close() error { r.closed = true; return nil }
func TestServeReaderRangeClosesBody(t *testing.T) {
	body := &rangeCloser{Reader: bytes.NewBufferString("abcd")}
	req := httptest.NewRequest("GET", "/tile", nil)
	req.Header.Set("Range", "bytes=0-1")
	ServeReaderRange(httptest.NewRecorder(), req, body)
	if !body.closed {
		t.Fatal("range body leaked")
	}
}
