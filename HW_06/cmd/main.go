package main

import (
	"anime/api/database/inmemory"
	"anime/api/http"
	"context"
	"log"
)

func main() {
	library := inmemory.NewDB()
	server := http.NewServer(context.Background(), library, ":8080")
	if err := server.Run(); err != nil {
		log.Println(err)
	}
	server.WaitForGracefulTermination()
}
