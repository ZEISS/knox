package authz

import (
	"context"

	"github.com/zeiss/knox/internal/ports"
	"github.com/zeiss/pkg/fga"

	openfga "github.com/openfga/go-sdk/client"
)

type writeTxImpl struct {
	client *openfga.OpenFgaClient
	store  fga.StoreTx
}

// NewWriteTx returns a new write transaction.
func NewWriteTx() fga.StoreTxFactory[ports.AuthzWriteTx] {
	return func(client *openfga.OpenFgaClient, storeTx fga.StoreTx) (ports.AuthzWriteTx, error) {
		return &writeTxImpl{client, storeTx}, nil
	}
}

// Add admin makes a user an admin.
func (tx *writeTxImpl) AddAdmin(ctx context.Context, user, team string) error {
	return tx.store.WriteTuple(
		ctx,
		fga.NewUser(fga.Namespace("user"), fga.String(user)),
		fga.NewObject(fga.Namespace("team"), fga.String(team)),
		fga.NewRelation(fga.String("admin")),
	)
}

// AddOwnerEnvironment creates owner of an environment.
func (tx *writeTxImpl) AddOwnerEnvironment(ctx context.Context, team, project, environment string) error {
	return tx.store.WriteTuple(
		ctx,
		fga.NewUser(fga.Namespace("project"), fga.Join(fga.DefaultSeparator, team, project)),
		fga.NewObject(fga.Namespace("environment"), fga.Join(fga.DefaultSeparator, team, project, environment)),
		fga.NewRelation(fga.String("owner")),
	)
}

// RemoveOwnerEnvironment removes owner of an environment.
func (tx *writeTxImpl) RemoveOwnerEnvironment(ctx context.Context, team, project, environment string) error {
	return tx.store.DeleteTuple(
		ctx,
		fga.NewUser(fga.Namespace("project"), fga.Join(fga.DefaultSeparator, team, project)),
		fga.NewObject(fga.Namespace("environment"), fga.Join(fga.DefaultSeparator, team, project, environment)),
		fga.NewRelation(fga.String("owner")),
	)
}
