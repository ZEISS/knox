// Package apis provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package apis

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	Api_keyScopes    = "api_key.Scopes"
	Basic_authScopes = "basic_auth.Scopes"
)

// Environment defines model for Environment.
type Environment struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        *string    `json:"id,omitempty"`
	IsLocked  *bool      `json:"isLocked,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Project   *struct {
		Id *string `json:"id,omitempty"`
	} `json:"project,omitempty"`
	Secret       *string    `json:"secret,omitempty"`
	Slug         *string    `json:"slug,omitempty"`
	StateVersion *int       `json:"stateVersion,omitempty"`
	UpdatedAt    *time.Time `json:"updatedAt,omitempty"`
	Username     *string    `json:"username,omitempty"`
}

// EnvironmentCreate defines model for EnvironmentCreate.
type EnvironmentCreate struct {
	Name     *string `json:"name,omitempty"`
	Secret   *string `json:"secret,omitempty"`
	Slug     *string `json:"slug,omitempty"`
	Username *string `json:"username,omitempty"`
}

// EnvironmentUpdate defines model for EnvironmentUpdate.
type EnvironmentUpdate struct {
	Name     *string `json:"name,omitempty"`
	Secret   *string `json:"secret,omitempty"`
	Slug     *string `json:"slug,omitempty"`
	Username *string `json:"username,omitempty"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// HealthResponse defines model for HealthResponse.
type HealthResponse struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`
	Status       *string `json:"status,omitempty"`
}

// LockInfo defines model for LockInfo.
type LockInfo struct {
	Created   *time.Time          `json:"created,omitempty"`
	Id        *openapi_types.UUID `json:"id,omitempty"`
	Info      *string             `json:"info,omitempty"`
	Operation *string             `json:"operation,omitempty"`
	Path      *string             `json:"path,omitempty"`
	Version   *string             `json:"version,omitempty"`
	Who       *string             `json:"who,omitempty"`
}

// Payload defines model for Payload.
type Payload = map[string]interface{}

// Project defines model for Project.
type Project struct {
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	Description *string    `json:"description,omitempty"`
	EnvCount    *int       `json:"envCount,omitempty"`
	Id          *string    `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Slug        *string    `json:"slug,omitempty"`
	Team        *struct {
		Id   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
		Slug *string `json:"slug,omitempty"`
	} `json:"team,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// ProjectCreate defines model for ProjectCreate.
type ProjectCreate struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	TeamId      *int    `json:"team_id,omitempty"`
}

// ProjectUpdate defines model for ProjectUpdate.
type ProjectUpdate struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	TeamId      *int    `json:"team_id,omitempty"`
}

// Snapshot defines model for Snapshot.
type Snapshot struct {
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	Description *string    `json:"description,omitempty"`
	Id          *string    `json:"id,omitempty"`
	RecordType  *string    `json:"record_type,omitempty"`
	RecordUuid  *string    `json:"record_uuid,omitempty"`
	Status      *string    `json:"status,omitempty"`
	Team        *struct {
		Id   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
		Slug *string `json:"slug,omitempty"`
	} `json:"team,omitempty"`
	Title     *string    `json:"title,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// SnapshotCreate defines model for SnapshotCreate.
type SnapshotCreate struct {
	Description *string             `json:"description,omitempty"`
	StateId     *openapi_types.UUID `json:"state_id,omitempty"`
	Title       *string             `json:"title,omitempty"`
}

// SnapshotUpdate defines model for SnapshotUpdate.
type SnapshotUpdate struct {
	Description *string `json:"description,omitempty"`
	TeamId      *int    `json:"team_id,omitempty"`
	Title       *string `json:"title,omitempty"`
}

// Task defines model for Task.
type Task struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        *string    `json:"id,omitempty"`
	RunAt     *time.Time `json:"runAt,omitempty"`
	Status    *string    `json:"status,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Team defines model for Team.
type Team struct {
	CreatedAt     *time.Time `json:"createdAt,omitempty"`
	Description   *string    `json:"description,omitempty"`
	Id            *string    `json:"id,omitempty"`
	Members       *[]string  `json:"members,omitempty"`
	Name          *string    `json:"name,omitempty"`
	ProjectsCount *int       `json:"projectsCount,omitempty"`
	Slug          *string    `json:"slug,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
	UsersCount    *int       `json:"usersCount,omitempty"`
}

// TeamCreate defines model for TeamCreate.
type TeamCreate struct {
	Description *string   `json:"description,omitempty"`
	Members     *[]string `json:"members,omitempty"`
	Name        *string   `json:"name,omitempty"`
	Slug        *string   `json:"slug,omitempty"`
}

// TeamUpdate defines model for TeamUpdate.
type TeamUpdate struct {
	Description *string   `json:"description,omitempty"`
	Members     *[]string `json:"members,omitempty"`
	Name        *string   `json:"name,omitempty"`
	Slug        *string   `json:"slug,omitempty"`
}

// User defines model for User.
type User struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Email     *string    `json:"email,omitempty"`
	Id        *string    `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Role      *string    `json:"role,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// UserCreate defines model for UserCreate.
type UserCreate struct {
	Email    *string `json:"email,omitempty"`
	Name     *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
	Role     *string `json:"role,omitempty"`
}

// UserUpdate defines model for UserUpdate.
type UserUpdate struct {
	Email    *string `json:"email,omitempty"`
	Name     *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
	Role     *string `json:"role,omitempty"`
}

// EnvironmentId defines model for environmentId.
type EnvironmentId = string

// Limit defines model for limit.
type Limit = int

// LockId defines model for lockId.
type LockId = openapi_types.UUID

// Offset defines model for offset.
type Offset = int

// ProjectId defines model for projectId.
type ProjectId = string

// TeamId defines model for teamId.
type TeamId = string

// UpdateEnvironmentStateParams defines parameters for UpdateEnvironmentState.
type UpdateEnvironmentStateParams struct {
	ID *LockId `form:"ID,omitempty" json:"ID,omitempty"`
}

// GetProjectParams defines parameters for GetProject.
type GetProjectParams struct {
	Limit  *int `form:"limit,omitempty" json:"limit,omitempty"`
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetProjectProjectIdEnvironmentParams defines parameters for GetProjectProjectIdEnvironment.
type GetProjectProjectIdEnvironmentParams struct {
	Limit  *int `form:"limit,omitempty" json:"limit,omitempty"`
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetSnapshotParams defines parameters for GetSnapshot.
type GetSnapshotParams struct {
	Limit  *int `form:"limit,omitempty" json:"limit,omitempty"`
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetTeamParams defines parameters for GetTeam.
type GetTeamParams struct {
	Limit  *int `form:"limit,omitempty" json:"limit,omitempty"`
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetUserParams defines parameters for GetUser.
type GetUserParams struct {
	Limit  *Limit  `form:"limit,omitempty" json:"limit,omitempty"`
	Offset *Offset `form:"offset,omitempty" json:"offset,omitempty"`
}

// LockEnvironmentJSONRequestBody defines body for LockEnvironment for application/json ContentType.
type LockEnvironmentJSONRequestBody = LockInfo

// UpdateEnvironmentStateJSONRequestBody defines body for UpdateEnvironmentState for application/json ContentType.
type UpdateEnvironmentStateJSONRequestBody = Payload

// UnlockEnvironmentJSONRequestBody defines body for UnlockEnvironment for application/json ContentType.
type UnlockEnvironmentJSONRequestBody = LockInfo

// PostProjectJSONRequestBody defines body for PostProject for application/json ContentType.
type PostProjectJSONRequestBody = ProjectCreate

// PutProjectIdJSONRequestBody defines body for PutProjectId for application/json ContentType.
type PutProjectIdJSONRequestBody = ProjectUpdate

// PostProjectProjectIdEnvironmentJSONRequestBody defines body for PostProjectProjectIdEnvironment for application/json ContentType.
type PostProjectProjectIdEnvironmentJSONRequestBody = EnvironmentCreate

// PutProjectProjectIdEnvironmentEnvironmentIdJSONRequestBody defines body for PutProjectProjectIdEnvironmentEnvironmentId for application/json ContentType.
type PutProjectProjectIdEnvironmentEnvironmentIdJSONRequestBody = EnvironmentUpdate

// CreateSnapshotJSONRequestBody defines body for CreateSnapshot for application/json ContentType.
type CreateSnapshotJSONRequestBody = SnapshotCreate

// PutSnapshotIdJSONRequestBody defines body for PutSnapshotId for application/json ContentType.
type PutSnapshotIdJSONRequestBody = SnapshotUpdate

// PostTeamJSONRequestBody defines body for PostTeam for application/json ContentType.
type PostTeamJSONRequestBody = TeamCreate

// PutTeamIdJSONRequestBody defines body for PutTeamId for application/json ContentType.
type PutTeamIdJSONRequestBody = TeamUpdate

// PostUserJSONRequestBody defines body for PostUser for application/json ContentType.
type PostUserJSONRequestBody = UserCreate

// PutUserIdJSONRequestBody defines body for PutUserId for application/json ContentType.
type PutUserIdJSONRequestBody = UserUpdate

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"V5v/AgAA//9HqNtoVWQAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
