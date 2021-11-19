package cache

import (
	"anime-redis/api/models"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type MangaRedisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewMangaRedisCache(host string, db int, expires time.Duration) MangaCache {
	return &MangaRedisCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}

func (m MangaRedisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     m.host,
		Password: "",
		DB:       m.db,
	})
}

func (m MangaRedisCache) Set(ctx context.Context, key string, value *models.Manga) {
	client := m.getClient()
	manga, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	_, err = client.Set(ctx, key, manga, m.expires*time.Second).Result()
	if err != nil {
		return
	}
}

func (m MangaRedisCache) Get(ctx context.Context, key string) *models.Manga {
	client := m.getClient()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}
	manga := new(models.Manga)
	err = json.Unmarshal([]byte(val), &manga)
	if err != nil {
		panic(err)
	}
	return manga
}
