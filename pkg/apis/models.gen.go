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
	Description *string `json:"description,omitempty"`
	RecordType  *string `json:"record_type,omitempty"`
	RecordUuid  *string `json:"record_uuid,omitempty"`
	TeamId      *int    `json:"team_id,omitempty"`
	Title       *string `json:"title,omitempty"`
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

// ProjectId defines model for projectId.
type ProjectId = string

// TeamId defines model for teamId.
type TeamId = string

// GetApiV1ProjectParams defines parameters for GetApiV1Project.
type GetApiV1ProjectParams struct {
	Limit  *int `form:"limit,omitempty" json:"limit,omitempty"`
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetApiV1ProjectProjectIdEnvironmentParams defines parameters for GetApiV1ProjectProjectIdEnvironment.
type GetApiV1ProjectProjectIdEnvironmentParams struct {
	Limit  *int `form:"limit,omitempty" json:"limit,omitempty"`
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetApiV1SnapshotParams defines parameters for GetApiV1Snapshot.
type GetApiV1SnapshotParams struct {
	Limit  *int `form:"limit,omitempty" json:"limit,omitempty"`
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetApiV1TeamParams defines parameters for GetApiV1Team.
type GetApiV1TeamParams struct {
	Limit  *int `form:"limit,omitempty" json:"limit,omitempty"`
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetApiV1UserParams defines parameters for GetApiV1User.
type GetApiV1UserParams struct {
	Limit  *int `form:"limit,omitempty" json:"limit,omitempty"`
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`
}

// PostApiV1ProjectJSONRequestBody defines body for PostApiV1Project for application/json ContentType.
type PostApiV1ProjectJSONRequestBody = ProjectCreate

// PutApiV1ProjectIdJSONRequestBody defines body for PutApiV1ProjectId for application/json ContentType.
type PutApiV1ProjectIdJSONRequestBody = ProjectUpdate

// PostApiV1ProjectProjectIdEnvironmentJSONRequestBody defines body for PostApiV1ProjectProjectIdEnvironment for application/json ContentType.
type PostApiV1ProjectProjectIdEnvironmentJSONRequestBody = EnvironmentCreate

// PutApiV1ProjectProjectIdEnvironmentEnvironmentIdJSONRequestBody defines body for PutApiV1ProjectProjectIdEnvironmentEnvironmentId for application/json ContentType.
type PutApiV1ProjectProjectIdEnvironmentEnvironmentIdJSONRequestBody = EnvironmentUpdate

// PostApiV1SnapshotJSONRequestBody defines body for PostApiV1Snapshot for application/json ContentType.
type PostApiV1SnapshotJSONRequestBody = SnapshotCreate

// PutApiV1SnapshotIdJSONRequestBody defines body for PutApiV1SnapshotId for application/json ContentType.
type PutApiV1SnapshotIdJSONRequestBody = SnapshotUpdate

// PostApiV1TeamJSONRequestBody defines body for PostApiV1Team for application/json ContentType.
type PostApiV1TeamJSONRequestBody = TeamCreate

// PutApiV1TeamIdJSONRequestBody defines body for PutApiV1TeamId for application/json ContentType.
type PutApiV1TeamIdJSONRequestBody = TeamUpdate

// PostApiV1UserJSONRequestBody defines body for PostApiV1User for application/json ContentType.
type PostApiV1UserJSONRequestBody = UserCreate

// PutApiV1UserIdJSONRequestBody defines body for PutApiV1UserId for application/json ContentType.
type PutApiV1UserIdJSONRequestBody = UserUpdate

// LockEnvironmentJSONRequestBody defines body for LockEnvironment for application/json ContentType.
type LockEnvironmentJSONRequestBody = LockInfo

// UpdateEnvironmentStateJSONRequestBody defines body for UpdateEnvironmentState for application/json ContentType.
type UpdateEnvironmentStateJSONRequestBody = Payload

// UnlockEnvironmentJSONRequestBody defines body for UnlockEnvironment for application/json ContentType.
type UnlockEnvironmentJSONRequestBody = LockInfo

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"KJKCowUTcvZ2+naKVjer/wIAAP//Nenf1nVkAAA=",
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
