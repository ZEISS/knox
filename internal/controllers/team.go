package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/knox/internal/models"
	"github.com/zeiss/knox/internal/ports"

	"github.com/go-playground/validator/v10"
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
	// ID is the ID of the team.
	ID uuid.UUID `json:"id" form:"id"`
}

// ListTeamsQuery ...
type ListTeamsQuery struct {
	Limit  int `json:"limit" form:"limit"`
	Offset int `json:"offset" form:"offset"`
}

// DeleteTeamCommand ...
type DeleteTeamCommand struct {
	// ID is the ID of the team.
	ID uuid.UUID `json:"id" form:"id"`
}

// TeamController ...
type TeamController interface {
	// CreateTeam creates a new team.
	CreateTeam(ctx context.Context, cmd CreateTeamCommand) error
	// GetTeam gets a team.
	GetTeam(ctx context.Context, query GetTeamQuery) (models.Team, error)
	// DeleteTeam deletes a team.
	DeleteTeam(ctx context.Context, cmd DeleteTeamCommand) error
	// ListTeams lists teams.
	ListTeams(ctx context.Context, query ListTeamsQuery) (tables.Results[models.Team], error)
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

	team := models.Team{
		Name:        cmd.Name,
		Description: cmd.Description,
	}

	return c.store.ReadWriteTx(ctx, func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateTeam(ctx, &team)
	})
}

// GetTeam gets a team.
func (c *TeamControllerImpl) GetTeam(ctx context.Context, query GetTeamQuery) (models.Team, error) {
	team := models.Team{
		ID: query.ID,
	}

	err := c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetTeam(ctx, &team)
	})

	return team, err
}

// DeleteTeam deletes a team.
func (c *TeamControllerImpl) DeleteTeam(ctx context.Context, cmd DeleteTeamCommand) error {
	return c.store.ReadWriteTx(ctx, func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteTeam(ctx, &models.Team{ID: cmd.ID})
	})
}

// ListTeams lists teams.
func (c *TeamControllerImpl) ListTeams(ctx context.Context, query ListTeamsQuery) (tables.Results[models.Team], error) {
	teams := tables.Results[models.Team]{}

	err := c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListTeams(ctx, &teams)
	})

	return teams, err
}
