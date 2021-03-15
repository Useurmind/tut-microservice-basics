package database

import (
	"database/sql"

	"github.com/rs/zerolog/log"
	migrate "github.com/rubenv/sql-migrate"
)


func getMigrationsSource() *migrate.MemoryMigrationSource {
	migrations := &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			{
				Id:   "add_people_table",
				Up:   []string{"CREATE TABLE people (id int GENERATED ALWAYS AS IDENTITY, name text, age int)"},
				Down: []string{"DROP TABLE people"},
			},
		},
	}

	return migrations
}

func MigrateDatabase(db *sql.DB) error {
	res, err := migrate.Exec(db, "postgres", getMigrationsSource(), migrate.Up)
	if err != nil {
		log.Error().Err(err).Msg("Applying migrations failed")
		return err
	}
	log.Info().Int("appliedMigrations", res).Msg("Applied migrations")

	return nil
}