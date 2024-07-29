package models

import (
	"time"

	"github.com/google/uuid"
)

// Lock is a lock for a project.
type Lock struct {
	// ID is the primary key of the lock.
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" params:"id"`
	// Info is the info of the lock.
	Info string `json:"info" form:"info"`
	// Operation is the operation of the lock.
	Operation string `json:"operation" form:"operation"`
	// Path is the path of the lock.
	Path string `json:"path" form:"path"`
	// Version is the version of the lock.
	Version string `json:"version" form:"version"`
	// Who is the who of the lock.
	Who string `json:"who" form:"who"`
	// Environment is the environment of the lock.
	Environment Environment `json:"environment" form:"environment"`
	// EnvironmentID is the environment id of the lock.
	EnvironmentID uuid.UUID `json:"environment_id" gorm:"uniqueIndex"`
	// CreatedAt is the time the team was created.
	CreatedAt time.Time
	// UpdatedAt is the time the team was last updated.
	UpdatedAt time.Time
}
