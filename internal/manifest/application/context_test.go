package application

import (
	"context"
	"testing"
)

func TestManifestRequestContextDoesNotLeak(t *testing.T) {
	s := New()
	canceled, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := s.Create(canceled, "roads", "v1", []string{"default"}, []byte("x")); err == nil {
		t.Fatal("canceled request was accepted")
	}
	if _, err := s.Create(context.Background(), "roads", "v2", []string{"default"}, []byte("y")); err != nil {
		t.Fatalf("later request inherited cancellation: %v", err)
	}
}
