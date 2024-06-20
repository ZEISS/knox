package handlers

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/zeiss/knox/internal/controllers"
	openapi "github.com/zeiss/knox/pkg/apis"
	"github.com/zeiss/knox/pkg/dto"
	"gorm.io/gorm"
)

var _ openapi.StrictServerInterface = (*apiHandlers)(nil)

type apiHandlers struct {
	locks     controllers.LocksController
	state     controllers.StateController
	snapshots controllers.SnapshotController
}

// NewAPIHandlers returns a new instance of APIHandlers.
func NewAPIHandlers(locks controllers.LocksController, state controllers.StateController, snapshots controllers.SnapshotController) *apiHandlers {
	return &apiHandlers{locks, state, snapshots}
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
// (GET /project)
func (a *apiHandlers) GetProject(ctx context.Context, request openapi.GetProjectRequestObject) (openapi.GetProjectResponseObject, error) {
	return nil, nil
}

// Create a new project
// (POST /project)
func (a *apiHandlers) PostProject(ctx context.Context, request openapi.PostProjectRequestObject) (openapi.PostProjectResponseObject, error) {
	return nil, nil
}

// Delete a project
// (DELETE /project/{id})
func (a *apiHandlers) DeleteProjectId(ctx context.Context, request openapi.DeleteProjectIdRequestObject) (openapi.DeleteProjectIdResponseObject, error) {
	return nil, nil
}

// Get a project
// (GET /project/{id})
func (a *apiHandlers) GetProjectId(ctx context.Context, request openapi.GetProjectIdRequestObject) (openapi.GetProjectIdResponseObject, error) {
	return nil, nil
}

// Update a project
// (PUT /project/{id})
func (a *apiHandlers) PutProjectId(ctx context.Context, request openapi.PutProjectIdRequestObject) (openapi.PutProjectIdResponseObject, error) {
	return nil, nil
}

// Get a list of environments
// (GET /project/{projectId}/environment)
func (a *apiHandlers) GetProjectProjectIdEnvironment(ctx context.Context, request openapi.GetProjectProjectIdEnvironmentRequestObject) (openapi.GetProjectProjectIdEnvironmentResponseObject, error) {
	return nil, nil
}

// Create a new environment
// (POST /project/{projectId}/environment)
func (a *apiHandlers) PostProjectProjectIdEnvironment(ctx context.Context, request openapi.PostProjectProjectIdEnvironmentRequestObject) (openapi.PostProjectProjectIdEnvironmentResponseObject, error) {
	return nil, nil
}

// Delete an environment
// (DELETE /project/{projectId}/environment/{environmentId})
func (a *apiHandlers) DeleteProjectProjectIdEnvironmentEnvironmentId(ctx context.Context, request openapi.DeleteProjectProjectIdEnvironmentEnvironmentIdRequestObject) (openapi.DeleteProjectProjectIdEnvironmentEnvironmentIdResponseObject, error) {
	return nil, nil
}

// Get an environment
// (GET /project/{projectId}/environment/{environmentId})
func (a *apiHandlers) GetProjectProjectIdEnvironmentEnvironmentId(ctx context.Context, request openapi.GetProjectProjectIdEnvironmentEnvironmentIdRequestObject) (openapi.GetProjectProjectIdEnvironmentEnvironmentIdResponseObject, error) {
	return nil, nil
}

// Update an environment
// (PUT /project/{projectId}/environment/{environmentId})
func (a *apiHandlers) PutProjectProjectIdEnvironmentEnvironmentId(ctx context.Context, request openapi.PutProjectProjectIdEnvironmentEnvironmentIdRequestObject) (openapi.PutProjectProjectIdEnvironmentEnvironmentIdResponseObject, error) {
	return nil, nil
}

// Get a list of snapshots
// (GET /snapshot)
func (a *apiHandlers) GetSnapshot(ctx context.Context, request openapi.GetSnapshotRequestObject) (openapi.GetSnapshotResponseObject, error) {
	return nil, nil
}

// Create a new snapshot
// (POST /snapshot)
func (a *apiHandlers) CreateSnapshot(ctx context.Context, request openapi.CreateSnapshotRequestObject) (openapi.CreateSnapshotResponseObject, error) {
	cmd := dto.FromCreateSnapshotRequestObject(request)

	snapshot, err := a.snapshots.CreateSnapshot(ctx, cmd)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return dto.ToCreateSnapshotResponseObject(snapshot), nil
}

// Delete a snapshot
// (DELETE /snapshot/{id})
func (a *apiHandlers) DeleteSnapshotId(ctx context.Context, request openapi.DeleteSnapshotIdRequestObject) (openapi.DeleteSnapshotIdResponseObject, error) {
	return nil, nil
}

// Get a snapshot
// (GET /snapshot/{id})
func (a *apiHandlers) GetSnapshotId(ctx context.Context, request openapi.GetSnapshotIdRequestObject) (openapi.GetSnapshotIdResponseObject, error) {
	return nil, nil
}

// Update a snapshot
// (PUT /snapshot/{id})
func (a *apiHandlers) PutSnapshotId(ctx context.Context, request openapi.PutSnapshotIdRequestObject) (openapi.PutSnapshotIdResponseObject, error) {
	return nil, nil
}

// Get a task
// (GET /task/{id})
func (a *apiHandlers) GetTaskId(ctx context.Context, request openapi.GetTaskIdRequestObject) (openapi.GetTaskIdResponseObject, error) {
	return nil, nil
}

// Get a list of teams
// (GET /team)
func (a *apiHandlers) GetTeam(ctx context.Context, request openapi.GetTeamRequestObject) (openapi.GetTeamResponseObject, error) {
	return nil, nil
}

// Create a new team
// (POST /team)
func (a *apiHandlers) PostTeam(ctx context.Context, request openapi.PostTeamRequestObject) (openapi.PostTeamResponseObject, error) {
	return nil, nil
}

// Delete a team
// (DELETE /team/{id})
func (a *apiHandlers) DeleteTeamId(ctx context.Context, request openapi.DeleteTeamIdRequestObject) (openapi.DeleteTeamIdResponseObject, error) {
	return nil, nil
}

// Get a team
// (GET /team/{id})
func (a *apiHandlers) GetTeamId(ctx context.Context, request openapi.GetTeamIdRequestObject) (openapi.GetTeamIdResponseObject, error) {
	return nil, nil
}

// Update a team
// (PUT /team/{id})
func (a *apiHandlers) PutTeamId(ctx context.Context, request openapi.PutTeamIdRequestObject) (openapi.PutTeamIdResponseObject, error) {
	return nil, nil
}

// Get a list of users
// (GET /user)
func (a *apiHandlers) GetUser(ctx context.Context, request openapi.GetUserRequestObject) (openapi.GetUserResponseObject, error) {
	return nil, nil
}

// Create a new user
// (POST /user)
func (a *apiHandlers) PostUser(ctx context.Context, request openapi.PostUserRequestObject) (openapi.PostUserResponseObject, error) {
	return nil, nil
}

// Delete a user
// (DELETE /user/{id})
func (a *apiHandlers) DeleteUserId(ctx context.Context, request openapi.DeleteUserIdRequestObject) (openapi.DeleteUserIdResponseObject, error) {
	return nil, nil
}

// Get a user
// (GET /user/{id})
func (a *apiHandlers) GetUserId(ctx context.Context, request openapi.GetUserIdRequestObject) (openapi.GetUserIdResponseObject, error) {
	return nil, nil
}

// Update a user
// (PUT /user/{id})
func (a *apiHandlers) PutUserId(ctx context.Context, request openapi.PutUserIdRequestObject) (openapi.PutUserIdResponseObject, error) {
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
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) { // the state was not found
		return openapi.GetEnvironmentState404JSONResponse{}, nil
	}

	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return dto.ToGetEnvironmentStateResponseObject(data), nil
}

// Update the state of Terraform environment
// (POST /client/{teamId}/{projectId}/{environmentId}/state)
func (a *apiHandlers) UpdateEnvironmentState(ctx context.Context, request openapi.UpdateEnvironmentStateRequestObject) (openapi.UpdateEnvironmentStateResponseObject, error) {
	cmd := dto.FromUpdateEnvironmentStateRequestObject(request)

	err := a.state.UpdateState(ctx, cmd)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return dto.ToUpdateEnvironmentStateResponseObject(), nil
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
