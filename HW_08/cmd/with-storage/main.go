package main

import (
	"anime-database/api/database"
	"anime-database/api/database/postgres"
	"anime-database/api/http"
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

	server := http.NewServer(context.Background(), library, ":8080")
	if err := server.Run(); err != nil {
		panic(err)
	}

	server.WaitForGracefulTermination()
}
