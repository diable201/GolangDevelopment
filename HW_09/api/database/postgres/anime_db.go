package postgres

import (
	"anime-redis/api/database"
	"anime-redis/api/models"
	"context"
	"fmt"
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

func (a AnimeRepository) Create(_ context.Context, anime *models.Anime) error {
	_, err := a.connection.Exec("INSERT INTO anime(genre_id, title, titlejapanese, "+
		"source, episodes, kind, score, status, synopsis) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", anime.GenreID, anime.Title,
		anime.TitleJapanese, anime.Source, anime.Episodes, anime.Kind, anime.Score,
		anime.Status, anime.Synopsis)
	if err != nil {
		return err
	}
	return nil
}

func (a AnimeRepository) All(_ context.Context, filter *models.AnimeFilter) ([]*models.Anime, error) {
	anime := make([]*models.Anime, 0)
	basicQuery := "SELECT * FROM anime"
	if filter.Query != nil {
		basicQuery = fmt.Sprintf("%s WHERE title ILIKE $1 OR titlejapanese ILIKE $2", basicQuery)
		if err := a.connection.Select(&anime, basicQuery,
			"%"+*filter.Query+"%", "%"+*filter.Query+"%"); err != nil {
			return nil, err
		}
		return anime, nil
	}
	if err := a.connection.Select(&anime, basicQuery); err != nil {
		return nil, err
	}
	return anime, nil
}

func (a AnimeRepository) ByID(_ context.Context, id int) (*models.Anime, error) {
	anime := new(models.Anime)
	if err := a.connection.Get(anime, "SELECT id, genre_id, title, titlejapanese, "+
		"source, episodes, kind, score, status, synopsis "+
		"FROM anime WHERE id = $1", id); err != nil {
		return nil, err
	}
	return anime, nil
}

func (a AnimeRepository) Update(_ context.Context, anime *models.Anime) error {
	_, err := a.connection.Exec("UPDATE anime SET genre_id = $1, title = $2, titlejapanese = $3, "+
		"source = $4, episodes = $5, kind = $6, score = $7, status = $8, synopsis = $9 WHERE id = $10",
		anime.GenreID, anime.Title, anime.TitleJapanese, anime.Source, anime.Episodes,
		anime.Kind, anime.Score, anime.Status, anime.Synopsis, anime.ID)
	if err != nil {
		return err
	}

	return nil
}

func (a AnimeRepository) Delete(_ context.Context, id int) error {
	_, err := a.connection.Exec("DELETE FROM anime WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
