package config

import "testing"

func TestDefaultConfigValid(t *testing.T) {
	if !Load().Valid() {
		t.Fatal("default config invalid")
	}
}
