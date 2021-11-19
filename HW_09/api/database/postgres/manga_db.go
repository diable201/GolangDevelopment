package postgres

import (
	"anime-redis/api/database"
	"anime-redis/api/models"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func (db *DB) Manga() database.MangaRepository {
	if db.manga == nil {
		db.manga = NewMangaRepository(db.connection)
	}

	return db.manga
}

type MangaRepository struct {
	connection *sqlx.DB
}

func NewMangaRepository(connection *sqlx.DB) database.MangaRepository {
	return &MangaRepository{connection: connection}
}

func (m MangaRepository) Create(_ context.Context, manga *models.Manga) error {
	_, err := m.connection.Exec("INSERT INTO manga(genre_id, title, titlejapanese, "+
		"volumes, chapters, score, status, synopsis) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", manga.GenreID, manga.Title,
		manga.TitleJapanese, manga.Volumes, manga.Chapters, manga.Score,
		manga.Status, manga.Synopsis)
	if err != nil {
		return err
	}
	return nil
}

func (m MangaRepository) All(_ context.Context, filter *models.MangaFilter) ([]*models.Manga, error) {
	manga := make([]*models.Manga, 0)
	basicQuery := "SELECT * FROM manga"
	if filter.Query != nil {
		basicQuery = fmt.Sprintf("%s WHERE title ILIKE $1 OR titlejapanese ILIKE $2", basicQuery)
		if err := m.connection.Select(&manga, basicQuery,
			"%"+*filter.Query+"%", "%"+*filter.Query+"%"); err != nil {
			return nil, err
		}
		return manga, nil
	}
	if err := m.connection.Select(&manga, basicQuery); err != nil {
		return nil, err
	}
	return manga, nil
}

func (m MangaRepository) ByID(_ context.Context, id int) (*models.Manga, error) {
	manga := new(models.Manga)
	if err := m.connection.Get(manga, "SELECT id, genre_id, title, titlejapanese, "+
		"volumes, chapters, score, status, synopsis "+
		"FROM manga WHERE id=$1", id); err != nil {
		return nil, err
	}
	return manga, nil
}

func (m MangaRepository) Update(_ context.Context, manga *models.Manga) error {
	_, err := m.connection.Exec("UPDATE manga SET genre_id = $1, title = $2, titlejapanese = $3, "+
		"volumes = $4, chapters = $5, score = $6, status = $7, synopsis = $8 WHERE id = $9",
		manga.GenreID, manga.Title, manga.TitleJapanese, manga.Volumes, manga.Chapters,
		manga.Score, manga.Status, manga.Synopsis, manga.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m MangaRepository) Delete(_ context.Context, id int) error {
	_, err := m.connection.Exec("DELETE FROM manga WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
