package adapter

import (
	"errors"
	"github.com/example/geospatial-tile-build-service/internal/encoding/domain"
	"testing"
)

func TestEncodeFeatureCheckedPreservesCause(t *testing.T) {
	_, err := EncodeFeatureChecked("", nil)
	if !errors.Is(err, domain.ErrTileEncode) {
		t.Fatalf("feature error lost: %v", err)
	}
}
