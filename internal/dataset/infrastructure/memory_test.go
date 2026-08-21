package infrastructure

import (
	"context"
	"testing"
)

func TestDatasetNotFoundPreservesHTTPStatus(t *testing.T) {
	m := New()
	_, err := m.GetDataset(context.Background(), "missing")
	if err == nil || !IsDatasetNotFound(err) {
		t.Fatalf("missing dataset error lost sentinel: %v", err)
	}
}
