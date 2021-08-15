package service

import (
	"projmng/pkg/code"
	"projmng/pkg/entity"
	"projmng/pkg/event"
)

type IProjectCommand interface {
	Create(proj event.Project) error
	Update(proj event.Project) error
	Delete(proj event.Project) error
}

type IProjectFinder interface {
	Get(id int) (entity.Project, error)
}

type ProjectService struct {
	projectCommand IProjectCommand
	projectUser    entity.IProjectUser
}

func (ps *ProjectService) CreateProject(evt event.Project) error {
	// verify user's auth
	// valicate whether project has a same name
	if err := ps.projectCommand.Create(evt); err != nil {
		return code.New(code.CreateProjectError, err)
	}
	// add self into project user group
	panic("")
}

func (ps *ProjectService) AddUsers(pu []*entity.ProjectUser) error {
	// notifiy users
	panic("")
}
