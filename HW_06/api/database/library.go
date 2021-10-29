package database

import (
	"anime/api/models"
	"context"
)

type Library interface {
	Anime() AnimeRepository
	Manga() MangaRepository
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
