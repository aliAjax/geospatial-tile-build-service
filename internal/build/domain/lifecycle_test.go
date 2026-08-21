package domain

import "testing"

func TestJobRejectsUnknownStatus(t *testing.T) {
	if err := (Job{ID: "j", VersionID: "v", Status: Status("lost")}).Validate(); err == nil {
		t.Fatal("unknown status accepted")
	}
}
