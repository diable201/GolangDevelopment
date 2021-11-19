package cache

import (
	"anime-redis/api/models"
	"context"
)

type Cache interface {
	Anime() AnimeCache
	Manga() MangaCache
}

type AnimeCache interface {
	Set(ctx context.Context, key string, value *models.Anime)
	Get(ctx context.Context, key string) *models.Anime
}

type MangaCache interface {
	Set(ctx context.Context, key string, value *models.Manga)
	Get(ctx context.Context, key string) *models.Manga
}
