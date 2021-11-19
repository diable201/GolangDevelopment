package main

import (
	"anime-redis/api/cache"
	"anime-redis/api/database"
	"anime-redis/api/database/postgres"
	"anime-redis/api/http"
	"context"
	_ "github.com/lib/pq"
)

func main() {
	url := "postgres://localhost:5432/anime"
	library := postgres.NewDB()
	if err := library.Connect(url); err != nil {
		panic(err)
	}
	defer func(store database.Library) {
		err := store.Close()
		if err != nil {
			panic(err)
		}
	}(library)

	cacheAnime := cache.NewAnimeRedisCache("localhost:6379", 0, 30)
	cacheManga := cache.NewMangaRedisCache("localhost:6379", 0, 30)

	server := http.NewServer(
		context.Background(),
		http.WithAddress(":8080"),
		http.WithLibrary(library),
		http.WithCache(cacheAnime, cacheManga),
	)
	if err := server.Run(); err != nil {
		panic(err)
	}

	server.WaitForGracefulTermination()
}
