package domain

import (
	"sort"
	"strings"
	"time"
)

type Query struct {
	Text     string
	Statuses []Status
	Limit    int
	Cursor   string
}

func (q Query) Normalize() Query {
	q.Text = strings.TrimSpace(strings.ToLower(q.Text))
	if q.Limit <= 0 || q.Limit > 100 {
		q.Limit = 50
	}
	return q
}
func Match(d Dataset, q Query) bool {
	q = q.Normalize()
	if q.Text != "" && !strings.Contains(strings.ToLower(d.Name), q.Text) {
		return false
	}
	return true
}
func SortDatasets(items []Dataset) []Dataset {
	out := append([]Dataset(nil), items...)
	sort.Slice(out, func(i, j int) bool {
		if out[i].CreatedAt.Equal(out[j].CreatedAt) {
			return out[i].ID < out[j].ID
		}
		return out[i].CreatedAt.Before(out[j].CreatedAt)
	})
	return out
}
func VersionAge(v Version, now time.Time) time.Duration {
	if v.CreatedAt == nil {
		return 0
	}
	return now.Sub(*v.CreatedAt)
}
func Latest(items []Version) Version {
	var out Version
	for _, v := range items {
		if out.CreatedAt == nil || (v.CreatedAt != nil && v.CreatedAt.After(*out.CreatedAt)) {
			out = v
		}
	}
	return out
}
