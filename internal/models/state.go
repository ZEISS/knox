package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-goth/adapters"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// State ...
type State struct {
	// ID is the primary key of the state.
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" params:"id"`
	// Team is the team of the state.
	Team adapters.GothTeam `json:"team" form:"team"`
	// TeamID is the team id of the state.
	TeamID uuid.UUID `json:"team_id" gorm:"uniqueIndex:idx_team_project_environment_version"`
	// Environment is the environment of the state.
	Environment Environment `json:"environment" form:"environment"`
	// EnvironmentID is the environment id of the state.
	EnvironmentID uuid.UUID `json:"environment_id" gorm:"uniqueIndex:idx_team_project_environment_version"`
	// Project is the project of the state.
	Project Project `json:"project" form:"project"`
	// ProjectID is the project id of the state.
	ProjectID uuid.UUID `json:"project_id" gorm:"uniqueIndex:idx_team_project_environment_version"`
	// Data is the data of the state.
	Data datatypes.JSON `json:"data" form:"data"`
	// Version is the version of the state.
	Version uint `json:"version" gorm:"type:integer;default:1;uniqueIndex:idx_team_project_environment_version" params:"version"`
	// CreatedAt is the time the state was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the state was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the state was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
