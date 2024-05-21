package migration

import (
	"database/sql"

	"github.com/DavidHuie/gomigrate"
)

func ApplyMigration(db *sql.DB) error {
	migrator, err := gomigrate.NewMigrator(db, gomigrate.Postgres{}, "./build/migrations")
	if err != nil {
		return err
	}
	err = migrator.Migrate()
	if err != nil {
		return err
	}
	return nil
}
