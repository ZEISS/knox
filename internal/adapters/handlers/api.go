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
	locks       controllers.LocksController
	state       controllers.StateController
	snapshots   controllers.SnapshotController
	team        controllers.TeamController
	project     controllers.ProjectController
	environment controllers.EnvironmentController
}

// NewAPIHandlers returns a new instance of APIHandlers.
func NewAPIHandlers(locks controllers.LocksController, state controllers.StateController, snapshots controllers.SnapshotController, team controllers.TeamController, project controllers.ProjectController, environment controllers.EnvironmentController) *apiHandlers {
	return &apiHandlers{locks, state, snapshots, team, project, environment}
}

// Get system health status
// (GET /_health)
func (h *apiHandlers) GetHealth(ctx context.Context, request openapi.GetHealthRequestObject) (openapi.GetHealthResponseObject, error) {
	return openapi.GetHealth200JSONResponse{}, nil // this is just up
}

// Get system readiness
// (GET /_ready)
func (h *apiHandlers) GetReady(ctx context.Context, request openapi.GetReadyRequestObject) (openapi.GetReadyResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// Get a list of projects
// (GET /project)
func (h *apiHandlers) GetProjects(ctx context.Context, request openapi.GetProjectsRequestObject) (openapi.GetProjectsResponseObject, error) {
	query := dto.FromGetProjectsRequestObject(request)

	results, err := h.project.ListProjects(ctx, query)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return dto.ToGetProjectsResponseObject(results), nil
}

// Create a new project
// (POST /project)
func (h *apiHandlers) CreateProject(ctx context.Context, request openapi.CreateProjectRequestObject) (openapi.CreateProjectResponseObject, error) {
	cmd := dto.FromCreateProjectRequestObject(request)

	err := h.project.CreateProject(ctx, cmd)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return dto.ToCreateProjectResponseObject(), nil
}

// Delete a project
// (DELETE /project/{id})
func (h *apiHandlers) DeleteProject(ctx context.Context, request openapi.DeleteProjectRequestObject) (openapi.DeleteProjectResponseObject, error) {
	cmd := dto.FromDeleteProjectRequestObject(request)

	err := h.project.DeleteProject(ctx, cmd)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return dto.ToDeleteProjectResponseObject(), nil
}

// Get a project
// (GET /project/{id})
func (h *apiHandlers) GetProject(ctx context.Context, request openapi.GetProjectRequestObject) (openapi.GetProjectResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// Update a project
// (PUT /project/{id})
func (h *apiHandlers) UpdateProject(ctx context.Context, request openapi.UpdateProjectRequestObject) (openapi.UpdateProjectResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// Get a list of environments
// (GET /project/{projectId}/environment)
func (h *apiHandlers) GetEnvironments(ctx context.Context, request openapi.GetEnvironmentsRequestObject) (openapi.GetEnvironmentsResponseObject, error) {
	query := dto.FromGetEnvironmentsRequestObject(request)

	results, err := h.environment.ListEnvironments(ctx, query)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return dto.ToGetEnvironmentsResponseObject(results), nil
}

// Create a new environment
// (POST /project/{projectId}/environment)
func (h *apiHandlers) CreateEnvironment(ctx context.Context, request openapi.CreateEnvironmentRequestObject) (openapi.CreateEnvironmentResponseObject, error) {
	cmd := dto.FromCreateEnvironmentRequestObject(request)

	err := h.environment.CreateEnvironment(ctx, cmd)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return dto.ToCreateEnvironmentResponseObject(), nil
}

// Delete an environment
// (DELETE /project/{projectId}/environment/{environmentId})
func (h *apiHandlers) DeleteEnvironment(ctx context.Context, request openapi.DeleteEnvironmentRequestObject) (openapi.DeleteEnvironmentResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// Get an environment
// (GET /project/{projectId}/environment/{environmentId})
func (h *apiHandlers) GetEnvironment(ctx context.Context, request openapi.GetEnvironmentRequestObject) (openapi.GetEnvironmentResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// Update an environment
// (PUT /project/{projectId}/environment/{environmentId})
func (h *apiHandlers) UpdateEnvironment(ctx context.Context, request openapi.UpdateEnvironmentRequestObject) (openapi.UpdateEnvironmentResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// Get a list of snapshots
// (GET /snapshot)
func (h *apiHandlers) GetSnapshots(ctx context.Context, request openapi.GetSnapshotsRequestObject) (openapi.GetSnapshotsResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// Delete a snapshot
// (DELETE /snapshot/{id})
func (h *apiHandlers) DeleteSnapshot(ctx context.Context, request openapi.DeleteSnapshotRequestObject) (openapi.DeleteSnapshotResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// Get a snapshot
// (GET /snapshot/{id})
func (h *apiHandlers) GetSnapshot(ctx context.Context, request openapi.GetSnapshotRequestObject) (openapi.GetSnapshotResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// Get a list of teams
// (GET /team)
func (h *apiHandlers) GetTeams(ctx context.Context, request openapi.GetTeamsRequestObject) (openapi.GetTeamsResponseObject, error) {
	cmd := dto.FromGetTeamsRequestObject(request)

	results, err := h.team.ListTeams(ctx, cmd)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return dto.ToGetTeamsResponseObject(results), nil
}

// Create a new team
// (POST /team)
func (h *apiHandlers) CreateTeam(ctx context.Context, request openapi.CreateTeamRequestObject) (openapi.CreateTeamResponseObject, error) {
	cmd := dto.FromCreateTeamRequestObject(request)

	err := h.team.CreateTeam(ctx, cmd)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return dto.ToCreateTeamResponseObject(), nil
}

// Delete a team
// (DELETE /team/{id})
func (h *apiHandlers) DeleteTeam(ctx context.Context, request openapi.DeleteTeamRequestObject) (openapi.DeleteTeamResponseObject, error) {
	cmd := dto.FromDeleteTeamRequestObject(request)

	err := h.team.DeleteTeam(ctx, cmd)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return dto.ToDeleteTeamResponseObject(), nil
}

// Get a team
// (GET /team/{id})
func (h *apiHandlers) GetTeam(ctx context.Context, request openapi.GetTeamRequestObject) (openapi.GetTeamResponseObject, error) {
	query := dto.FromGetTeamRequestObject(request)

	team, err := h.team.GetTeam(ctx, query)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return dto.ToGetTeamResponseObject(team), nil
}

// Update a team
// (PUT /team/{id})
func (h *apiHandlers) UpdateTeam(ctx context.Context, request openapi.UpdateTeamRequestObject) (openapi.UpdateTeamResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// Get a list of users
// (GET /user)
func (h *apiHandlers) GetUsers(ctx context.Context, request openapi.GetUsersRequestObject) (openapi.GetUsersResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// Create a new user
// (POST /user)
func (h *apiHandlers) PostUser(ctx context.Context, request openapi.PostUserRequestObject) (openapi.PostUserResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// Delete a user
// (DELETE /user/{id})
func (h *apiHandlers) DeleteUser(ctx context.Context, request openapi.DeleteUserRequestObject) (openapi.DeleteUserResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// Get a user
// (GET /user/{id})
func (h *apiHandlers) GetUser(ctx context.Context, request openapi.GetUserRequestObject) (openapi.GetUserResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// Update a user
// (PUT /user/{id})
func (h *apiHandlers) UpdateUser(ctx context.Context, request openapi.UpdateUserRequestObject) (openapi.UpdateUserResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
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
