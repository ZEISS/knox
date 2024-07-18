package oas

import (
	"context"

	openapi "github.com/zeiss/knox/pkg/apis"

	"github.com/gofiber/fiber/v2"
	"github.com/openfga/go-sdk/client"
)

// User is the subject.
type User string

// String returns the string representation of the user.
func (u User) String() string {
	return string(u)
}

// Relation is the object.
type Relation string

// String returns the string representation of the relation.
func (r Relation) String() string {
	return string(r)
}

// Object is the action.
type Object string

// String returns the string representation of the object.
func (o Object) String() string {
	return string(o)
}

const (
	NoopUser     User     = ""
	NoopRelation Relation = ""
	NoopObject   Object   = ""
)

// Checker is the interface for the FGA authz checker.
type Checker interface {
	Allowed(ctx context.Context, user User, relation Relation, object Object) (bool, error)
}

var _ Checker = (*fgaImpl)(nil)

type fgaImpl struct {
	client *client.OpenFgaClient
}

// NewChecker returns a new FGA authz checker.
func NewChecker(c *client.OpenFgaClient) *fgaImpl {
	return &fgaImpl{c}
}

// Allowed returns true if user has the relation with the object.
func (f *fgaImpl) Allowed(ctx context.Context, user User, relation Relation, object Object) (bool, error) {
	body := client.ClientCheckRequest{
		User:     user.String(),
		Relation: relation.String(),
		Object:   object.String(),
	}

	allowed, err := f.client.Check(ctx).Body(body).Execute()
	if err != nil {
		return false, err
	}

	return allowed.GetAllowed(), nil
}

type noopImpl struct{}

// NewNoop returns a new Noop authz checker.
func NewNoop() *noopImpl {
	return &noopImpl{}
}

// Allowed returns true if user has the relation with the object.
func (n *noopImpl) Allowed(ctx context.Context, user User, relation Relation, object Object) (bool, error) {
	return false, nil
}

// NoopResolvers is a map of Noop resolvers.
func NoopResolvers() map[string]AuthzResolverFunc {
	return map[string]AuthzResolverFunc{}
}

// Config ...
type Config struct {
	// Next defines a function to skip the current middleware.
	Next func(c *fiber.Ctx) bool
	// Checker defines a function to check the authorization.
	Checker Checker
	// Resolvers defines the resolvers for a specific operation.
	Resolvers map[string]AuthzResolverFunc
	// DefaultError defines the default error.
	DefaultError error
}

// DefaultConfig contains the default configuration.
var DefaultConfig = Config{
	Checker:      NewNoop(),
	Resolvers:    NoopResolvers(),
	DefaultError: fiber.ErrForbidden,
}

// NoopResolver is a resolver that always returns Noop values.
func NoopResolver() AuthzResolverFunc {
	return func(_ *fiber.Ctx) (User, Relation, Object, error) {
		return NoopUser, NoopRelation, NoopObject, nil
	}
}

// AuthzResolverFunc is a function to resolve the authz values.
type AuthzResolverFunc func(ctx *fiber.Ctx) (User, Relation, Object, error)

// NewAuthz returns a new authz middleware.
// nolint:contextcheck
func NewAuthz(config ...Config) openapi.StrictMiddlewareFunc {
	cfg := configDefault(config...)

	return func(f openapi.StrictHandlerFunc, operationID string) openapi.StrictHandlerFunc {
		return func(c *fiber.Ctx, args interface{}) (interface{}, error) {
			if cfg.Next != nil && cfg.Next(c) {
				return f(c, args)
			}

			resolver, ok := cfg.Resolvers[operationID]
			if !ok {
				return nil, cfg.DefaultError
			}

			user, relation, object, err := resolver(c)
			if err != nil {
				return nil, cfg.DefaultError
			}

			allowed, err := cfg.Checker.Allowed(c.Context(), user, relation, object)
			if err != nil {
				return nil, cfg.DefaultError
			}

			if !allowed {
				return nil, cfg.DefaultError
			}

			return f(c, args)
		}
	}
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return DefaultConfig
	}

	// Override default config
	cfg := config[0]

	if cfg.Checker == nil {
		cfg.Checker = DefaultConfig.Checker
	}

	if cfg.DefaultError == nil {
		cfg.DefaultError = DefaultConfig.DefaultError
	}

	return cfg
}
