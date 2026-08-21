package domain

import "time"

type Release struct {
	DatasetID, VersionID, Channel string
	Active                        bool
	UpdatedAt                     time.Time
}
