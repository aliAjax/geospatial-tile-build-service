package http

import (
	"encoding/json"
	"io"
	"net/http"
)

type Envelope struct {
	Data  any `json:"data,omitempty"`
	Error any `json:"error,omitempty"`
}

func ReadJSON(r *http.Request, dst any) error {
	dec := json.NewDecoder(io.LimitReader(r.Body, 8<<20))
	dec.DisallowUnknownFields()
	return dec.Decode(dst)
}
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func WriteError(w http.ResponseWriter, status int, code, message string) {
	WriteJSON(w, status, Envelope{Error: map[string]string{"code": code, "message": message}})
}
