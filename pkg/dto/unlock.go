package dto

import (
	"github.com/zeiss/knox/internal/controllers"
	openapi "github.com/zeiss/knox/pkg/apis"
	"github.com/zeiss/knox/pkg/utils"
)

// FromUnlockEnvironmentRequestObject ...
func FromUnlockEnvironmentRequestObject(request openapi.UnlockEnvironmentRequestObject) controllers.UnlockControllerCommand {
	return controllers.UnlockControllerCommand{
		ID:          utils.PtrUUID(request.Body.Id),
		Team:        request.TeamId,
		Project:     request.ProjectId,
		Environment: request.EnvironmentId,
	}
}

// ToUnlockEnvironmentResponseObject ...
func ToUnlockEnvironmentResponseObject() openapi.UnlockEnvironmentResponseObject {
	res := openapi.UnlockEnvironment200Response{}

	return res
}
