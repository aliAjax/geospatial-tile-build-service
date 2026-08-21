package http

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strings"
)

func ETag(data []byte) string {
	h := sha256.Sum256(data)
	return `"` + hex.EncodeToString(h[:])[:32] + `"`
}
func Conditional(w http.ResponseWriter, r *http.Request, tag string) bool {
	w.Header().Set("ETag", tag)
	return strings.TrimSpace(r.Header.Get("If-None-Match")) == tag
}
func CacheHeaders(w http.ResponseWriter, maxAge int) {
	w.Header().Set("Cache-Control", "public, max-age="+itoa(maxAge))
	w.Header().Set("Vary", "Accept-Encoding")
}
func itoa(v int) string {
	if v == 0 {
		return "0"
	}
	s := ""
	for v > 0 {
		s = string(rune('0'+v%10)) + s
		v /= 10
	}
	return s
}
