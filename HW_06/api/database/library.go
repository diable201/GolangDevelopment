package database

import (
	"anime/api/models"
	"context"
)

type Library interface {
	Create(ctx context.Context, anime *models.Anime) error
	All(ctx context.Context) ([]*models.Anime, error)
	ByID(ctx context.Context, id int) (*models.Anime, error)
	Update(ctx context.Context, anime *models.Anime) error
	Delete(ctx context.Context, id int) error
}
