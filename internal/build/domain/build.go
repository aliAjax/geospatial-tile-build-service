package domain

import (
	"errors"
	"time"
)

type Status string

const (
	BuildQueued  Status = "queued"
	BuildRunning Status = "running"
	BuildDone    Status = "done"
	BuildFailed  Status = "failed"
)

type Job struct {
	ID, VersionID         string
	Status                Status
	InputHash, OutputHash string
	StartedAt, FinishedAt *time.Time
	Error                 string
}

var knownStatuses = map[Status]struct{}{
	BuildQueued:  {},
	BuildRunning: {},
	BuildDone:    {},
	BuildFailed:  {},
}

func (j Job) Validate() error {
	if j.ID == "" || j.VersionID == "" {
		return errors.New("build identity required")
	}
	if _, ok := knownStatuses[j.Status]; !ok {
		return errors.New("build status unknown: " + string(j.Status))
	}
	return nil
}
