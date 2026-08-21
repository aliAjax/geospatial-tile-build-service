package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestContentMiddlewarePreservesContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	seen := false
	h := RequireContentType(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) { seen = r.Context().Err() == context.Canceled }))
	h.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", "/v1/x", nil).WithContext(ctx))
	if !seen {
		t.Fatal("request cancellation was replaced")
	}
}
