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

func MigrateDatabase() error {
	db, err := GetConnection()
	if err != nil {
		log.Error().Err(err).Msg("Could not connect to database")
		return err
	}
	defer db.Close()
	err = migrateDatabase(db.DB)
	if err != nil {
		log.Error().Err(err).Msg("Could not migrate database")
		return err
	}

	return nil
}

func migrateDatabase(db *sql.DB) error {
	res, err := migrate.Exec(db, "postgres", getMigrationsSource(), migrate.Up)
	if err != nil {
		log.Error().Err(err).Msg("Applying migrations failed")
		return err
	}
	log.Info().Int("appliedMigrations", res).Msg("Applied migrations")

	return nil
}