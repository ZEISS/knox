package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/zeiss/knox/internal/controllers"
	openapi "github.com/zeiss/knox/pkg/apis"
	"github.com/zeiss/knox/pkg/dto"
)

var _ openapi.StrictServerInterface = (*apiHandlers)(nil)

type apiHandlers struct {
	locks controllers.LocksController
	state controllers.StateController
}

// NewAPIHandlers returns a new instance of APIHandlers.
func NewAPIHandlers(locks controllers.LocksController, state controllers.StateController) *apiHandlers {
	return &apiHandlers{locks, state}
}

// Get system health status
// (GET /_health)
func (a *apiHandlers) GetHealth(ctx context.Context, request openapi.GetHealthRequestObject) (openapi.GetHealthResponseObject, error) {
	return nil, nil
}

// Get system readiness
// (GET /_ready)
func (a *apiHandlers) GetReady(ctx context.Context, request openapi.GetReadyRequestObject) (openapi.GetReadyResponseObject, error) {
	return nil, nil
}

// Get a list of projects
// (GET /api/v1/project)
func (a *apiHandlers) GetApiV1Project(ctx context.Context, request openapi.GetApiV1ProjectRequestObject) (openapi.GetApiV1ProjectResponseObject, error) {
	return nil, nil
}

// Create a new project
// (POST /api/v1/project)
func (a *apiHandlers) PostApiV1Project(ctx context.Context, request openapi.PostApiV1ProjectRequestObject) (openapi.PostApiV1ProjectResponseObject, error) {
	return nil, nil
}

// Delete a project
// (DELETE /api/v1/project/{id})
func (a *apiHandlers) DeleteApiV1ProjectId(ctx context.Context, request openapi.DeleteApiV1ProjectIdRequestObject) (openapi.DeleteApiV1ProjectIdResponseObject, error) {
	return nil, nil
}

// Get a project
// (GET /api/v1/project/{id})
func (a *apiHandlers) GetApiV1ProjectId(ctx context.Context, request openapi.GetApiV1ProjectIdRequestObject) (openapi.GetApiV1ProjectIdResponseObject, error) {
	return nil, nil
}

// Update a project
// (PUT /api/v1/project/{id})
func (a *apiHandlers) PutApiV1ProjectId(ctx context.Context, request openapi.PutApiV1ProjectIdRequestObject) (openapi.PutApiV1ProjectIdResponseObject, error) {
	return nil, nil
}

// Get a list of environments
// (GET /api/v1/project/{projectId}/environment)
func (a *apiHandlers) GetApiV1ProjectProjectIdEnvironment(ctx context.Context, request openapi.GetApiV1ProjectProjectIdEnvironmentRequestObject) (openapi.GetApiV1ProjectProjectIdEnvironmentResponseObject, error) {
	return nil, nil
}

// Create a new environment
// (POST /api/v1/project/{projectId}/environment)
func (a *apiHandlers) PostApiV1ProjectProjectIdEnvironment(ctx context.Context, request openapi.PostApiV1ProjectProjectIdEnvironmentRequestObject) (openapi.PostApiV1ProjectProjectIdEnvironmentResponseObject, error) {
	return nil, nil
}

// Delete an environment
// (DELETE /api/v1/project/{projectId}/environment/{environmentId})
func (a *apiHandlers) DeleteApiV1ProjectProjectIdEnvironmentEnvironmentId(ctx context.Context, request openapi.DeleteApiV1ProjectProjectIdEnvironmentEnvironmentIdRequestObject) (openapi.DeleteApiV1ProjectProjectIdEnvironmentEnvironmentIdResponseObject, error) {
	return nil, nil
}

// Get an environment
// (GET /api/v1/project/{projectId}/environment/{environmentId})
func (a *apiHandlers) GetApiV1ProjectProjectIdEnvironmentEnvironmentId(ctx context.Context, request openapi.GetApiV1ProjectProjectIdEnvironmentEnvironmentIdRequestObject) (openapi.GetApiV1ProjectProjectIdEnvironmentEnvironmentIdResponseObject, error) {
	return nil, nil
}

// Update an environment
// (PUT /api/v1/project/{projectId}/environment/{environmentId})
func (a *apiHandlers) PutApiV1ProjectProjectIdEnvironmentEnvironmentId(ctx context.Context, request openapi.PutApiV1ProjectProjectIdEnvironmentEnvironmentIdRequestObject) (openapi.PutApiV1ProjectProjectIdEnvironmentEnvironmentIdResponseObject, error) {
	return nil, nil
}

// Get a list of snapshots
// (GET /api/v1/snapshot)
func (a *apiHandlers) GetApiV1Snapshot(ctx context.Context, request openapi.GetApiV1SnapshotRequestObject) (openapi.GetApiV1SnapshotResponseObject, error) {
	return nil, nil
}

// Create a new snapshot
// (POST /api/v1/snapshot)
func (a *apiHandlers) PostApiV1Snapshot(ctx context.Context, request openapi.PostApiV1SnapshotRequestObject) (openapi.PostApiV1SnapshotResponseObject, error) {
	return nil, nil
}

// Delete a snapshot
// (DELETE /api/v1/snapshot/{id})
func (a *apiHandlers) DeleteApiV1SnapshotId(ctx context.Context, request openapi.DeleteApiV1SnapshotIdRequestObject) (openapi.DeleteApiV1SnapshotIdResponseObject, error) {
	return nil, nil
}

// Get a snapshot
// (GET /api/v1/snapshot/{id})
func (a *apiHandlers) GetApiV1SnapshotId(ctx context.Context, request openapi.GetApiV1SnapshotIdRequestObject) (openapi.GetApiV1SnapshotIdResponseObject, error) {
	return nil, nil
}

// Update a snapshot
// (PUT /api/v1/snapshot/{id})
func (a *apiHandlers) PutApiV1SnapshotId(ctx context.Context, request openapi.PutApiV1SnapshotIdRequestObject) (openapi.PutApiV1SnapshotIdResponseObject, error) {
	return nil, nil
}

// Get a task
// (GET /api/v1/task/{id})
func (a *apiHandlers) GetApiV1TaskId(ctx context.Context, request openapi.GetApiV1TaskIdRequestObject) (openapi.GetApiV1TaskIdResponseObject, error) {
	return nil, nil
}

// Get a list of teams
// (GET /api/v1/team)
func (a *apiHandlers) GetApiV1Team(ctx context.Context, request openapi.GetApiV1TeamRequestObject) (openapi.GetApiV1TeamResponseObject, error) {
	return nil, nil
}

// Create a new team
// (POST /api/v1/team)
func (a *apiHandlers) PostApiV1Team(ctx context.Context, request openapi.PostApiV1TeamRequestObject) (openapi.PostApiV1TeamResponseObject, error) {
	return nil, nil
}

// Delete a team
// (DELETE /api/v1/team/{id})
func (a *apiHandlers) DeleteApiV1TeamId(ctx context.Context, request openapi.DeleteApiV1TeamIdRequestObject) (openapi.DeleteApiV1TeamIdResponseObject, error) {
	return nil, nil
}

// Get a team
// (GET /api/v1/team/{id})
func (a *apiHandlers) GetApiV1TeamId(ctx context.Context, request openapi.GetApiV1TeamIdRequestObject) (openapi.GetApiV1TeamIdResponseObject, error) {
	return nil, nil
}

// Update a team
// (PUT /api/v1/team/{id})
func (a *apiHandlers) PutApiV1TeamId(ctx context.Context, request openapi.PutApiV1TeamIdRequestObject) (openapi.PutApiV1TeamIdResponseObject, error) {
	return nil, nil
}

// Get a list of users
// (GET /api/v1/user)
func (a *apiHandlers) GetApiV1User(ctx context.Context, request openapi.GetApiV1UserRequestObject) (openapi.GetApiV1UserResponseObject, error) {
	return nil, nil
}

// Create a new user
// (POST /api/v1/user)
func (a *apiHandlers) PostApiV1User(ctx context.Context, request openapi.PostApiV1UserRequestObject) (openapi.PostApiV1UserResponseObject, error) {
	return nil, nil
}

// Delete a user
// (DELETE /api/v1/user/{id})
func (a *apiHandlers) DeleteApiV1UserId(ctx context.Context, request openapi.DeleteApiV1UserIdRequestObject) (openapi.DeleteApiV1UserIdResponseObject, error) {
	return nil, nil
}

// Get a user
// (GET /api/v1/user/{id})
func (a *apiHandlers) GetApiV1UserId(ctx context.Context, request openapi.GetApiV1UserIdRequestObject) (openapi.GetApiV1UserIdResponseObject, error) {
	return nil, nil
}

// Update a user
// (PUT /api/v1/user/{id})
func (a *apiHandlers) PutApiV1UserId(ctx context.Context, request openapi.PutApiV1UserIdRequestObject) (openapi.PutApiV1UserIdResponseObject, error) {
	return nil, nil
}

// Lock the state of Terraform environment
// (POST /client/{teamId}/{projectId}/{environmentId}/lock)
func (a *apiHandlers) LockEnvironment(ctx context.Context, request openapi.LockEnvironmentRequestObject) (openapi.LockEnvironmentResponseObject, error) {
	cmd := dto.FromLockEnvironmentRequestObject(request)

	err := a.locks.Lock(ctx, cmd)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return dto.ToLockEnvironmentResponseObject(), nil
}

// Get the state of Terraform environment
// (GET /client/{teamId}/{projectId}/{environmentId}/state)
func (a *apiHandlers) GetEnvironmentState(ctx context.Context, request openapi.GetEnvironmentStateRequestObject) (openapi.GetEnvironmentStateResponseObject, error) {
	query := dto.FromGetEnvironmentStateRequestObject(request)

	data, err := a.state.GetState(ctx, query)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return dto.ToGetEnvironmentStateResponseObject(data), nil
}

// Update the state of Terraform environment
// (POST /client/{teamId}/{projectId}/{environmentId}/state)
func (a *apiHandlers) UpdateEnvironmentState(ctx context.Context, request openapi.UpdateEnvironmentStateRequestObject) (openapi.UpdateEnvironmentStateResponseObject, error) {
	return nil, nil
}

// Unlock the state of Terraform environment
// (POST /client/{teamId}/{projectId}/{environmentId}/unlock)
func (a *apiHandlers) UnlockEnvironment(ctx context.Context, request openapi.UnlockEnvironmentRequestObject) (openapi.UnlockEnvironmentResponseObject, error) {
	cmd := dto.FromUnlockEnvironmentRequestObject(request)

	err := a.locks.Unlock(ctx, cmd)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return dto.ToUnlockEnvironmentResponseObject(), nil
}
