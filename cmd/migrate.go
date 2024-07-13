package cmd

import (
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/knox/internal/adapters/database"
	"github.com/zeiss/knox/internal/models"

	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := gorm.Open(postgres.Open(config.Flags.DatabaseURI), &gorm.Config{
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

		return store.Migrate(
			cmd.Context(),
			&models.Environment{},
			&models.Lock{},
			&models.Project{},
			&models.Snapshot{},
			&models.State{},
		)
	},
}
