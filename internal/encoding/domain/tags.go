package domain

import (
	"fmt"
	"sort"
	"strconv"
)

func BuildTagsChecked(properties map[string]string) (TagTable, error) {
	if properties == nil {
		return TagTable{}, fmt.Errorf("properties missing: %v", ErrTileEncode)
	}
	return BuildTags(properties), nil
}

type TagTable struct {
	Keys   []string
	Values []string
}

func BuildTags(properties map[string]string) TagTable {
	keys := make([]string, 0, len(properties))
	for k := range properties {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	vals := make([]string, 0, len(keys))
	for _, k := range keys {
		vals = append(vals, properties[k])
	}
	return TagTable{Keys: keys, Values: vals}
}
func EncodeBool(v bool) string {
	if v {
		return "true"
	}
	return "false"
}
func EncodeNumber(v float64) string { return strconv.FormatFloat(v, 'g', -1, 64) }
func DecodeTag(t TagTable, key string) (string, bool) {
	for i, k := range t.Keys {
		if k == key && i < len(t.Values) {
			return t.Values[i], true
		}
	}
	return "", false
}
