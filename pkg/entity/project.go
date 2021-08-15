package entity

import "time"

type Project struct {
	Id          int
	Name        string
	Description string
	Creator     *User
	StartedAt   time.Time
	EndedAt     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
