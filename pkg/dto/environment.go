package dto

import (
	"github.com/zeiss/knox/internal/controllers"
	openapi "github.com/zeiss/knox/pkg/apis"
	"github.com/zeiss/knox/pkg/utils"
)

// FromCreateEnvironmentRequestObject ...
func FromCreateEnvironmentRequestObject(req openapi.CreateEnvironmentRequestObject) controllers.CreateEnvironmentCommand {
	return controllers.CreateEnvironmentCommand{
		Team:     req.TeamId,
		Project:  req.ProjectId,
		Name:     utils.PtrStr(req.Body.Name),
		Username: utils.PtrStr(req.Body.Username),
		Password: utils.PtrStr(req.Body.Secret),
	}
}

// ToCreateEnvironmentResponseObject ...
func ToCreateEnvironmentResponseObject() openapi.CreateEnvironment201JSONResponse {
	res := openapi.CreateEnvironment201JSONResponse{}

	return res
}
