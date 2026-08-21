package application

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"
)

type Change struct {
	Kind  string
	Value string
}

func DiffLayers(before, after []string) []Change {
	b := map[string]bool{}
	a := map[string]bool{}
	for _, v := range before {
		b[v] = true
	}
	for _, v := range after {
		a[v] = true
	}
	out := []Change{}
	for v := range a {
		if !b[v] {
			out = append(out, Change{"added", v})
		}
	}
	for v := range b {
		if !a[v] {
			out = append(out, Change{"removed", v})
		}
	}
	sort.Slice(out, func(i, j int) bool { return strings.Compare(out[i].Value, out[j].Value) < 0 })
	return out
}
func DigestParts(parts ...string) string {
	h := sha256.New()
	for _, v := range parts {
		h.Write([]byte(v))
		h.Write([]byte{0})
	}
	return hex.EncodeToString(h.Sum(nil))
}
