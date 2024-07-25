package dto

import (
	"github.com/zeiss/knox/internal/controllers"
	openapi "github.com/zeiss/knox/pkg/apis"
)

// FromGetEnvironmentStateRequestObject ...
func FromGetEnvironmentStateRequestObject(request openapi.GetEnvironmentStateRequestObject) controllers.GetStateControllerQuery {
	return controllers.GetStateControllerQuery{
		Team:        request.TeamId,
		Project:     request.ProjectId,
		Environment: request.EnvironmentId,
	}
}

// ToGetEnvironmentStateResponseObject ...
func ToGetEnvironmentStateResponseObject(data map[string]interface{}) openapi.GetEnvironmentStateResponseObject {
	res := openapi.GetEnvironmentState200JSONResponse(data)

	return res
}

// FromUpdateEnvironmentStateRequestObject ...
func FromUpdateEnvironmentStateRequestObject(request openapi.UpdateEnvironmentStateRequestObject) controllers.UpdateStateControllerCommand {
	return controllers.UpdateStateControllerCommand{
		Team:        request.TeamId,
		Project:     request.ProjectId,
		Environment: request.EnvironmentId,
		State:       request.Body,
	}
}

// ToUpdateEnvironmentStateResponseObject ...
func ToUpdateEnvironmentStateResponseObject() openapi.UpdateEnvironmentState200JSONResponse {
	res := openapi.UpdateEnvironmentState200JSONResponse{}

	return res
}
