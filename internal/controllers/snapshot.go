package controllers

import (
	"context"

	"github.com/zeiss/knox/internal/models"
	"github.com/zeiss/knox/internal/ports"

	"github.com/google/uuid"
)

var _ SnapshotController = (*SnapshotControllerImpl)(nil)

// CreateSnapshotCommand ...
type CreateSnapshotCommand struct {
	Title       string
	Description string
	StateID     uuid.UUID
}

// SnapshotController ...
type SnapshotController interface {
	// CreateSnapshot ...
	CreateSnapshot(ctx context.Context, cmd CreateSnapshotCommand) (models.Snapshot, error)
}

// SnapshotControllerImpl ...
type SnapshotControllerImpl struct {
	store ports.Datastore
}

// NewSnapshotController ...
func NewSnapshotController(store ports.Datastore) *SnapshotControllerImpl {
	return &SnapshotControllerImpl{store}
}

// CreateSnapshot ...
func (c *SnapshotControllerImpl) CreateSnapshot(ctx context.Context, cmd CreateSnapshotCommand) (models.Snapshot, error) {
	state := models.State{
		ID: cmd.StateID,
	}

	err := c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetState(ctx, &state)
	})
	if err != nil {
		return models.Snapshot{}, err
	}

	snapshot := models.Snapshot{
		Title:         cmd.Title,
		Description:   &cmd.Description,
		TeamID:        state.TeamID,
		EnvironmentID: state.EnvironmentID,
		ProjectID:     state.ProjectID,
		Data:          state.Data,
	}

	err = c.store.ReadWriteTx(ctx, func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateSnapshot(ctx, &snapshot)
	})
	if err != nil {
		return models.Snapshot{}, err
	}

	return snapshot, nil
}
