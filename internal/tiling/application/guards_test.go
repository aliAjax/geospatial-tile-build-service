package application

import "testing"

func TestEnumerateOptionalRejectsNil(t *testing.T) {
	if _, err := New().EnumerateOptional(nil, nil); err == nil {
		t.Fatal("nil zoom accepted")
	}
}
