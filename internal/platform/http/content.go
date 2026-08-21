package http

import (
	"net/http"
	"strings"
)

func RequireContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodPatch {
			if !strings.Contains(r.Header.Get("Content-Type"), "application/json") && strings.Contains(r.URL.Path, "/v1/") {
				WriteError(w, 415, "unsupported_media_type", "application/json required")
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
func NoStore(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}
