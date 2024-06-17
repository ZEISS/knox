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
func ToGetEnvironmentStateResponseObject(data []byte) openapi.GetEnvironmentStateResponseObject {
	res := openapi.GetEnvironmentState200JSONResponse{}

	return res
}
