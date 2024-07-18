package authz

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zeiss/knox/pkg/oas"
)

// Resolvers is a map of operationID to AuthzResolverFunc.
func Resolvers() oas.ResolverMap {
	return oas.ResolverMap{
		"GetTeam": func(c *fiber.Ctx) (oas.User, oas.Relation, oas.Object, error) {
			return oas.NoopUser, oas.NoopRelation, oas.NoopObject, nil
		},
	}
}
