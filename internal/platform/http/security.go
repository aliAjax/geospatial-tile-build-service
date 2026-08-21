package http

import "net/http"

func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Referrer-Policy", "no-referrer")
		next.ServeHTTP(w, r)
	})
}
func Recover(next http.Handler, fail func(http.ResponseWriter, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if v := recover(); v != nil {
				if fail != nil {
					fail(w, panicError{v})
				} else {
					http.Error(w, "internal server error", 500)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}

type panicError struct{ v any }

func (p panicError) Error() string { return "panic" }
