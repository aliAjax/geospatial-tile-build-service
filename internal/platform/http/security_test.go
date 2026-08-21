package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecoverMiddlewareCatchesPanicAndLogs(t *testing.T) {
	called := false
	h := Recover(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { panic("boom") }), func(w http.ResponseWriter, _ error) {
		called = true
		w.WriteHeader(http.StatusInternalServerError)
	})
	r := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if !called || w.Code != http.StatusInternalServerError {
		t.Fatalf("panic was not converted: called=%v status=%d", called, w.Code)
	}
}
