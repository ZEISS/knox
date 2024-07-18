package oidc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/zeiss/knox/internal/authn"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hashicorp/go-retryablehttp"
	middleware "github.com/oapi-codegen/fiber-middleware"
)

type contextKey int

const jwtToken contextKey = iota

var jwkRefreshInterval = 48 * time.Hour

var (
	ErrNoAuthHeader      = fiber.NewError(fiber.StatusUnauthorized, "Authorization header is missing")
	ErrInvalidAuthHeader = fiber.NewError(fiber.StatusUnauthorized, "Authorization header is invalid")
	ErrInvalidToken      = fiber.NewError(fiber.StatusUnauthorized, "token is invalid")
	ErrInvalidIssuer     = fiber.NewError(fiber.StatusUnauthorized, "issuer is invalid")
	ErrClaimsInvalid     = fiber.NewError(fiber.StatusUnauthorized, "claims are invalid")
	ErrInvalidAudiance   = fiber.NewError(fiber.StatusUnauthorized, "audience is invalid")
	ErrInvalidSubject    = fiber.NewError(fiber.StatusUnauthorized, "subject is invalid")
)

// Validator is an interface for validating tokens
type Validator interface {
	// Validate validates the provided token.
	Validate(req *http.Request) (*authn.AuthClaims, error)
}

// RemoteOidcValidator is an OIDC validator that validates tokens using a remote OIDC provider.
type RemoteOidcValidator struct {
	MainIssuer    string
	IssuerAliases []string
	Audience      string

	JwksURI string
	JWKs    *keyfunc.JWKS

	httpClient *http.Client
}

// NewRemoteOidcValidator creates a new RemoteOidcValidator.
func NewRemoteOidcValidator(mainIssuer string, issuerAliases []string, audience string) (*RemoteOidcValidator, error) {
	client := retryablehttp.NewClient()
	client.Logger = nil

	oidc := &RemoteOidcValidator{
		MainIssuer:    mainIssuer,
		IssuerAliases: issuerAliases,
		Audience:      audience,
		httpClient:    client.StandardClient(),
	}

	oidcConfig, err := oidc.GetConfiguration()
	if err != nil {
		return nil, fmt.Errorf("error fetching OIDC configuration: %w", err)
	}

	oidc.JwksURI = oidcConfig.JWKsURI
	jwks, err := oidc.GetKeys()
	if err != nil {
		return nil, fmt.Errorf("error fetching OIDC keys: %w", err)
	}

	oidc.JWKs = jwks

	return oidc, nil
}

func (oidc *RemoteOidcValidator) GetKeys() (*keyfunc.JWKS, error) {
	jwks, err := keyfunc.Get(oidc.JwksURI, keyfunc.Options{
		Client:          oidc.httpClient,
		RefreshInterval: jwkRefreshInterval,
	})
	if err != nil {
		return nil, fmt.Errorf("error fetching keys from %v: %w", oidc.JwksURI, err)
	}

	return jwks, nil
}

// GetConfiguration fetches the OIDC configuration from the issuer.
// nolint:noctx
func (oidc *RemoteOidcValidator) GetConfiguration() (*authn.OidcConfig, error) {
	wellKnown := strings.TrimSuffix(oidc.MainIssuer, "/") + "/.well-known/openid-configuration"
	req, err := http.NewRequest("GET", wellKnown, nil)
	if err != nil {
		return nil, fmt.Errorf("error forming request to get OIDC: %w", err)
	}

	res, err := oidc.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error getting OIDC: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code getting OIDC: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	oidcConfig := &authn.OidcConfig{}
	if err := json.Unmarshal(body, oidcConfig); err != nil {
		return nil, fmt.Errorf("failed parsing document: %w", err)
	}

	if oidcConfig.Issuer == "" {
		return nil, errors.New("missing issuer value")
	}

	if oidcConfig.JWKsURI == "" {
		return nil, errors.New("missing jwks_uri value")
	}

	return oidcConfig, nil
}

// Authenticate returns a nil error and the AuthClaims info (if available) if the subject is authenticated or a
func Authenticate(v Validator) openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		c := middleware.GetFiberContext(ctx)

		principal, err := v.Validate(input.RequestValidationInput.Request)
		if err != nil {
			return err
		}

		usrCtx := context.WithValue(c.UserContext(), jwtToken, principal)
		// nolint:contextcheck
		c.SetUserContext(usrCtx)

		return nil
	}
}

// Validate validates the provided token.
func (oidc *RemoteOidcValidator) Validate(req *http.Request) (*authn.AuthClaims, error) {
	jwtParser := jwt.NewParser(
		jwt.WithValidMethods([]string{"RS256"}),
		jwt.WithIssuedAt(),
		jwt.WithExpirationRequired(),
	)

	// Now, we need to get the JWS from the request, to match the request expectations
	// against request contents.
	jws, err := GetJWSFromRequest(req)
	if err != nil {
		return nil, fmt.Errorf("getting jws: %w", err)
	}

	token, err := jwtParser.Parse(jws, func(token *jwt.Token) (any, error) {
		return oidc.JWKs.Keyfunc(token)
	})

	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrClaimsInvalid
	}

	validIssuers := []string{
		oidc.MainIssuer,
	}
	validIssuers = append(validIssuers, oidc.IssuerAliases...)

	ok = slices.ContainsFunc(validIssuers, func(issuer string) bool {
		v := jwt.NewValidator(jwt.WithIssuer(issuer))
		err := v.Validate(claims)
		return err == nil
	})

	if !ok {
		return nil, ErrInvalidIssuer
	}

	if err := jwt.NewValidator(jwt.WithAudience(oidc.Audience)).Validate(claims); err != nil {
		return nil, ErrInvalidAudiance
	}

	// optional subject
	subject := ""
	if subjectClaim, ok := claims["sub"]; ok {
		if subject, ok = subjectClaim.(string); !ok {
			return nil, ErrInvalidSubject
		}
	}

	principal := &authn.AuthClaims{
		Subject: subject,
		Scopes:  make(map[string]bool),
	}

	// optional scopes
	if scopeKey, ok := claims["scope"]; ok {
		if scope, ok := scopeKey.(string); ok {
			scopes := strings.Split(scope, " ")
			for _, s := range scopes {
				principal.Scopes[s] = true
			}
		}
	}

	return principal, nil
}

func (oidc *RemoteOidcValidator) Close() {
	oidc.JWKs.EndBackground()
}

// GetJWSFromRequest extracts a JWS string from an Authorization: Bearer <jws> header
func GetJWSFromRequest(req *http.Request) (string, error) {
	authHdr := req.Header.Get("Authorization")
	// Check for the Authorization header.
	if authHdr == "" {
		return "", ErrNoAuthHeader
	}

	// We expect a header value of the form "Bearer <token>", with 1 space after
	// Bearer, per spec.
	prefix := "Bearer "
	if !strings.HasPrefix(authHdr, prefix) {
		return "", ErrInvalidAuthHeader
	}
	return strings.TrimPrefix(authHdr, prefix), nil
}

// GetJWTFromContext extracts the JWT token from the context.
func GetJWTFromContext(ctx context.Context) (*authn.AuthClaims, bool) {
	principal, ok := ctx.Value(jwtToken).(*authn.AuthClaims)

	return principal, ok
}
