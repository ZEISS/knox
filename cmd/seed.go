package cmd

import (
	"github.com/zeiss/fiber-goth/adapters"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/knox/internal/models"

	"github.com/spf13/cobra"
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

var Seed = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := gorm.Open(postgres.Open(config.Flags.DatabaseURI), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: config.Flags.DatabaseTablePrefix,
			},
		})
		if err != nil {
			return err
		}

		seeder := seed.NewSeeder(conn)
		return seeder.Seed(cmd.Context(), seeds...)
	},
}
