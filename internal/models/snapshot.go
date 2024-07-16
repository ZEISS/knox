package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-goth/adapters"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Snapshot ...
type Snapshot struct {
	// ID is the primary key of the snapshot.
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" params:"id"`
	// Title is the name of the snapshot.
	Title string `json:"name" form:"name" validate:"required,alphanum,gt=3,lt=255"`
	// Description is the description of the snapshot.
	Description *string `json:"description" form:"description" validate:"omitempty,max=1024"`
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
	// Version is the version of the state.
	Version int `json:"version" gorm:"uniqueIndex:idx_team_project_environment_version"`
	// Data is the data of the state.
	Data datatypes.JSON `json:"data" form:"data"`
	// CreatedAt is the time the snapshot was created.
	CreatedAt time.Time
	// UpdatedAt is the time the snapshot was last updated.
	UpdatedAt time.Time
	// DeletedAt is the time the snapshot was deleted.
	DeletedAt gorm.DeletedAt
}
