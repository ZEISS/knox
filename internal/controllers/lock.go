package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/knox/internal/models"
	"github.com/zeiss/knox/internal/ports"
	"github.com/zeiss/pkg/dbx"
)

var _ LocksController = (*LocksControllerImpl)(nil)

// LockControllerCommand ...
type LockControllerCommand struct {
	// ID is the ID of the lock.
	ID uuid.UUID `json:"id" form:"id"`
	// Team is the team of the lock.
	Team string `json:"team" form:"team"`
	// Project is the project of the lock.
	Project string `json:"project" form:"project"`
	// Environment is the environment of the lock.
	Environment string `json:"environment" form:"environment"`
	// Info is the info of the lock.
	Info string `json:"info" form:"info"`
	// Operation is the operation of the lock.
	Operation string `json:"operation" form:"operation"`
	// Path is the path of the lock.
	Path string `json:"path" form:"path"`
	// Version is the version of the lock.
	Version string `json:"version" form:"version"`
	// Who is the who of the lock.
	Who string `json:"who" form:"who"`
}

// UnlockControllerCommand ...
type UnlockControllerCommand struct {
	// ID is the ID of the lock.
	ID uuid.UUID `json:"id" form:"id"`
	// Team is the team of the lock.
	Team string `json:"team" form:"team"`
	// Project is the project of the lock.
	Project string `json:"project" form:"project"`
	// Environment is the environment of the lock.
	Environment string `json:"environment" form:"environment"`
}

// LocksController ...
type LocksController interface {
	// Lock ...
	Lock(ctx context.Context, cmd LockControllerCommand) error
	// Unlock ...
	Unlock(ctx context.Context, cmd UnlockControllerCommand) error
}

// LocksControllerImpl is the controller for operators.
type LocksControllerImpl struct {
	store dbx.Database[ports.ReadTx, ports.ReadWriteTx]
}

// NewLocksController returns a new LocksControllerImpl.
func NewLocksController(store dbx.Database[ports.ReadTx, ports.ReadWriteTx]) *LocksControllerImpl {
	return &LocksControllerImpl{store}
}

// Lock ...
func (c *LocksControllerImpl) Lock(ctx context.Context, cmd LockControllerCommand) error {
	team := models.Team{
		Name: cmd.Team,
	}

	err := c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
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

	l := models.Lock{}
	l.ID = cmd.ID
	l.EnvironmentID = env.ID

	return c.store.ReadWriteTx(ctx, func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateLock(ctx, &l)
	})
}

// Unlock ...
func (c *LocksControllerImpl) Unlock(ctx context.Context, cmd UnlockControllerCommand) error {
	l := models.Lock{
		ID: cmd.ID,
	}

	return c.store.ReadWriteTx(ctx, func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteLock(ctx, &l)
	})
}
