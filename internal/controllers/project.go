package controllers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/zeiss/fiber-goth/adapters"
	"github.com/zeiss/fiber-htmx/components/tables"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/knox/internal/models"
	"github.com/zeiss/knox/internal/ports"
	"github.com/zeiss/knox/pkg/utils"
)

var _ ProjectController = (*ProjectControllerImpl)(nil)

// CreateProjectCommand ...
type CreateProjectCommand struct {
	Team        string `json:"team" form:"team" validate:"required"`
	Name        string `json:"name" form:"name" validate:"required,min=1,max=255,alphanum,lowercase"`
	Description string `json:"description" form:"description" validate:"omitempty,min=3,max=2048"`
}

// ListProjectsQuery ...
type ListProjectsQuery struct {
	Team   string `json:"team" form:"team"`
	Limit  int    `json:"limit" form:"limit"`
	Offset int    `json:"offset" form:"offset"`
	Sort   string `json:"sort" form:"sort"`
}

// DeleteProjectCommand ...
type DeleteProjectCommand struct {
	Name string `json:"name" form:"name" validate:"required"`
}

// ProjectControllerImpl ...
type ProjectControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
}

// ProjectController ...
type ProjectController interface {
	// CreateProject ...
	CreateProject(ctx context.Context, cmd CreateProjectCommand) error
	// ListProjects ...
	ListProjects(ctx context.Context, cmd ListProjectsQuery) (tables.Results[models.Project], error)
	// DeleteProject ...
	DeleteProject(ctx context.Context, cmd DeleteProjectCommand) error
}

// NewProjectController ...
func NewProjectController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ProjectControllerImpl {
	return &ProjectControllerImpl{store}
}

// CreateProject ...
func (c *ProjectControllerImpl) CreateProject(ctx context.Context, cmd CreateProjectCommand) error {
	validate = validator.New()

	err := validate.Struct(cmd)
	if err != nil {
		return err
	}

	team := adapters.GothTeam{
		Slug: cmd.Team,
	}

	err = c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
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

// ListProjects ...
func (c *ProjectControllerImpl) ListProjects(ctx context.Context, cmd ListProjectsQuery) (tables.Results[models.Project], error) {
	teams := tables.Results[models.Project]{
		Limit:  cmd.Limit,
		Offset: cmd.Offset,
		Sort:   cmd.Sort,
	}

	err := c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListProjects(ctx, cmd.Team, &teams)
	})
	if err != nil {
		return teams, err
	}

	return teams, nil
}

// DeleteProject ...
func (c *ProjectControllerImpl) DeleteProject(ctx context.Context, cmd DeleteProjectCommand) error {
	err := validate.Struct(cmd)
	if err != nil {
		return err
	}

	project := models.Project{
		Name: cmd.Name,
	}

	return c.store.ReadWriteTx(ctx, func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteProject(ctx, &project)
	})
}