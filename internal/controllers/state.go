package controllers

import (
	"context"
	"encoding/json"

	"github.com/zeiss/knox/internal/models"
	"github.com/zeiss/knox/internal/ports"
	openapi "github.com/zeiss/knox/pkg/apis"

	"github.com/google/uuid"
	seed "github.com/zeiss/gorm-seed"
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
	// LockID is the ID of the lock.
	LockID uuid.UUID `json:"lock_id" form:"lock_id"`
	// State is the state of the lock.
	State *map[string]interface{} `json:"state" form:"state"`
}

// StateController ...
type StateController interface {
	// GetState ...
	GetState(ctx context.Context, query GetStateControllerQuery) (map[string]interface{}, error)
	// UpdateState ...
	UpdateState(ctx context.Context, cmd UpdateStateControllerCommand) error
}

// StateControllerImpl is the controller for the state.
type StateControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
}

// NewStateController returns a new LocksControllerImpl.
func NewStateController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *StateControllerImpl {
	return &StateControllerImpl{store}
}

// GetState ...
func (c *StateControllerImpl) GetState(ctx context.Context, query GetStateControllerQuery) (map[string]interface{}, error) {
	var data map[string]interface{}

	team := models.Team{
		Name: query.Team,
	}

	err := c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetTeam(ctx, &team)
	})
	if err != nil {
		return data, err
	}

	project := models.Project{
		Name:    query.Project,
		OwnerID: team.ID,
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
		return tx.GetEnvironment(ctx, query.Team, query.Project, &env)
	})
	if err != nil {
		return data, err
	}

	state := models.State{
		ProjectID:     project.ID,
		EnvironmentID: env.ID,
	}
	err = c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetState(ctx, &state)
	})
	if err != nil {
		return data, err
	}

	var payload openapi.Payload
	err = json.Unmarshal(state.Data, &payload)
	if err != nil {
		return data, err
	}

	return payload, nil
}

// UpdateState ...
func (c *StateControllerImpl) UpdateState(ctx context.Context, cmd UpdateStateControllerCommand) error {
	lock := models.Lock{
		ID: cmd.LockID,
	}

	err := c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetLock(ctx, &lock)
	})
	if err != nil {
		return err
	}

	team := models.Team{
		Name: cmd.Team,
	}

	err = c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetTeam(ctx, &team)
	})
	if err != nil {
		return err
	}

	project := models.Project{
		Name:    cmd.Project,
		OwnerID: team.ID,
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
		return tx.GetEnvironment(ctx, cmd.Team, cmd.Project, &env)
	})
	if err != nil {
		return err
	}

	b, err := json.Marshal(cmd.State)
	if err != nil {
		return err
	}

	state := models.State{
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
