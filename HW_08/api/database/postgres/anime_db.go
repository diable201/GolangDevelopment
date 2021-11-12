package postgres

import (
	"anime-database/api/database"
	"anime-database/api/models"
	"context"
	"github.com/jmoiron/sqlx"
)

func (db *DB) Anime() database.AnimeRepository {
	if db.anime == nil {
		db.anime = NewAnimeRepository(db.connection)
	}

	return db.anime
}

type AnimeRepository struct {
	connection *sqlx.DB
}

func NewAnimeRepository(connection *sqlx.DB) database.AnimeRepository {
	return &AnimeRepository{connection: connection}
}

func (m AnimeRepository) Create(_ context.Context, anime *models.Anime) error {
	_, err := m.connection.Exec("INSERT INTO anime(genre_id, title, titlejapanese, "+
		"source, episodes, kind, score, status, synopsis) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", anime.GenreID, anime.Title,
		anime.TitleJapanese, anime.Source, anime.Episodes, anime.Kind, anime.Score,
		anime.Status, anime.Synopsis)
	if err != nil {
		return err
	}
	return nil
}

func (m AnimeRepository) All(_ context.Context) ([]*models.Anime, error) {
	anime := make([]*models.Anime, 0)
	if err := m.connection.Select(&anime, "SELECT * FROM anime"); err != nil {
		return nil, err
	}

	return anime, nil
}

func (m AnimeRepository) ByID(_ context.Context, id int) (*models.Anime, error) {
	anime := new(models.Anime)
	if err := m.connection.Get(anime, "SELECT id, genre_id, title, titlejapanese, "+
		"source, episodes, kind, score, status, synopsis "+
		"FROM anime WHERE id=$1", id); err != nil {
		return nil, err
	}
	return anime, nil
}

func (m AnimeRepository) Update(_ context.Context, anime *models.Anime) error {
	_, err := m.connection.Exec("UPDATE anime SET genre_id = $1, title = $2, titlejapanese = $3, "+
		"source = $4, episodes = $5, kind = $6, score = $7, status = $8, synopsis = $9 WHERE id = $10",
		anime.GenreID, anime.Title, anime.TitleJapanese, anime.Source, anime.Episodes,
		anime.Kind, anime.Score, anime.Status, anime.Synopsis, anime.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m AnimeRepository) Delete(_ context.Context, id int) error {
	_, err := m.connection.Exec("DELETE FROM anime WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
