package controllers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/knox/internal/models"
	"github.com/zeiss/knox/internal/ports"
	"github.com/zeiss/knox/pkg/utils"
	"github.com/zeiss/pkg/dbx"
)

var _ ProjectController = (*ProjectControllerImpl)(nil)

// CreateProjectCommand ...
type CreateProjectCommand struct {
	TeamName    string `json:"team_name" form:"team_name" validate:"required"`
	Name        string `json:"name" form:"name" validate:"required,min=1,max=255,alphanum,lowercase"`
	Description string `json:"description" form:"description" validate:"omitempty,min=3,max=2048"`
}

// ListProjectsQuery ...
type ListProjectsQuery struct {
	TeamName string `json:"team_name" form:"team_name" validate:"required"`
	Limit    int    `json:"limit" form:"limit"`
	Offset   int    `json:"offset" form:"offset"`
	Sort     string `json:"sort" form:"sort"`
}

// GetProjectQuery ...
type GetProjectQuery struct {
	TeamName    string `json:"team_name" form:"team_name" validate:"required"`
	ProjectName string `json:"project_name" form:"project_name" validate:"required"`
}

// DeleteProjectCommand ...
type DeleteProjectCommand struct {
	TeamName    string `json:"team_name" form:"team_name" validate:"required"`
	ProjectName string `json:"project_name" form:"project_name" validate:"required"`
}

// ProjectControllerImpl ...
type ProjectControllerImpl struct {
	store dbx.Database[ports.ReadTx, ports.ReadWriteTx]
}

// ProjectController ...
type ProjectController interface {
	// CreateProject ...
	CreateProject(ctx context.Context, cmd CreateProjectCommand) error
	// GetProject ...
	GetProject(ctx context.Context, cmd GetProjectQuery) (models.Project, error)
	// ListProjects ...
	ListProjects(ctx context.Context, cmd ListProjectsQuery) (tables.Results[models.Project], error)
	// DeleteProject ...
	DeleteProject(ctx context.Context, cmd DeleteProjectCommand) error
}

// NewProjectController ...
func NewProjectController(store dbx.Database[ports.ReadTx, ports.ReadWriteTx]) *ProjectControllerImpl {
	return &ProjectControllerImpl{store}
}

// CreateProject ...
func (c *ProjectControllerImpl) CreateProject(ctx context.Context, cmd CreateProjectCommand) error {
	validate = validator.New()

	err := validate.Struct(&cmd)
	if err != nil {
		return err
	}

	team := models.Team{
		Name: cmd.TeamName,
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
		Owner:       team,
	}

	return c.store.ReadWriteTx(ctx, func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateProject(ctx, &project)
	})
}

// GetProject ...
func (c *ProjectControllerImpl) GetProject(ctx context.Context, cmd GetProjectQuery) (models.Project, error) {
	project := models.Project{
		Name: cmd.ProjectName,
	}

	err := c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetProject(ctx, &project)
	})
	if err != nil {
		return project, err
	}

	return project, nil
}

// ListProjects ...
func (c *ProjectControllerImpl) ListProjects(ctx context.Context, cmd ListProjectsQuery) (tables.Results[models.Project], error) {
	teams := tables.Results[models.Project]{
		Limit:  cmd.Limit,
		Offset: cmd.Offset,
		Sort:   cmd.Sort,
	}

	err := c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListProjects(ctx, cmd.TeamName, &teams)
	})
	if err != nil {
		return teams, err
	}

	return teams, nil
}

// DeleteProject ...
func (c *ProjectControllerImpl) DeleteProject(ctx context.Context, cmd DeleteProjectCommand) error {
	return c.store.ReadWriteTx(ctx, func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteProject(ctx, cmd.TeamName, &models.Project{Name: cmd.ProjectName})
	})
}
