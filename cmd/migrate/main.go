package main

import (
	"context"
	"log"
	"os"

	"github.com/zeiss/fiber-goth/adapters"
	"github.com/zeiss/knox/internal/adapters/database"
	"github.com/zeiss/knox/internal/configs"
	"github.com/zeiss/knox/internal/models"

	"github.com/katallaxie/pkg/logger"
	"github.com/spf13/cobra"
	seed "github.com/zeiss/gorm-seed"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var seeds = []seed.Seed{
	{
		Name: "create team",
		Run: func(db *gorm.DB) error {
			team := adapters.GothTeam{
				Name: "Zeiss",
				Slug: "zeiss",
			}

			err := db.Create(&team).Error
			if err != nil {
				return err
			}

			project := models.Project{
				Name:   "demo",
				TeamID: team.ID,
			}

			err = db.Create(&project).Error
			if err != nil {
				return err
			}

			return db.Create([]models.Environment{
				{
					ProjectID: project.ID,
					Name:      "dev",
				},
			}).Error
		},
	},
}

var cfg = configs.New()

var rootCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(cmd.Context())
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfg.Flags.Addr, "addr", ":8080", "addr")
	rootCmd.PersistentFlags().StringVar(&cfg.Flags.DB.Database, "db-database", cfg.Flags.DB.Database, "Database name")
	rootCmd.PersistentFlags().StringVar(&cfg.Flags.DB.Username, "db-username", cfg.Flags.DB.Username, "Database user")
	rootCmd.PersistentFlags().StringVar(&cfg.Flags.DB.Password, "db-password", cfg.Flags.DB.Password, "Database password")
	rootCmd.PersistentFlags().IntVar(&cfg.Flags.DB.Port, "db-port", cfg.Flags.DB.Port, "Database port")

	rootCmd.SilenceUsage = true
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	logger.RedirectStdLog(logger.LogSink)

	dsn := "host=host.docker.internal user=example password=example dbname=example port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
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

	seeder := seed.NewSeeder(conn)
	err = seeder.Seed(ctx, seeds...)
	if err != nil {
		panic(err)
	}

	return nil
}
