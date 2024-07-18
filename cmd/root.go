package cmd

import (
	"context"
	"log"

	"github.com/zeiss/knox/internal/adapters/database"
	"github.com/zeiss/knox/internal/adapters/handlers"
	"github.com/zeiss/knox/internal/controllers"
	openapi "github.com/zeiss/knox/pkg/apis"
	"github.com/zeiss/knox/pkg/auth"
	"github.com/zeiss/knox/pkg/cfg"
	"github.com/zeiss/knox/pkg/oas"
	"github.com/zeiss/knox/pkg/utils"

	"github.com/gofiber/fiber/v2"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
	requestid "github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/katallaxie/pkg/server"
	"github.com/kelseyhightower/envconfig"
	middleware "github.com/oapi-codegen/fiber-middleware"
	openfga "github.com/openfga/go-sdk/client"
	"github.com/spf13/cobra"
	seed "github.com/zeiss/gorm-seed"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var config *cfg.Config

func init() {
	config = cfg.New()

	err := envconfig.Process("", config.Flags)
	if err != nil {
		log.Fatal(err)
	}

	Root.AddCommand(Seed)
	Root.AddCommand(Migrate)

	Root.PersistentFlags().StringVar(&config.Flags.Addr, "addr", config.Flags.Addr, "addr")
	Root.PersistentFlags().StringVar(&config.Flags.DatabaseURI, "db-uri", config.Flags.DatabaseURI, "Database URI")
	Root.PersistentFlags().StringVar(&config.Flags.DatabaseTablePrefix, "db-table-prefix", config.Flags.DatabaseTablePrefix, "Database table prefix")
	Root.PersistentFlags().StringVar(&config.Flags.FGAApiUrl, "fga-api-url", config.Flags.FGAApiUrl, "FGA API URL")
	Root.PersistentFlags().StringVar(&config.Flags.FGAStoreID, "fga-store-id", config.Flags.FGAStoreID, "FGA Store ID")
	Root.PersistentFlags().StringVar(&config.Flags.FGAAuthorizationModelID, "fga-authorization-model-id", config.Flags.FGAAuthorizationModelID, "FGA Authorization Model ID")

	Root.SilenceUsage = true
}

var Root = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		srv := NewWebSrv(config)

		s, _ := server.WithContext(cmd.Context())
		s.Listen(srv, false)

		return s.Wait()
	},
}

var _ server.Listener = (*WebSrv)(nil)

// WebSrv is the server that implements the Noop interface.
type WebSrv struct {
	cfg *cfg.Config
}

// NewWebSrv returns a new instance of NoopSrv.
func NewWebSrv(cfg *cfg.Config) *WebSrv {
	return &WebSrv{cfg}
}

// Start starts the server.
func (s *WebSrv) Start(ctx context.Context, ready server.ReadyFunc, run server.RunFunc) func() error {
	return func() error {
		conn, err := gorm.Open(postgres.Open(s.cfg.Flags.DatabaseURI), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: config.Flags.DatabaseTablePrefix,
			},
		})
		if err != nil {
			return err
		}

		store, err := seed.NewDatabase(conn, database.NewReadTx(), database.NewWriteTx())
		if err != nil {
			return err
		}

		fga, err := openfga.NewSdkClient(
			&openfga.ClientConfiguration{
				ApiUrl:               s.cfg.Flags.FGAApiUrl,
				StoreId:              s.cfg.Flags.FGAStoreID,
				AuthorizationModelId: s.cfg.Flags.FGAAuthorizationModelID,
			},
		)
		if err != nil {
			return err
		}

		swagger, err := openapi.GetSwagger()
		if err != nil {
			return err
		}
		swagger.Servers = nil

		c := fiber.Config{
			ErrorHandler: utils.DefaultErrorHandler,
		}

		app := fiber.New(c)
		app.Use(requestid.New())
		app.Use(logger.New())

		validatorOptions := &middleware.Options{}
		validatorOptions.Options.AuthenticationFunc = auth.NewAuthenticator(auth.WithBasicAuthenticator(auth.NewBasicAuthenticator(store)))

		// validatorOptions.ErrorHandler = authz.NewOpenAPIErrorHandler()

		app.Use(middleware.OapiRequestValidatorWithOptions(swagger, validatorOptions))

		lc := controllers.NewLocksController(store)
		sc := controllers.NewStateController(store)
		ssc := controllers.NewSnapshotController(store)
		tc := controllers.NewTeamController(store)
		pc := controllers.NewProjectController(store)
		ec := controllers.NewEnvironmentController(store)

		authz := oas.NewAuthz(
			oas.Config{
				Checker: oas.NewChecker(fga),
				Resolvers: map[string]oas.AuthzResolverFunc{
					"GetTeam": func(c *fiber.Ctx) (oas.User, oas.Relation, oas.Object, error) {
						return oas.NoopUser, oas.NoopRelation, oas.NoopObject, nil
					},
				},
			},
		)

		handlers := handlers.NewAPIHandlers(lc, sc, ssc, tc, pc, ec)
		handler := openapi.NewStrictHandler(handlers, []openapi.StrictMiddlewareFunc{authz})
		openapi.RegisterHandlers(app, handler)

		err = app.Listen(s.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
