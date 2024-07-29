package dto

import (
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/knox/internal/controllers"
	"github.com/zeiss/knox/internal/models"
	openapi "github.com/zeiss/knox/pkg/apis"
	"github.com/zeiss/knox/pkg/utils"
)

// FromCreateTeamRequestObject ...
func FromCreateTeamRequestObject(request openapi.CreateTeamRequestObject) controllers.CreateTeamCommand {
	return controllers.CreateTeamCommand{
		Name:        utils.PtrStr(request.Body.Name),
		Description: utils.PtrStr(request.Body.Description),
	}
}

// ToCreateTeamResponseObject ...
func ToCreateTeamResponseObject() openapi.CreateTeamResponseObject {
	res := openapi.CreateTeam201JSONResponse{}

	return res
}

// FromGetTeamRequestObject ...
func FromGetTeamRequestObject(request openapi.GetTeamRequestObject) controllers.GetTeamQuery {
	return controllers.GetTeamQuery{
		TeamName: request.TeamName,
	}
}

// ToGetTeamResponseObject ...
func ToGetTeamResponseObject(team models.Team) openapi.GetTeamResponseObject {
	res := openapi.GetTeam200JSONResponse(openapi.Team{
		Id:          utils.StrPtr(team.ID.String()),
		Name:        utils.StrPtr(team.Name),
		Description: utils.StrPtr(team.Description),
	})

	return res
}

// FromDeleteTeamRequestObject ...
func FromDeleteTeamRequestObject(request openapi.DeleteTeamRequestObject) controllers.DeleteTeamCommand {
	return controllers.DeleteTeamCommand{
		TeamName: request.TeamName,
	}
}

// ToDeleteTeamResponseObject ...
func ToDeleteTeamResponseObject() openapi.DeleteTeamResponseObject {
	res := openapi.DeleteTeam204Response{}

	return res
}

// FromGetTeamsRequestObject ...
func FromGetTeamsRequestObject(request openapi.GetTeamsRequestObject) controllers.ListTeamsQuery {
	return controllers.ListTeamsQuery{
		Limit:  utils.PtrInt(request.Params.Limit),
		Offset: utils.PtrInt(request.Params.Offset),
	}
}

// ToGetTeamsResponseObject ...
func ToGetTeamsResponseObject(results tables.Results[models.Team]) openapi.GetTeamsResponseObject {
	res := openapi.GetTeams200JSONResponse{}

	teams := []openapi.Team{}
	for _, team := range results.Rows {
		teams = append(teams, openapi.Team{
			Id:          utils.StrPtr(team.ID.String()),
			Name:        utils.StrPtr(team.Name),
			Description: utils.StrPtr(team.Description),
		})
	}
	res.Teams = &teams

	return res
}
