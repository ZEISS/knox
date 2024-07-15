package dto

import (
	"github.com/zeiss/knox/internal/controllers"
	openapi "github.com/zeiss/knox/pkg/apis"
	"github.com/zeiss/knox/pkg/utils"
)

// FromCreateProjectRequestObject ...
func FromCreateProjectRequestObject(req openapi.CreateProjectRequestObject) controllers.CreateProjectCommand {
	return controllers.CreateProjectCommand{
		Team:        req.TeamId,
		Name:        utils.PtrStr(req.Body.Name),
		Description: utils.PtrStr(req.Body.Description),
	}
}

// ToCreateProjectResponseObject ...
func ToCreateProjectResponseObject() openapi.CreateProject201JSONResponse {
	res := openapi.CreateProject201JSONResponse{}

	return res
}
