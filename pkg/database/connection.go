package database

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)


func GetConnection() (*sqlx.DB, error) {
	dbUser, err := getEnvRequired("DB_USER")
	if err != nil {
		return nil, err
	}
	dbPass, err := getEnvRequired("DB_PASS")
	if err != nil {
		return nil, err
	}
	dbHost, err := getEnvRequired("DB_HOST")
	if err != nil {
		return nil, err
	}
	dbPort, err := getEnvRequired("DB_PORT")
	if err != nil {
		return nil, err
	}
	dbName, err := getEnvRequired("DB_NAME")
	if err != nil {
		return nil, err
	}

	connString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbName, dbUser, dbPass)

	log.Info().Str("host", dbHost).Str("port", dbPort).Str("dbName", dbName).Str("user", dbUser).Msg("Connect to database")
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, err
	}

	return db, err
}

func getEnvRequired(env string) (string, error) {
	value := os.Getenv(env)
	if value == "" {
		return "", fmt.Errorf("%s must be set, but was empty", env)
	}
	return value, nil
}