package http

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func ServeReaderRange(w http.ResponseWriter, r *http.Request, body io.ReadCloser) {
	if body == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := io.ReadAll(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ServeRange(w, r, data)
}

type ByteRange struct{ Start, End int64 }

func ParseRange(header string, size int64) (ByteRange, error) {
	if !strings.HasPrefix(header, "bytes=") {
		return ByteRange{}, fmt.Errorf("invalid range")
	}
	p := strings.Split(strings.TrimPrefix(header, "bytes="), "-")
	if len(p) != 2 {
		return ByteRange{}, fmt.Errorf("invalid range")
	}
	start, _ := strconv.ParseInt(p[0], 10, 64)
	end := size - 1
	if p[1] != "" {
		n, e := strconv.ParseInt(p[1], 10, 64)
		if e != nil {
			return ByteRange{}, e
		}
		end = n
	}
	if start < 0 || start > end || end >= size {
		return ByteRange{}, fmt.Errorf("range out of bounds")
	}
	return ByteRange{Start: start, End: end}, nil
}
func ServeRange(w http.ResponseWriter, r *http.Request, data []byte) {
	br, err := ParseRange(r.Header.Get("Range"), int64(len(data)))
	if err != nil {
		w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
		return
	}
	w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", br.Start, br.End, len(data)))
	w.WriteHeader(http.StatusPartialContent)
	_, _ = w.Write(data[br.Start : br.End+1])
}
