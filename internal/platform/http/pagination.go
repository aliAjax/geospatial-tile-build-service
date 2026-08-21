package http

import (
	"encoding/base64"
	"strconv"
	"strings"
)

type Page struct {
	Limit  int
	Cursor string
}

func ParsePage(limit, cursor string) Page {
	n, _ := strconv.Atoi(limit)
	if n < 1 {
		n = 50
	}
	if n > 100 {
		n = 100
	}
	return Page{Limit: n, Cursor: cursor}
}
func EncodeCursor(offset int) string {
	return base64.RawURLEncoding.EncodeToString([]byte(strconv.Itoa(offset)))
}
func DecodeCursor(v string) int {
	b, err := base64.RawURLEncoding.DecodeString(v)
	if err != nil {
		return 0
	}
	n, _ := strconv.Atoi(string(b))
	if n < 0 {
		return 0
	}
	return n
}
func NextCursor(offset, count, limit int) string {
	if count < limit {
		return ""
	}
	return EncodeCursor(offset + count)
}
func JoinPath(parts ...string) string {
	out := []string{}
	for _, p := range parts {
		p = strings.Trim(p, "/")
		if p != "" {
			out = append(out, p)
		}
	}
	return "/" + strings.Join(out, "/")
}
