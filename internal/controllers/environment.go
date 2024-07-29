package controllers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/zeiss/fiber-htmx/components/tables"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/knox/internal/models"
	"github.com/zeiss/knox/internal/ports"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

var _ ProjectController = (*ProjectControllerImpl)(nil)

// CreateEnvironmentCommand ...
type CreateEnvironmentCommand struct {
	TeamName    string `json:"team_name" form:"team_name" validate:"required"`
	ProjectName string `json:"project_name" form:"project_name" validate:"required"`
	Name        string `json:"name" form:"name" validate:"required,min=1,max=255,alphanum,lowercase"`
	Username    string `json:"username" form:"username" validate:"required,min=1,max=255"`
	Password    string `json:"password" form:"password" validate:"required,min=1,max=255"`
}

// ListEnvironmentsQuery ...
type ListEnvironmentsQuery struct {
	TeamName    string `json:"team_name" form:"team_name" validate:"required"`
	ProjectName string `json:"project_name" form:"project_name" validate:"required"`
	Limit       int    `json:"limit" form:"limit" validate:"omitempty,min=1,max=100"`
	Offset      int    `json:"offset" form:"offset" validate:"omitempty,min=0"`
}

// EnvironmentControllerImpl ...
type EnvironmentControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
}

// EnvironmentController ...
type EnvironmentController interface {
	// CreateEnvironment ...
	CreateEnvironment(ctx context.Context, cmd CreateEnvironmentCommand) error
	// ListEnvironments ...
	ListEnvironments(ctx context.Context, query ListEnvironmentsQuery) (tables.Results[models.Environment], error)
}

// NewEnvironmentController ...
func NewEnvironmentController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *EnvironmentControllerImpl {
	return &EnvironmentControllerImpl{store}
}

// CreateEnvironment ...
func (c *EnvironmentControllerImpl) CreateEnvironment(ctx context.Context, cmd CreateEnvironmentCommand) error {
	validate = validator.New()

	err := validate.Struct(cmd)
	if err != nil {
		return err
	}

	project := models.Project{
		Name: cmd.ProjectName,
	}

	err = c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetProject(ctx, &project)
	})
	if err != nil {
		return err
	}

	environment := models.Environment{
		ProjectID: project.ID,
		Name:      cmd.Name,
		Username:  cmd.Username,
		Password:  cmd.Password,
	}

	return c.store.ReadWriteTx(ctx, func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateEnvironment(ctx, &environment)
	})
}

// ListEnvironments ...
func (c *EnvironmentControllerImpl) ListEnvironments(ctx context.Context, query ListEnvironmentsQuery) (tables.Results[models.Environment], error) {
	validate = validator.New()

	results := tables.Results[models.Environment]{
		Limit:  query.Limit,
		Offset: query.Offset,
	}

	err := validate.Struct(query)
	if err != nil {
		return results, err
	}

	err = c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListEnvironments(ctx, query.TeamName, query.ProjectName, &results)
	})
	if err != nil {
		return results, err
	}

	return results, nil
}
