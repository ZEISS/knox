package ports

import (
	"context"
)

// AuthzWriteTx provides methods for transactional read and write operations.
type AuthzWriteTx interface {
	// AddAdmin is a method that makes a user an admin.
	AddAdmin(ctx context.Context, user, team string) error
	// AddOwnerEnvironment is a method that creates owner of an environment.
	AddOwnerEnvironment(ctx context.Context, team, project, environment string) error
	// RemoveOwnerEnvironment is a method that removes owner of an environment.
	RemoveOwnerEnvironment(ctx context.Context, team, project, environment string) error
}
