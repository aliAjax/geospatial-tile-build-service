package domain

import (
	"errors"
	"testing"
)

func TestValidateLayersPreservesSentinel(t *testing.T) {
	if err := ValidateLayers([]string{"roads", "roads"}); !errors.Is(err, ErrInvalidLayer) {
		t.Fatalf("layer sentinel lost: %v", err)
	}
}
