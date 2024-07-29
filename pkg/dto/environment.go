package dto

import (
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/knox/internal/controllers"
	"github.com/zeiss/knox/internal/models"
	openapi "github.com/zeiss/knox/pkg/apis"
	"github.com/zeiss/knox/pkg/utils"
)

// FromCreateEnvironmentRequestObject ...
func FromCreateEnvironmentRequestObject(req openapi.CreateEnvironmentRequestObject) controllers.CreateEnvironmentCommand {
	return controllers.CreateEnvironmentCommand{
		TeamName:    req.TeamName,
		ProjectName: req.ProjectName,
		Name:        utils.PtrStr(req.Body.Name),
		Username:    utils.PtrStr(req.Body.Username),
		Password:    utils.PtrStr(req.Body.Secret),
	}
}

// ToCreateEnvironmentResponseObject ...
func ToCreateEnvironmentResponseObject() openapi.CreateEnvironment201JSONResponse {
	res := openapi.CreateEnvironment201JSONResponse{}

	return res
}

// FromGetEnvironmentsRequestObject ...
func FromGetEnvironmentsRequestObject(req openapi.GetEnvironmentsRequestObject) controllers.ListEnvironmentsQuery {
	return controllers.ListEnvironmentsQuery{
		TeamName:    req.TeamName,
		ProjectName: req.ProjectName,
		Limit:       utils.PtrInt(req.Params.Limit),
		Offset:      utils.PtrInt(req.Params.Offset),
	}
}

// ToGetEnvironmentsResponseObject ...
func ToGetEnvironmentsResponseObject(results tables.Results[models.Environment]) openapi.GetEnvironments200JSONResponse {
	res := openapi.GetEnvironments200JSONResponse{}

	environments := make([]openapi.Environment, results.GetLen())

	for i, environment := range results.GetRows() {
		environments[i] = openapi.Environment{
			Id:   utils.StrPtr(environment.ID.String()),
			Name: utils.StrPtr(environment.Name),
		}
	}

	res.Environments = &environments

	return res
}

// FromGetEnvironmentRequestObject ...
func FromGetEnvironmentRequestObject(req openapi.GetEnvironmentRequestObject) controllers.GetEnvironmentQuery {
	return controllers.GetEnvironmentQuery{
		TeamName:        req.TeamName,
		ProjectName:     req.ProjectName,
		EnvironmentName: req.EnvironmentName,
	}
}

// ToGetEnvironmentResponseObject ...
func ToGetEnvironmentResponseObject(environment models.Environment) openapi.GetEnvironment200JSONResponse {
	res := openapi.GetEnvironment200JSONResponse{
		Id:   utils.StrPtr(environment.ID.String()),
		Name: utils.StrPtr(environment.Name),
	}

	return res
}

// FromDeleteEnvironmentRequestObject ...
func FromDeleteEnvironmentRequestObject(req openapi.DeleteEnvironmentRequestObject) controllers.DeleteEnvironmentCommand {
	return controllers.DeleteEnvironmentCommand{
		TeamName:        req.TeamName,
		ProjectName:     req.ProjectName,
		EnvironmentName: req.EnvironmentName,
	}
}

// ToDeleteEnvironmentResponseObject ...
func ToDeleteEnvironmentResponseObject() openapi.DeleteEnvironment204Response {
	res := openapi.DeleteEnvironment204Response{}

	return res
}
