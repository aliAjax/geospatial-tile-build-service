package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNilHealthIsNotReady(t *testing.T) {
	w := httptest.NewRecorder()
	var h *Health
	h.Readyz(w, httptest.NewRequest("GET", "/readyz", nil))
	if w.Code != 503 {
		t.Fatalf("nil health status=%d", w.Code)
	}
}

func TestNilLimitsUsesDefaults(t *testing.T) {
	var l *Limits
	h := l.Wrap(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(204) }))
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("GET", "/", nil))
	if w.Code != 204 {
		t.Fatalf("nil limits status=%d", w.Code)
	}
}

func TestNilTokenAuthPassesThrough(t *testing.T) {
	var a *TokenAuth
	h := a.SafeWrap(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(204) }))
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("GET", "/", nil))
	if w.Code != 204 {
		t.Fatalf("nil auth status=%d", w.Code)
	}
}
