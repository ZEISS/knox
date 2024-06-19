package models

import (
	"time"

	"github.com/google/uuid"

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
	// State is the state of the snapshot.
	State State `json:"state" form:"state"`
	// StateID is the state id of the snapshot.
	StateID uuid.UUID `json:"state_id" gorm:"uniqueIndex:idx_state_snapshot"`
	// Data is the data of the state.
	Data datatypes.JSON `json:"data" form:"data"`
	// CreatedAt is the time the snapshot was created.
	CreatedAt time.Time
	// UpdatedAt is the time the snapshot was last updated.
	UpdatedAt time.Time
	// DeletedAt is the time the snapshot was deleted.
	DeletedAt gorm.DeletedAt
}
