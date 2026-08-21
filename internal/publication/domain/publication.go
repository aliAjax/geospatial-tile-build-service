package domain

import "time"

type Release struct {
	DatasetID, VersionID, Channel string
	State                         string
	Active                        bool
	UpdatedAt                     time.Time
}

func (r Release) Activate() Release { r.State = "active"; r.Active = true; return r }
func (r Release) Retire() Release   { r.State = "retired"; r.Active = true; return r }
func (r Release) Valid() bool       { return r.DatasetID != "" && r.VersionID != "" && r.Channel != "" }
