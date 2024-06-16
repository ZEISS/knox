package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/zeiss/knox/internal/ports"

	"gorm.io/gorm"
)

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
	return d.conn.WithContext(ctx).AutoMigrate()
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

	if err := tx.Error; err != nil && !errors.Is(err, sql.ErrTxDone) {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

// ReadTx starts a read only transaction.
func (d *database) ReadTx(ctx context.Context, fn func(context.Context, ports.ReadTx) error) error {
	tx := d.conn.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := fn(ctx, &datastoreTx{tx}); err != nil {
		tx.Rollback()
	}

	if err := tx.Error; err != nil && !errors.Is(err, sql.ErrTxDone) {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
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
