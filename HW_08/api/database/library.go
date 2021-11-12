package database

import (
	"anime-database/api/models"
	"context"
)

type Library interface {
	Connect(url string) error
	Close() error
	Anime() AnimeRepository
	Manga() MangaRepository
	Genres() GenresRepository
}

type AnimeRepository interface {
	Create(ctx context.Context, anime *models.Anime) error
	All(ctx context.Context) ([]*models.Anime, error)
	ByID(ctx context.Context, id int) (*models.Anime, error)
	Update(ctx context.Context, anime *models.Anime) error
	Delete(ctx context.Context, id int) error
}

type MangaRepository interface {
	Create(ctx context.Context, manga *models.Manga) error
	All(ctx context.Context) ([]*models.Manga, error)
	ByID(ctx context.Context, id int) (*models.Manga, error)
	Update(ctx context.Context, manga *models.Manga) error
	Delete(ctx context.Context, id int) error
}

type GenresRepository interface {
	Create(ctx context.Context, genre *models.Genre) error
	All(ctx context.Context) ([]*models.Genre, error)
	ByID(ctx context.Context, id int) (*models.Genre, error)
	Update(ctx context.Context, genre *models.Genre) error
	Delete(ctx context.Context, id int) error
}
