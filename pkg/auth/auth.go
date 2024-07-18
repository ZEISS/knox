package auth

import (
	"context"
	"encoding/base64"
	"strings"

	"github.com/zeiss/knox/internal/ports"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	middleware "github.com/oapi-codegen/fiber-middleware"
	seed "github.com/zeiss/gorm-seed"
)

// OpenAPIAuthenticatorOpts are the OpenAPI authenticator options.
type OpenAPIAuthenticatorOpts struct {
	BasicAuthenticator openapi3filter.AuthenticationFunc
}

// Conigure the OpenAPI authenticator.
func (o *OpenAPIAuthenticatorOpts) Conigure(opts ...OpenAPIAuthenticatorOpt) {
	for _, opt := range opts {
		opt(o)
	}
}

// OpenAPIAuthenticatorOpt is a function that sets an option on the OpenAPI authenticator.
type OpenAPIAuthenticatorOpt func(*OpenAPIAuthenticatorOpts)

func OpenAPIAuthenticatorDefaultOpts() OpenAPIAuthenticatorOpts {
	return OpenAPIAuthenticatorOpts{}
}

// NoopBasicAuthenticator is a no-op basic authenticator.
func NoopBasicAuthenticator(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	return nil
}

// WithBasicAuthenticator sets the basic authenticator.
func WithBasicAuthenticator(auth openapi3filter.AuthenticationFunc) OpenAPIAuthenticatorOpt {
	return func(o *OpenAPIAuthenticatorOpts) {
		o.BasicAuthenticator = auth
	}
}

// NewBasicAuthenticator returns a new basic authenticator.
func NewBasicAuthenticator(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		c := middleware.GetFiberContext(ctx)

		// Get authorization header
		auth := c.Get(fiber.HeaderAuthorization)

		if len(auth) <= 6 || !utils.EqualFold(auth[:6], "basic ") {
			return fiber.NewError(fiber.StatusForbidden, "forbidden")
		}

		// Decode the header contents
		raw, err := base64.StdEncoding.DecodeString(auth[6:])
		if err != nil {
			return fiber.NewError(fiber.StatusForbidden, "forbidden")
		}

		// Get the credentials
		creds := utils.UnsafeString(raw)

		index := strings.Index(creds, ":")
		if index == -1 {
			return fiber.NewError(fiber.StatusForbidden, "forbidden")
		}

		username := creds[:index]
		password := creds[index+1:]

		teamId, ok := input.RequestValidationInput.PathParams["teamId"]
		if !ok {
			return fiber.NewError(fiber.StatusForbidden, "forbidden")
		}

		projectID, ok := input.RequestValidationInput.PathParams["projectId"]
		if !ok {
			return fiber.NewError(fiber.StatusForbidden, "forbidden")
		}

		environmentId, ok := input.RequestValidationInput.PathParams["environmentId"]
		if !ok {
			return fiber.NewError(fiber.StatusForbidden, "forbidden")
		}

		err = store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
			return tx.AuthenticateClient(ctx, teamId, projectID, environmentId, username, password)
		})
		if err != nil {
			return fiber.NewError(fiber.StatusForbidden, "forbidden")
		}

		return nil
	}
}

// NewAuthenticator returns a new authenticator.
func NewAuthenticator(opts ...OpenAPIAuthenticatorOpt) openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		options := OpenAPIAuthenticatorDefaultOpts()
		options.Conigure(opts...)

		if input.SecuritySchemeName == "basic_auth" {
			return options.BasicAuthenticator(ctx, input)
		}

		return nil
	}
}
