package domain

import "testing"

func TestParseOptionalRejectsEmpty(t *testing.T) {
	if _, err := ParseOptional(""); err == nil {
		t.Fatal("empty path accepted")
	}
}
func TestValidateOptionalRejectsNil(t *testing.T) {
	if err := ValidateOptional(nil); err == nil {
		t.Fatal("nil coordinate accepted")
	}
}
func TestValidateMatrixRejectsNil(t *testing.T) {
	if err := ValidateMatrix(nil); err == nil {
		t.Fatal("nil matrix accepted")
	}
}
