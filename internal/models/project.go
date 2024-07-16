package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-goth/adapters"
	"gorm.io/gorm"
)

// Project ...
type Project struct {
	// ID is the primary key of the team.
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	// Name is the name of the project.
	Name string `json:"name" gorm:"uniqueIndex:idx_team_name"`
	// Team is the team of the project.
	Team adapters.GothTeam `json:"team"`
	// TeamID is the team id of the project.
	TeamID uuid.UUID `json:"team_id" gorm:"uniqueIndex:idx_team_name"`
	// Description is the description of the project.
	Description *string `json:"description"`
	// Environments are the environments in the project.
	Environments []Environment `json:"environments" gorm:"foreignKey:ProjectID"`
	// CreatedAt is the time the team was created.
	CreatedAt time.Time
	// UpdatedAt is the time the team was last updated.
	UpdatedAt time.Time
	// DeletedAt is the time the team was deleted.
	DeletedAt gorm.DeletedAt
}
