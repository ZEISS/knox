package dto

import (
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/knox/internal/controllers"
	"github.com/zeiss/knox/internal/models"
	openapi "github.com/zeiss/knox/pkg/apis"
	"github.com/zeiss/knox/pkg/utils"
)

// FromCreateProjectRequestObject ...
func FromCreateProjectRequestObject(req openapi.CreateProjectRequestObject) controllers.CreateProjectCommand {
	return controllers.CreateProjectCommand{
		TeamID:      req.TeamId,
		Name:        utils.PtrStr(req.Body.Name),
		Description: utils.PtrStr(req.Body.Description),
	}
}

// ToCreateProjectResponseObject ...
func ToCreateProjectResponseObject() openapi.CreateProject201JSONResponse {
	res := openapi.CreateProject201JSONResponse{}

	return res
}

// FromGetProjectsRequestObject ...
func FromGetProjectsRequestObject(req openapi.GetProjectsRequestObject) controllers.ListProjectsQuery {
	return controllers.ListProjectsQuery{
		Limit:  utils.PtrInt(req.Params.Limit),
		Offset: utils.PtrInt(req.Params.Offset),
		TeamID: req.TeamId,
	}
}

// ToGetProjectsResponseObject ...
func ToGetProjectsResponseObject(results tables.Results[models.Project]) openapi.GetProjects200JSONResponse {
	res := openapi.GetProjects200JSONResponse{}

	projects := make([]openapi.Project, results.GetLen())

	for i, project := range results.GetRows() {
		projects[i] = openapi.Project{
			Id:          utils.StrPtr(project.ID.String()),
			Name:        utils.StrPtr(project.Name),
			Description: project.Description,
		}
	}

	res.Projects = &projects

	return res
}

// FromDeleteProjectRequestObject ...
func FromDeleteProjectRequestObject(req openapi.DeleteProjectRequestObject) controllers.DeleteProjectCommand {
	return controllers.DeleteProjectCommand{
		ID: req.ProjectId,
	}
}

// ToDeleteProjectResponseObject ...
func ToDeleteProjectResponseObject() openapi.DeleteProject204Response {
	res := openapi.DeleteProject204Response{}

	return res
}
