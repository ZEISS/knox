package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/zeiss/knox/internal/models"
	"github.com/zeiss/knox/internal/ports"

	"gorm.io/gorm"
)

// QueryError is an error that occurred while executing a query.
type QueryError struct {
	Query string
	Err   error
}

// Error ...
func (e *QueryError) Error() string { return e.Query + ": " + e.Err.Error() }

// Unwrap ...
func (e *QueryError) Unwrap() error { return e.Err }

type database struct {
	conn *gorm.DB
}

// NewDatastore returns a new instance of db.
func NewDB(conn *gorm.DB) (ports.Datastore, error) {
	return &database{
		conn: conn,
	}, nil
}

// Close closes the database connection.
func (d *database) Close() error {
	sqlDB, err := d.conn.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

// RunMigrations runs the database migrations.
func (d *database) Migrate(ctx context.Context) error {
	return d.conn.WithContext(ctx).AutoMigrate(
		&models.Environment{},
		&models.Project{},
		&models.Team{},
		&models.Lock{},
		&models.State{},
	)
}

// ReadWriteTx starts a read only transaction.
func (d *database) ReadWriteTx(ctx context.Context, fn func(context.Context, ports.ReadWriteTx) error) error {
	tx := d.conn.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := fn(ctx, &datastoreTx{tx}); err != nil {
		tx.Rollback()
	}

	if err := tx.Commit().Error; err != nil && !errors.Is(err, sql.ErrTxDone) {
		return err
	}

	return nil
}

// NewQueryError returns a new QueryError.
func NewQueryError(query string, err error) *QueryError {
	return &QueryError{
		Query: query,
		Err:   err,
	}
}

// ReadTx starts a read only transaction.
func (d *database) ReadTx(ctx context.Context, fn func(context.Context, ports.ReadTx) error) error {
	tx := d.conn.WithContext(ctx).Begin()
	if tx.Error != nil {
		return NewQueryError("begin read transaction", tx.Error)
	}

	err := fn(ctx, &datastoreTx{tx})
	if err != nil {
		tx.Rollback()
	}

	if err := tx.Commit().Error; err != nil && !errors.Is(err, sql.ErrTxDone) {
		return NewQueryError("commit read transaction", err)
	}

	if err != nil {
		return NewQueryError("commit read transaction", err)
	}

	return nil
}

var (
	_ ports.ReadTx      = (*datastoreTx)(nil)
	_ ports.ReadWriteTx = (*datastoreTx)(nil)
)

type datastoreTx struct {
	tx *gorm.DB
}

// CreateLock creates a new lock.
func (tx *datastoreTx) CreateLock(ctx context.Context, lock *models.Lock) error {
	return tx.tx.Create(lock).Error
}

// DeleteLock deletes a lock.
func (tx *datastoreTx) DeleteLock(ctx context.Context, lock *models.Lock) error {
	return tx.tx.Delete(lock).Error
}

// GetProject ...
func (tx *datastoreTx) GetProject(ctx context.Context, project *models.Project) error {
	return tx.tx.Where(project).First(project).Error
}

// GetEnvironment ...
func (tx *datastoreTx) GetEnvironment(ctx context.Context, environment *models.Environment) error {
	return tx.tx.Where(environment).First(environment).Error
}

// GetTeam ...
func (tx *datastoreTx) GetTeam(ctx context.Context, team *models.Team) error {
	return tx.tx.Where(team).First(team).Error
}

// GetState ...
func (tx *datastoreTx) GetState(ctx context.Context, state *models.State) error {
	return tx.tx.Where(state).First(state).Error
}

// UpdateState...
func (tx *datastoreTx) UpdateState(ctx context.Context, state *models.State) error {
	return tx.tx.Where(state).Save(state).Error
}
