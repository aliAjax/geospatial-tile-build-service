package domain

import (
	"errors"
	"testing"
)

func TestEncodeCheckedPreservesCause(t *testing.T) {
	_, err := EncodeChecked(Layer{})
	if !errors.Is(err, ErrTileEncode) {
		t.Fatalf("tile error lost: %v", err)
	}
}
func TestBuildTagsCheckedPreservesCause(t *testing.T) {
	_, err := BuildTagsChecked(nil)
	if !errors.Is(err, ErrTileEncode) {
		t.Fatalf("tag error lost: %v", err)
	}
}

func TestBuildTagsCheckedRejectsEmpty(t *testing.T) {
	if _, err := BuildTagsChecked(map[string]string{}); err == nil {
		t.Fatal("empty tags accepted")
	}
}
