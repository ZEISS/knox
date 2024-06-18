package ports

import (
	"context"
	"io"

	"github.com/zeiss/knox/internal/models"
)

// Migration is a method that runs the migration.
type Migration interface {
	// Migrate is a method that runs the migration.
	Migrate(context.Context) error
}

// Datastore provides methods for transactional operations.
type Datastore interface {
	// ReadTx starts a read only transaction.
	ReadTx(context.Context, func(context.Context, ReadTx) error) error
	// ReadWriteTx starts a read write transaction.
	ReadWriteTx(context.Context, func(context.Context, ReadWriteTx) error) error

	io.Closer
	Migration
}

// ReadTx provides methods for transactional read operations.
type ReadTx interface {
	// GetProject ...
	GetProject(context.Context, *models.Project) error
	// GetEnvironment ...
	GetEnvironment(context.Context, *models.Environment) error
	// GetTeam ...
	GetTeam(context.Context, *models.Team) error
	// GetState ...
	GetState(context.Context, *models.State) error
	// GetLock ...
	GetLock(context.Context, *models.Lock) error
}

// ReadWriteTx provides methods for transactional read and write operations.
type ReadWriteTx interface {
	// CreateLock creates a new lock.
	CreateLock(context.Context, *models.Lock) error
	// DeleteLock deletes a lock.
	DeleteLock(context.Context, *models.Lock) error
	// UpdateState creates a new state.
	UpdateState(context.Context, *models.State) error

	ReadTx
}
