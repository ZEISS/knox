package authz

import (
	"context"

	"github.com/zeiss/knox/internal/ports"
	authx "github.com/zeiss/pkg/authx/fga"

	openfga "github.com/openfga/go-sdk/client"
)

type writeTxImpl struct {
	client *openfga.OpenFgaClient
	store  authx.StoreTx
}

// NewWriteTx returns a new write transaction.
func NewWriteTx() authx.StoreTxFactory[ports.AuthzWriteTx] {
	return func(client *openfga.OpenFgaClient, storeTx authx.StoreTx) (ports.AuthzWriteTx, error) {
		return &writeTxImpl{client, storeTx}, nil
	}
}

// Add admin makes a user an admin.
func (tx *writeTxImpl) AddAdmin(ctx context.Context, user, team string) error {
	return tx.store.WriteTuple(
		ctx,
		authx.NewUser(authx.Namespace("user"), authx.String(user)),
		authx.NewObject(authx.Namespace("team"), authx.String(team)),
		authx.NewRelation(authx.String("admin")),
	)
}

// AddOwnerEnvironment creates owner of an environment.
func (tx *writeTxImpl) AddOwnerEnvironment(ctx context.Context, team, project, environment string) error {
	return tx.store.WriteTuple(
		ctx,
		authx.NewUser(authx.Namespace("project"), authx.Join(authx.DefaultSeparator, team, project)),
		authx.NewObject(authx.Namespace("environment"), authx.Join(authx.DefaultSeparator, team, project, environment)),
		authx.NewRelation(authx.String("owner")),
	)
}

// RemoveOwnerEnvironment removes owner of an environment.
func (tx *writeTxImpl) RemoveOwnerEnvironment(ctx context.Context, team, project, environment string) error {
	return tx.store.DeleteTuple(
		ctx,
		authx.NewUser(authx.Namespace("project"), authx.Join(authx.DefaultSeparator, team, project)),
		authx.NewObject(authx.Namespace("environment"), authx.Join(authx.DefaultSeparator, team, project, environment)),
		authx.NewRelation(authx.String("owner")),
	)
}
