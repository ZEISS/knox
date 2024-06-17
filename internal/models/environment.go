package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Environment ...
type Environment struct {
	// ID is the primary key of the team.
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" params:"id"`
	// Name is the name of the environment.
	Name string `json:"name" form:"name" validate:"required,alphanum,gt=3,lt=255"`
	// Description is the description of the environment.
	Description *string `json:"description" form:"description" validate:"omitempty,max=1024"`
	// Project is the project of the environment.
	Project Project `json:"project" form:"project"`
	// ProjectID is the project id of the environment.
	ProjectID uuid.UUID `json:"project_id"`
	// CreatedAt is the time the team was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the team was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the team was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
