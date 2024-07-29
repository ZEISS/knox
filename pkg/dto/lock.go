package dto

import (
	"github.com/zeiss/knox/internal/controllers"
	openapi "github.com/zeiss/knox/pkg/apis"
	"github.com/zeiss/knox/pkg/utils"
)

// FromLockEnvironmentRequestObject ...
func FromLockEnvironmentRequestObject(request openapi.LockEnvironmentRequestObject) controllers.LockControllerCommand {
	return controllers.LockControllerCommand{
		ID:          utils.PtrUUID(request.Body.Id),
		Team:        request.TeamName,
		Project:     request.ProjectName,
		Environment: request.EnvironmentName,
		Info:        utils.PtrStr(request.Body.Info),
	}
}

// ToLockEnvironmentResponseObject ...
func ToLockEnvironmentResponseObject() openapi.LockEnvironmentResponseObject {
	res := openapi.LockEnvironment200JSONResponse{}

	return res
}
