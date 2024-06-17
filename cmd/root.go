package cmd

import (
	"context"

	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/knox/internal/adapters/database"
	"github.com/zeiss/knox/internal/adapters/handlers"
	"github.com/zeiss/knox/internal/controllers"
	openapi "github.com/zeiss/knox/pkg/apis"
	"github.com/zeiss/knox/pkg/cfg"
	"github.com/zeiss/knox/pkg/utils"

	"github.com/gofiber/fiber/v2"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
	requestid "github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/katallaxie/pkg/server"
	middleware "github.com/oapi-codegen/fiber-middleware"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var config *cfg.Config

func init() {
	config = cfg.New()

	Root.PersistentFlags().StringVar(&config.Flags.Addr, "addr", ":8080", "addr")
	Root.PersistentFlags().StringVar(&config.Flags.DB.Addr, "db-addr", config.Flags.DB.Addr, "Database address")
	Root.PersistentFlags().StringVar(&config.Flags.DB.Database, "db-database", config.Flags.DB.Database, "Database name")
	Root.PersistentFlags().StringVar(&config.Flags.DB.Username, "db-username", config.Flags.DB.Username, "Database user")
	Root.PersistentFlags().StringVar(&config.Flags.DB.Password, "db-password", config.Flags.DB.Password, "Database password")
	Root.PersistentFlags().IntVar(&config.Flags.DB.Port, "db-port", config.Flags.DB.Port, "Database port")

	Root.SilenceUsage = true
}

var Root = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := cfg.New()
		srv := NewWebSrv(cfg)

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
		conn, err := gorm.Open(postgres.Open(s.cfg.DSN()), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: "knox_",
			},
		})
		if err != nil {
			return err
		}

		db, err := database.NewDB(conn)
		if err != nil {
			return err
		}

		err = db.Migrate(ctx)
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
		validatorOptions.Options.AuthenticationFunc = authz.NewOpenAPIAuthenticator(authz.WithAuthzChecker(authz.NewFake(true)))
		validatorOptions.ErrorHandler = authz.NewOpenAPIErrorHandler()

		app.Use(middleware.OapiRequestValidatorWithOptions(swagger, validatorOptions))

		lc := controllers.NewLocksController(db)
		sc := controllers.NewStateController(db)

		handlers := handlers.NewAPIHandlers(lc, sc)
		handler := openapi.NewStrictHandler(handlers, nil)
		openapi.RegisterHandlers(app, handler)

		err = app.Listen(s.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
