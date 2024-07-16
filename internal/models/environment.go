package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Environment ...
type Environment struct {
	// ID is the primary key of the team.
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	// Name is the name of the environment.
	Name string `json:"name" gorm:"uniqueIndex:idx_project_environment"`
	// Project is the project of the environment.
	Project Project `json:"project" form:"project"`
	// ProjectID is the project id of the environment.
	ProjectID uuid.UUID `json:"project_id" gorm:"uniqueIndex:idx_project_environment"`
	// Username is the username of the environment.
	Username string `json:"username"`
	// Password is the password of the environment.
	Password string `json:"password"`
	// Description is the description of the environment.
	Description *string `json:"description" form:"description"`
	// CreatedAt is the time the team was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the team was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the team was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// HashPassword substitutes User.Password with its bcrypt hash
func (e *Environment) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	e.Password = string(hash)

	return nil
}

// ComparePassword compares User.Password hash with raw password
func (e *Environment) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(password))
}

// BeforeCreate gorm hook
func (e *Environment) BeforeCreate(db *gorm.DB) (err error) {
	return e.HashPassword()
}
