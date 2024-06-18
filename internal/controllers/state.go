package controllers

import (
	"context"
	"encoding/json"

	"github.com/zeiss/knox/internal/models"
	"github.com/zeiss/knox/internal/ports"
	"gorm.io/datatypes"
)

var _ StateController = (*StateControllerImpl)(nil)

// GetStateControllerQuery ...
type GetStateControllerQuery struct {
	// Team is the team of the lock.
	Team string `json:"team" form:"team"`
	// Project is the project of the lock.
	Project string `json:"project" form:"project"`
	// Environment is the environment of the lock.
	Environment string `json:"environment" form:"environment"`
}

// UpdateStateControllerCommand ...
type UpdateStateControllerCommand struct {
	// Team is the team of the lock.
	Team string `json:"team" form:"team"`
	// Project is the project of the lock.
	Project string `json:"project" form:"project"`
	// Environment is the environment of the lock.
	Environment string `json:"environment" form:"environment"`
	// State is the state of the lock.
	State *map[string]interface{} `json:"state" form:"state"`
}

// StateController ...
type StateController interface {
	// GetState ...
	GetState(ctx context.Context, query GetStateControllerQuery) ([]byte, error)
	// UpdateState ...
	UpdateState(ctx context.Context, cmd UpdateStateControllerCommand) error
}

// StateControllerImpl is the controller for the state.
type StateControllerImpl struct {
	store ports.Datastore
}

// NewStateController returns a new LocksControllerImpl.
func NewStateController(store ports.Datastore) *StateControllerImpl {
	return &StateControllerImpl{store}
}

// GetState ...
func (c *StateControllerImpl) GetState(ctx context.Context, query GetStateControllerQuery) ([]byte, error) {
	var data []byte

	team := models.Team{
		Slug: query.Team,
	}

	err := c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetTeam(ctx, &team)
	})
	if err != nil {
		return data, err
	}

	project := models.Project{
		Name:   query.Project,
		TeamID: team.ID,
	}

	err = c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetProject(ctx, &project)
	})
	if err != nil {
		return data, err
	}

	env := models.Environment{
		Name:      query.Environment,
		ProjectID: project.ID,
	}

	err = c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetEnvironment(ctx, &env)
	})
	if err != nil {
		return data, err
	}

	state := models.State{
		TeamID:        team.ID,
		ProjectID:     project.ID,
		EnvironmentID: env.ID,
	}
	err = c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetState(ctx, &state)
	})
	if err != nil {
		return data, err
	}

	return state.Data, nil
}

// UpdateState ...
func (c *StateControllerImpl) UpdateState(ctx context.Context, cmd UpdateStateControllerCommand) error {
	team := models.Team{
		Slug: cmd.Team,
	}

	err := c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetTeam(ctx, &team)
	})
	if err != nil {
		return err
	}

	project := models.Project{
		Name:   cmd.Project,
		TeamID: team.ID,
	}

	err = c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetProject(ctx, &project)
	})
	if err != nil {
		return err
	}

	env := models.Environment{
		Name:      cmd.Environment,
		ProjectID: project.ID,
	}

	err = c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetEnvironment(ctx, &env)
	})
	if err != nil {
		return err
	}

	b, err := json.Marshal(cmd.State)
	if err != nil {
		return err
	}

	state := models.State{
		TeamID:        team.ID,
		ProjectID:     project.ID,
		EnvironmentID: env.ID,
		Data:          datatypes.JSON(b),
	}
	err = c.store.ReadWriteTx(ctx, func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.UpdateState(ctx, &state)
	})
	if err != nil {
		return err
	}

	return nil
}