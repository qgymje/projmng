package entity

import "time"

type ProjectUser struct {
	Id        int
	Project   *Project
	User      *User
	CreatedAt time.Time
	DeletedAt time.Time
}

type IProjectUser interface {
	Add(pu *ProjectUser) error
	Delete(pu *ProjectUser) error
	ListUser(proj *Project) ([]*User, error)
	ListByUser(uid int) ([]*Project, error)
}
