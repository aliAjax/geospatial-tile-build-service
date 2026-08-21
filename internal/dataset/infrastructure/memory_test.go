package infrastructure

import (
	"context"
	"errors"
	"testing"
)

func TestDatasetNotFoundPreservesHTTPStatus(t *testing.T) {
	m := New()
	_, err := m.GetDataset(context.Background(), "missing")
	if err == nil || !errors.Is(err, ErrDatasetNotFound) {
		t.Fatalf("missing dataset error lost sentinel: %v", err)
	}
}
