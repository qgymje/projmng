package service

import "projmng/pkg/entity"

type ProjectService struct {
	project     entity.IProject
	projectUser entity.IProjectUser
}

func (ps *ProjectService) CreateProject(project *entity.Project) error {
	// verify user's auth
	// valicate whether project has a same name
	// add self into project user group
	panic("")
}

func (ps *ProjectService) AddUsers(pu []*entity.ProjectUser) error {
	// notifiy users
	panic("")
}
