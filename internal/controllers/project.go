package controllers

import (
	"context"

	"github.com/zeiss/fiber-goth/adapters"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/knox/internal/models"
	"github.com/zeiss/knox/internal/ports"
	"github.com/zeiss/knox/pkg/utils"
)

var _ ProjectController = (*ProjectControllerImpl)(nil)

// CreateProjectCommand ...
type CreateProjectCommand struct {
	Team        string `json:"team" form:"team"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

// ProjectControllerImpl ...
type ProjectControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
}

// ProjectController ...
type ProjectController interface {
	// CreateProject ...
	CreateProject(ctx context.Context, cmd CreateProjectCommand) error
}

// NewProjectController ...
func NewProjectController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ProjectControllerImpl {
	return &ProjectControllerImpl{store}
}

// CreateProject ...
func (c *ProjectControllerImpl) CreateProject(ctx context.Context, cmd CreateProjectCommand) error {
	team := adapters.GothTeam{
		Slug: cmd.Team,
	}

	err := c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetTeam(ctx, &team)
	})
	if err != nil {
		return err
	}

	project := models.Project{
		Name:        cmd.Name,
		Description: utils.StrPtr(cmd.Description),
		Team:        team,
	}

	return c.store.ReadWriteTx(ctx, func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateProject(ctx, &project)
	})
}
