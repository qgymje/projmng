package event

import "time"

type ProjectEventType int

const (
	EventTypeCreate ProjectEventType = iota
	EventTypeUpdate
	EventTypeDelete
)

type Project struct {
	ProjectID   int
	Type        ProjectEventType
	Name        string
	Description string
	CreatorId   int
	StartedAt   time.Time
	EndedAt     time.Time
	EventTime   time.Time
}
