package migrate

import (
	"basketManager/config"
	"basketManager/db"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Imported for its side effects
)

const (
	flagPath = "path"
)

func main(path string, cfg config.Config) error {
	database := db.WithRetry(db.Create, cfg.Postgres)

	defer func() {
		if err := database.Close(); err != nil {
			logrus.Error(err.Error())
		}
	}()

	driver, err := postgres.WithInstance(database.DB(), &postgres.Config{})
	if err != nil {
		return err
	}

	dbName := cfg.Postgres.DBName

	m, err := migrate.NewWithDatabaseInstance("file://"+path, dbName, driver)
	if err != nil {
		return fmt.Errorf("migration: %s", err)
	}

	if err := m.Up(); errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("No change detected. All migrations have already been applied!")
		return nil
	} else if err != nil {
		return fmt.Errorf("migration up: %s", err)
	}

	return nil
}

// Register migrate command.
func Register(root *cobra.Command, cfg config.Config) {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Provides DB migration functionality",

		PreRunE: func(cmd *cobra.Command, args []string) error {
			path, err := cmd.Flags().GetString(flagPath)
			if err != nil {
				return fmt.Errorf("error parsing %s flag: %w", flagPath, err)
			}

			if path == "" {
				return fmt.Errorf("%s flag is required", flagPath)
			}

			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := cmd.Flags().GetString(flagPath)
			if err != nil {
				return err
			}

			if err := main(path, cfg); err != nil {
				return err
			}

			cmd.Println("migrations ran successfully")

			return nil
		},
	}

	cmd.Flags().StringP(flagPath, "p", "", "migration folder path")

	root.AddCommand(cmd)
}
