package controllers

import (
	"context"

	"github.com/zeiss/knox/internal/ports"

	"github.com/go-playground/validator/v10"
	"github.com/zeiss/fiber-goth/adapters"
	seed "github.com/zeiss/gorm-seed"
)

var _ TeamController = (*TeamControllerImpl)(nil)

// CreateTeamCommand ...
type CreateTeamCommand struct {
	// Name is the name of the team.
	Name string `json:"name" form:"name"`
	// Description is the description of the team.
	Description string `json:"description" form:"description"`
	// Slug is the slug of the team.
	Slug string `json:"slug" form:"slug"`
}

// GetTeamQuery ...
type GetTeamQuery struct {
	// Slug is the slug of the team.
	Slug string `json:"slug" form:"slug"`
}

// DeleteTeamCommand ...
type DeleteTeamCommand struct {
	// Slug is the slug of the team.
	Slug string `json:"slug" form:"slug"`
}

// TeamController ...
type TeamController interface {
	// CreateTeam creates a new team.
	CreateTeam(ctx context.Context, cmd CreateTeamCommand) error
	// GetTeam gets a team.
	GetTeam(ctx context.Context, query GetTeamQuery) (adapters.GothTeam, error)
	// DeleteTeam deletes a team.
	DeleteTeam(ctx context.Context, cmd DeleteTeamCommand) error
}

// TeamControllerImpl is the controller for teams.
type TeamControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
}

// NewTeamController returns a new instance of TeamController.
func NewTeamController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *TeamControllerImpl {
	return &TeamControllerImpl{store}
}

// CreateTeam creates a new team.
func (c *TeamControllerImpl) CreateTeam(ctx context.Context, cmd CreateTeamCommand) error {
	validate = validator.New()

	err := validate.Struct(cmd)
	if err != nil {
		return err
	}

	team := adapters.GothTeam{
		Name:        cmd.Name,
		Description: cmd.Description,
		Slug:        cmd.Slug,
	}

	return c.store.ReadWriteTx(ctx, func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateTeam(ctx, &team)
	})
}

// GetTeam gets a team.
func (c *TeamControllerImpl) GetTeam(ctx context.Context, query GetTeamQuery) (adapters.GothTeam, error) {
	team := adapters.GothTeam{
		Slug: query.Slug,
	}

	err := c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetTeam(ctx, &team)
	})

	return team, err
}

// DeleteTeam deletes a team.
func (c *TeamControllerImpl) DeleteTeam(ctx context.Context, cmd DeleteTeamCommand) error {
	team := adapters.GothTeam{
		Slug: cmd.Slug,
	}

	return c.store.ReadWriteTx(ctx, func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteTeam(ctx, &team)
	})
}
