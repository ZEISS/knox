package controllers

import (
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/knox/internal/ports"
)

var _ TeamController = (*TeamControllerImpl)(nil)

// TeamController ...
type TeamController interface{}

// TeamControllerImpl is the controller for teams.
type TeamControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
}

// NewTeamController returns a new instance of TeamController.
func NewTeamController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *TeamControllerImpl {
	return &TeamControllerImpl{store}
}
