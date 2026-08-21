package httpapi

import (
	"encoding/json"
	"net/http"
	"time"
)

type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}
type ErrorBody struct {
	Code      string    `json:"code"`
	Message   string    `json:"message"`
	RequestID string    `json:"request_id"`
	Time      time.Time `json:"time"`
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(ErrorResponse{Error: ErrorBody{Code: code, Message: message, Time: time.Now().UTC()}})
}
func method(w http.ResponseWriter, r *http.Request, want string) bool {
	if r.Method != want {
		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "method must be "+want)
		return false
	}
	return true
}
