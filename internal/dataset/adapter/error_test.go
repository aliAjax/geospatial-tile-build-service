package adapter

import (
	"errors"
	"testing"
)

func TestDecodeDatasetPreservesJSONError(t *testing.T) {
	_, err := DecodeDataset([]byte("{"))
	if !errors.Is(err, ErrDatasetJSON) {
		t.Fatalf("json sentinel lost: %v", err)
	}
}
