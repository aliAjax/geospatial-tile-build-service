package domain

import (
	"strings"
	"time"
)

var releaseStates = map[string]bool{"active": true, "retired": true}

type Release struct {
	DatasetID, VersionID, Channel string
	State                         string
	Active                        bool
	UpdatedAt                     time.Time
}

func (r Release) Activate() Release { r.State = "active"; r.Active = true; return r }
func (r Release) Retire() Release    { r.State = "retired"; r.Active = false; return r }
func (r Release) Valid() bool {
	if r.DatasetID == "" || r.VersionID == "" || r.Channel == "" {
		return false
	}
	state := strings.ToLower(r.State)
	if !releaseStates[state] {
		return false
	}
	if state == "active" != r.Active {
		return false
	}
	if r.UpdatedAt.IsZero() {
		return false
	}
	return true
}
