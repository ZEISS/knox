package authz

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zeiss/knox/internal/authn/oidc"
	"github.com/zeiss/knox/internal/authz/fga"
)

// Resolvers is a map of operationID to AuthzResolverFunc.
func Resolvers() fga.ResolverMap {
	return fga.ResolverMap{
		"GetTeam": func(c *fiber.Ctx) (fga.User, fga.Relation, fga.Object, error) {
			user := fga.NoopUser
			relation := fga.NoopRelation
			object := fga.NoopObject

			principal, ok := oidc.GetJWTFromContext(c.UserContext())
			if !ok {
				return user, relation, object, fiber.ErrUnauthorized
			}
			user = fga.User("user:" + principal.Subject)
			relation = fga.Relation("admin")
			object = fga.Object("team:" + c.Params("teamId", ""))

			return user, relation, object, nil
		},
	}
}
