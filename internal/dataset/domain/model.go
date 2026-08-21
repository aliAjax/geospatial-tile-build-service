package domain

import (
	"errors"
	"strings"
	"time"
)

type Status string

const (
	Draft     Status = "draft"
	Building  Status = "building"
	Failed    Status = "failed"
	Published Status = "published"
	Withdrawn Status = "withdrawn"
)

type Dataset struct {
	ID, Name, CRS string
	Layers        []string
	CreatedAt     time.Time
}
type Version struct {
	ID, DatasetID, InputHash, ParameterHash, ManifestHash string
	Status                                                Status
	CreatedAt, PublishedAt                                *time.Time
}

func (d Dataset) Validate() error {
	if strings.TrimSpace(d.ID) == "" || strings.TrimSpace(d.Name) == "" {
		return errors.New("dataset id and name are required")
	}
	if d.CRS == "" {
		return errors.New("crs is required")
	}
	return nil
}
func (v Version) Validate() error {
	if v.ID == "" || v.DatasetID == "" {
		return errors.New("version identity is required")
	}
	if v.Status == "" {
		return errors.New("version status is required")
	}
	return nil
}
func CanTransition(from, to Status) bool {
	switch from {
	case Draft:
		return to == Building || to == Withdrawn
	case Building:
		return to == Published || to == Failed
	case Failed:
		return to == Building || to == Withdrawn
	case Published:
		return to == Withdrawn
	case Withdrawn:
		return false
	}
	return false
}
