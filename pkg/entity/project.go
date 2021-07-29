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

type IProject interface {
	Create(proj *Project) error
	Update(proj *Project) error
	Delete(proj *Project) error
	Get(id int) (*Project, error)
}
