package authz

import (
	"context"

	"github.com/zeiss/knox/internal/ports"

	openfga "github.com/openfga/go-sdk/client"
	"github.com/zeiss/pkg/authz"
)

type writeTxImpl struct {
	client *openfga.OpenFgaClient
	store  authz.StoreTx
}

// NewWriteTx returns a new write transaction.
func NewWriteTx() authz.StoreTxFactory[ports.AuthzWriteTx] {
	return func(client *openfga.OpenFgaClient, storeTx authz.StoreTx) (ports.AuthzWriteTx, error) {
		return &writeTxImpl{client, storeTx}, nil
	}
}

// Add admin makes a user an admin.
func (tx *writeTxImpl) AddAdmin(ctx context.Context, user, team string) error {
	return tx.store.WriteTuple(
		ctx,
		authz.NewUser(authz.Namespace("user"), authz.String(user)),
		authz.NewObject(authz.Namespace("team"), authz.String(team)),
		authz.NewRelation(authz.String("admin")),
	)
}

// AddOwnerEnvironment creates owner of an environment.
func (tx *writeTxImpl) AddOwnerEnvironment(ctx context.Context, team, project, environment string) error {
	return tx.store.WriteTuple(
		ctx,
		authz.NewUser(authz.Namespace("project"), authz.Join(authz.DefaultSeparator, team, project)),
		authz.NewObject(authz.Namespace("environment"), authz.Join(authz.DefaultSeparator, team, project, environment)),
		authz.NewRelation(authz.String("owner")),
	)
}

// RemoveOwnerEnvironment removes owner of an environment.
func (tx *writeTxImpl) RemoveOwnerEnvironment(ctx context.Context, team, project, environment string) error {
	return tx.store.DeleteTuple(
		ctx,
		authz.NewUser(authz.Namespace("project"), authz.Join(authz.DefaultSeparator, team, project)),
		authz.NewObject(authz.Namespace("environment"), authz.Join(authz.DefaultSeparator, team, project, environment)),
		authz.NewRelation(authz.String("owner")),
	)
}
