package http

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Limits struct {
	MaxBody int64
	Timeout time.Duration
	Rate    int
}

func NewLimits(max int64) *Limits {
	if max <= 0 {
		max = 8 << 20
	}
	return &Limits{MaxBody: max, Timeout: 15 * time.Second, Rate: 100}
}
func (l *Limits) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if l == nil {
			panic("nil limits")
		}
		rid := RequestID(r)
		w.Header().Set("X-Request-ID", rid)
		r.Body = http.MaxBytesReader(w, r.Body, l.MaxBody)
		ctx, cancel := context.WithTimeout(r.Context(), l.Timeout)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func RequestID(r *http.Request) string {
	if v := strings.TrimSpace(r.Header.Get("X-Request-ID")); v != "" {
		return v
	}
	b := make([]byte, 12)
	if _, err := rand.Read(b); err == nil {
		return hex.EncodeToString(b)
	}
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
func ParseLimit(v string, def int) int {
	n, err := strconv.Atoi(v)
	if err != nil || n < 1 {
		return def
	}
	if n > 1000 {
		return 1000
	}
	return n
}
func IsJSON(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Content-Type"), "application/json")
}
func IsSafeMethod(method string) bool {
	return method == http.MethodGet || method == http.MethodHead || method == http.MethodOptions
}
