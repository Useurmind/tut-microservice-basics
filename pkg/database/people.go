package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/useurmind/tut-microservice-basics/pkg/model"
)

func GetPeople(ctx context.Context, db *sqlx.DB) ([]*model.People, error) {
	people := []*model.People{}
	err := db.SelectContext(ctx, &people, "SELECT * FROM people ORDER BY name ASC")
	if err != nil {
		return nil, err
	}

	return people, nil
}

func InsertPeople(ctx context.Context, db *sqlx.DB, people *model.People) (*model.People, error) {	
	row := db.QueryRowContext(ctx, "INSERT INTO people (name, age) VALUES ($1, $2) RETURNING id", people.Name, people.Age)
	peopleId := int64(-1)
	err := row.Scan(&peopleId)
	if err != nil {
		return nil, err
	}

	people.ID = peopleId

	return people, nil
}

func GetPeopleById(ctx context.Context, db *sqlx.DB, id int) (*model.People, error) {
	log.Info().Int("id", id).Msg("get people by id")
	people := model.People{}
	err := db.GetContext(ctx, &people, "SELECT * FROM people WHERE id=$1", id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &people, nil
}

func DeletePeopleById(ctx context.Context, db *sqlx.DB, id int) (bool, error) {
	log.Info().Int("id", id).Msg("delete people by id")
	row := db.QueryRowContext(ctx, "DELETE FROM people WHERE id=$1 RETURNING id", id)
	peopleId := int64(-1)
	err := row.Scan(&peopleId)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
