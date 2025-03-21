// Package apis provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
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
	BasicAuthScopes = "basicAuth.Scopes"
	OpenIdScopes    = "openId.Scopes"
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
	StateVersion *int       `json:"stateVersion,omitempty"`
	UpdatedAt    *time.Time `json:"updatedAt,omitempty"`
	Username     *string    `json:"username,omitempty"`
}

// EnvironmentCreate defines model for EnvironmentCreate.
type EnvironmentCreate struct {
	Name     *string `json:"name,omitempty"`
	Secret   *string `json:"secret,omitempty"`
	Username *string `json:"username,omitempty"`
}

// EnvironmentUpdate defines model for EnvironmentUpdate.
type EnvironmentUpdate struct {
	Name     *string `json:"name,omitempty"`
	Secret   *string `json:"secret,omitempty"`
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
	Team        *struct {
		Id   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	} `json:"team,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// ProjectCreate defines model for ProjectCreate.
type ProjectCreate struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

// ProjectUpdate defines model for ProjectUpdate.
type ProjectUpdate struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	TeamId      *int    `json:"team_id,omitempty"`
}

// Snapshot defines model for Snapshot.
type Snapshot struct {
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
	Description *string    `json:"description,omitempty"`
	Id          *string    `json:"id,omitempty"`
	Title       *string    `json:"title,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	Version     *int       `json:"version,omitempty"`
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

// State defines model for State.
type State struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Id        *string    `json:"id,omitempty"`
	State     *string    `json:"state,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Version   *int       `json:"version,omitempty"`
}

// Team defines model for Team.
type Team struct {
	CreatedAt     *time.Time `json:"createdAt,omitempty"`
	Description   *string    `json:"description,omitempty"`
	Id            *string    `json:"id,omitempty"`
	Members       *[]string  `json:"members,omitempty"`
	Name          *string    `json:"name,omitempty"`
	ProjectsCount *int       `json:"projectsCount,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
	UsersCount    *int       `json:"usersCount,omitempty"`
}

// TeamCreate defines model for TeamCreate.
type TeamCreate struct {
	Description *string   `json:"description,omitempty"`
	Members     *[]string `json:"members,omitempty"`
	Name        *string   `json:"name,omitempty"`
}

// TeamUpdate defines model for TeamUpdate.
type TeamUpdate struct {
	Description *string   `json:"description,omitempty"`
	Members     *[]string `json:"members,omitempty"`
	Name        *string   `json:"name,omitempty"`
}

// EnvironmentName defines model for environmentName.
type EnvironmentName = string

// Limit defines model for limit.
type Limit = int

// Offset defines model for offset.
type Offset = int

// ProjectName defines model for projectName.
type ProjectName = string

// SnapshotId defines model for snapshotId.
type SnapshotId = openapi_types.UUID

// TeamName defines model for teamName.
type TeamName = string

// GetTeamsParams defines parameters for GetTeams.
type GetTeamsParams struct {
	Limit  *Limit  `form:"limit,omitempty" json:"limit,omitempty"`
	Offset *Offset `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetProjectsParams defines parameters for GetProjects.
type GetProjectsParams struct {
	Limit  *int `form:"limit,omitempty" json:"limit,omitempty"`
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetEnvironmentsParams defines parameters for GetEnvironments.
type GetEnvironmentsParams struct {
	Limit  *Limit  `form:"limit,omitempty" json:"limit,omitempty"`
	Offset *Offset `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetSnapshotsParams defines parameters for GetSnapshots.
type GetSnapshotsParams struct {
	Limit  *Limit  `form:"limit,omitempty" json:"limit,omitempty"`
	Offset *Offset `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetStatesParams defines parameters for GetStates.
type GetStatesParams struct {
	Limit  *Limit  `form:"limit,omitempty" json:"limit,omitempty"`
	Offset *Offset `form:"offset,omitempty" json:"offset,omitempty"`
}

// LockEnvironmentJSONRequestBody defines body for LockEnvironment for application/json ContentType.
type LockEnvironmentJSONRequestBody = LockInfo

// UpdateEnvironmentStateJSONRequestBody defines body for UpdateEnvironmentState for application/json ContentType.
type UpdateEnvironmentStateJSONRequestBody = Payload

// UnlockEnvironmentJSONRequestBody defines body for UnlockEnvironment for application/json ContentType.
type UnlockEnvironmentJSONRequestBody = LockInfo

// CreateTeamJSONRequestBody defines body for CreateTeam for application/json ContentType.
type CreateTeamJSONRequestBody = TeamCreate

// UpdateTeamJSONRequestBody defines body for UpdateTeam for application/json ContentType.
type UpdateTeamJSONRequestBody = TeamUpdate

// CreateProjectJSONRequestBody defines body for CreateProject for application/json ContentType.
type CreateProjectJSONRequestBody = ProjectCreate

// UpdateProjectJSONRequestBody defines body for UpdateProject for application/json ContentType.
type UpdateProjectJSONRequestBody = ProjectUpdate

// CreateEnvironmentJSONRequestBody defines body for CreateEnvironment for application/json ContentType.
type CreateEnvironmentJSONRequestBody = EnvironmentCreate

// UpdateEnvironmentJSONRequestBody defines body for UpdateEnvironment for application/json ContentType.
type UpdateEnvironmentJSONRequestBody = EnvironmentUpdate

// CreateSnapshotJSONRequestBody defines body for CreateSnapshot for application/json ContentType.
type CreateSnapshotJSONRequestBody = SnapshotCreate

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xcWXPcuBH+Kygkj7Q43myqtubN2RzrZJNyWdrkwVbJGLJniDUJcAFQWlk1/z0FgOfw",
	"HIqc0cE3ibga3V+f7OED9ngUcwZMSbx+wDERJAIFwvwH7JYKziJg6j8kAv2IMrzGMVEBdjAzz2qzHCzg",
	"t4QK8PFaiQQcLL0AIqKXq/tYL5FKULbD+72DQxpRpYd82JIkVHj9duXYY35LQNwX59iZDbtRpmAHwmzH",
	"t1sJ1f1atktn9uwXC/4reF3XL8847uqSkVgGXL33W7YuTejaectFRBRe4yShemb9JAUk6rhCPnwM/fts",
	"0EDlbwUIDI4Ej0EoCmbQE0AU+O9UhVqfKHijqDn2YHMHU7/hTAdT+TP3vkJ5cMN5CITpUZZesbYslVGd",
	"ssZj9jk9fGOW6auCJyyuaptLRRT8F4SknDWhyMFJ7B97/0SCaLlOE3kl9v9omF2/aitzOq42nopfzJXP",
	"Q4UQXHwEGXMmGygAPfxvkJLshu74E5BQBaO3tBBJ5MDTNMLfsy1v1aNjtajHPmjzaE+rDejDiariuqRV",
	"2og0DdzWlKEYuwv4QD58IPchJ9UbbCgjxojXZ7ep+Ajj44P0BI1bLw7s9keeMNWs7S2mqxX52voOtEyt",
	"uzTx72irs29na5tR6ePUEeSmJ7UZjpEnWfbeVLhZ9u81Mi5TnzsRkEKYGHstuFBUhc33H+F8btudWRfL",
	"xoLEONCbgbaq7aJdhI3FVAd0jiRENZ5/GkS1IEZmNJ0BMVeNJm8GS91y9wiiTZrjUAWRbJa+fUCEIPdD",
	"okvZ4RNGRoDtW7ZxdawOTsmRNtrGquG8tNnYMxFU3V/qfMbStCGSeu8SG+CYPMfkGvppIaxAqTgNk5hN",
	"4jZABIi/ZwL+5/+usgTTLDejh+v3pQiswgp8FVCJqEQqAPQvxn9H7z68RzIGj26pZ+KyC/ucSkSQBHFL",
	"PUAqIApFhJEdSHQFQhANOFRK0yUizEdZdikvPrPP7CoAAWYjhsgtoSHZhIB+gjBCXkCEQlsuzGFm9pcv",
	"X2QAYfiZBXqGgJgj4vvoq6ZGX0uuXfcbUCkvdlQFyeaCclcPunq+a3aU5cVWQdInEojwAjugF5nzPjOc",
	"G12sCcEls4PfXqwuVpksSEzxGv/pYnXxFtsw1cjUvQlMIK//3tlcIw9wtfjwP0DZUN8kwjbaNyu/W62M",
	"geJMpQkuieMwlYH7q7TALXLlPwrY4jX+g1tUV9w0WXYPkgkj/qrYL++lgkjLwtJr8PznM5Hg8ztmlSSJ",
	"Ih17GzYhaSdYAlGa3zhYkZ3E60/4Z3oLDKTE13qpeyOA+PddfP9oJjwRthtqkeKIeB7ECilBtlvqnU0M",
	"jKt2mtpEoxdQI4NCLB/zZ1YuXkiBKfchKwDt3YdSOWvvPhzU9vZuyL2vxoJz2SBInbqWK0FOpaD4qZkd",
	"xRQ3L0Ttnd655brbgOmHVcr9tS12gVR/4Rabkwg1z973Rp6z4bl8Tg02ieeBlNskRCJHl4O/X30/2fHV",
	"IksDDR9B8kR4YMC75Qnzp1afXhLeMwWCkRBdgrgFgcyCisM3mCy5+k/XGhaFOmkmG+9rgmXEt83+tKRh",
	"JfSPV7I8NG+zlqVTbGrx5BVtJi3IakSLEsynBNqnPEYHnBZnYdOB5wfl6X1GCcX7RVmetbJYUJ/BZySs",
	"OzT7xYwvwVlfcDZIL76zerHo0Dw6ZKD6aB3SeJVdUdSVmXCsGtg2hAGgThsMHh3/VItlESjiE9UwkndS",
	"1MuPRVdEQx2bKxIeVWfMGZtX47oAY2q9tQJdQx1uqFN6imjOKoBNsRNBIZVKg1ilgMtAawF47eDf32zp",
	"BsQbkqjg25vtztCdMmZd7dFhSRjmjRu7kG9ImDZyyJh4ptRp8n9s7FuYv71NVxA/oiwrLxveJSq4yURB",
	"fe9gMzPN8KXZrdh6s5HxPCa7VNRuNNpvJz2pCQL6OcpevT8z/FnOIYIY3Bn4jUDfybGWG+9S/GPdbQg2",
	"L65i8K/meYrBsTFN3UovLn4mUFp5ITIWkGVz+Km7oU0ffALb2BVhTIrJ1ezWboH4dH6/Gd/alyat9ZAp",
	"IDOPF05f385cpFhwOScu0/JEKzSbXG+WzHfmUh+yOY+qKgxr+650itcTloeh7d7lDvHaNi8uaytLcVDi",
	"lnVULrmbzt3iAuCZ0uSYv+5JjzJOPjmzXu3unDm/yvFUl1Q69CKyrDgXdgNOOu1rtcbbn/E8HlZHFnKX",
	"HOn0OVIXnCbNlJzeXzcdJlMFaRI0ZBQXeI3dk2VXTwD/q1NYx0WjpnPnL0adBBA/7S0+upLbnn2eR6Vm",
	"C2xOk7IuqnuqrHWq2Kr8wlkObLSST+mF+ZlfQh6yb1A+We49aGjYP3uK/CrTWqgCvPl1+qtxjB05/Fk6",
	"Z2byjvVfi8+c+ldUv47o0vCLKAEMaVF5Vjp1J6gapVNjfXK9wa2/HPKsetuWKsopqyjsSaqkM+hzOoea",
	"W73JrNrrDImLX+1vD3qc2qLBk0Wrr1t956n0vPJO8PqnimYuEi3W4oSFosXfnzlad/OPC3RV1y7zSU/Z",
	"AjnLTwJ6m0sq4h5UDcy/s/QS20taim2yhPfMJBU60NdFknPsFbrrg09MzVyuKrDZgLx07EUUqmQBqSY8",
	"zuEP3Ifiq6YDijrPAvT9S0pfcl3qP2fooukButMborxo/K1OYzkXSE/2+u6EhlsRBd1RvJ2xhPCPCOEL",
	"Lg8Lns2XMqaMnBctPFoLj/p9eqokS5GyVnMwLNcysBdORJh+GVGuXTfkHgkDLtX6h9UPK7y/3v8/AAD/",
	"/3Rjw0qyYAAA",
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
