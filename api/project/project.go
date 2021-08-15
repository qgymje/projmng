package project

import (
	"projmng/pkg/service"

	"github.com/gin-gonic/gin"
)

type ProjectAPI struct {
	projectService *service.ProjectService
}

func (api *ProjectAPI) Create(*gin.Context) {
}
