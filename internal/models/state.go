package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// State ...
type State struct {
	// ID is the primary key of the state.
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" params:"id"`
	// Team is the team of the state.
	Team Team `json:"team" form:"team"`
	// TeamID is the team id of the state.
	TeamID uuid.UUID `json:"team_id" gorm:"uniqueIndex:idx_team_project_environment"`
	// Environment is the environment of the state.
	Environment Environment `json:"environment" form:"environment"`
	// EnvironmentID is the environment id of the state.
	EnvironmentID uuid.UUID `json:"environment_id" gorm:"uniqueIndex:idx_team_project_environment"`
	// Project is the project of the state.
	Project Project `json:"project" form:"project"`
	// ProjectID is the project id of the state.
	ProjectID uuid.UUID `json:"project_id" gorm:"uniqueIndex:idx_team_project_environment"`
	// Data is the data of the state.
	Data datatypes.JSON `json:"data" form:"data"`
}
