package cache

import (
	"anime-kafka/api/models"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type AnimeRedisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewAnimeRedisCache(host string, db int, expires time.Duration) AnimeCache {
	return &AnimeRedisCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}

func (a AnimeRedisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     a.host,
		Password: "",
		DB:       a.db,
	})
}

func (a AnimeRedisCache) Set(ctx context.Context, key string, value *models.Anime) {
	client := a.getClient()
	anime, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	_, err = client.Set(ctx, key, anime, a.expires*time.Second).Result()
	if err != nil {
		return
	}
}

func (a AnimeRedisCache) Get(ctx context.Context, key string) *models.Anime {
	client := a.getClient()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}
	anime := new(models.Anime)
	err = json.Unmarshal([]byte(val), &anime)
	if err != nil {
		panic(err)
	}
	return anime
}
