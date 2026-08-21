package http

import (
	"crypto/subtle"
	"net/http"
)

type TokenAuth struct{ Token string }

func (a TokenAuth) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if a.Token != "" {
			got := r.Header.Get("Authorization")
			want := "Bearer " + a.Token
			if subtle.ConstantTimeCompare([]byte(got), []byte(want)) != 1 {
				WriteError(w, 401, "unauthorized", "valid bearer token required")
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
