package application

import (
	"context"
	"github.com/example/geospatial-tile-build-service/internal/publication/domain"
	"testing"
	"time"
)

func TestRetryPublicationReachesPublishedState(t *testing.T) {
	s := New()
	if _, err := s.Publish(context.Background(), "roads", "v1", "stable"); err != nil {
		t.Fatal(err)
	}
	r, err := s.Publish(context.Background(), "roads", "v2", "stable")
	if err != nil {
		t.Fatal(err)
	}
	if !r.Active || r.VersionID != "v2" {
		t.Fatalf("retry did not become current: %+v", r)
	}
	cur, err := s.Current(context.Background(), "roads", "stable")
	if err != nil || !cur.Active || cur.VersionID != "v2" {
		t.Fatalf("current release disagrees: %+v %v", cur, err)
	}
}

func TestPublicationCurrentReleaseActive(t *testing.T) {
	s := New()
	_, _ = s.Publish(context.Background(), "roads", "v1", "stable")
	_, _ = s.Publish(context.Background(), "roads", "v2", "stable")
	cur, err := s.Current(context.Background(), "roads", "stable")
	if err != nil || !cur.Active || cur.VersionID != "v2" {
		t.Fatalf("current release is stale: %+v %v", cur, err)
	}
}

func TestPublicationHistoryLatest(t *testing.T) {
	h := NewHistory()
	h.Add(domain.Release{DatasetID: "roads", VersionID: "v1", Channel: "stable", Active: false, UpdatedAt: time.Unix(10, 0)})
	h.Add(domain.Release{DatasetID: "roads", VersionID: "v2", Channel: "stable", Active: true, UpdatedAt: time.Unix(20, 0)})
	latest, ok := h.Latest("roads", "stable")
	if !ok || latest.VersionID != "v2" || !latest.Active {
		t.Fatalf("history latest is stale: %+v %v", latest, ok)
	}
}
