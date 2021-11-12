package postgres

import (
	"anime-database/api/database"
	"anime-database/api/models"
	"context"
	"github.com/jmoiron/sqlx"
)

func (db *DB) Genres() database.GenresRepository {
	if db.genres == nil {
		db.genres = NewGenresRepository(db.connection)
	}

	return db.genres
}

type GenresRepository struct {
	connection *sqlx.DB
}

func NewGenresRepository(connection *sqlx.DB) database.GenresRepository {
	return &GenresRepository{connection: connection}
}

func (g GenresRepository) Create(_ context.Context, genre *models.Genre) error {
	_, err := g.connection.Exec("INSERT INTO genres(name) VALUES ($1)", genre.Name)
	if err != nil {
		return err
	}

	return nil
}

func (g GenresRepository) All(_ context.Context) ([]*models.Genre, error) {
	genres := make([]*models.Genre, 0)
	if err := g.connection.Select(&genres, "SELECT * FROM genres"); err != nil {
		return nil, err
	}

	return genres, nil
}

func (g GenresRepository) ByID(_ context.Context, id int) (*models.Genre, error) {
	genre := new(models.Genre)
	if err := g.connection.Get(genre, "SELECT id, name FROM genres WHERE id=$1", id); err != nil {
		return nil, err
	}

	return genre, nil
}

func (g GenresRepository) Update(_ context.Context, genre *models.Genre) error {
	_, err := g.connection.Exec("UPDATE genres SET name = $1 WHERE id = $2", genre.Name, genre.ID)
	if err != nil {
		return err
	}

	return nil
}

func (g GenresRepository) Delete(_ context.Context, id int) error {
	_, err := g.connection.Exec("DELETE FROM genres WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
