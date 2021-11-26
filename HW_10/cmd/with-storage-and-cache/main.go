package main

import (
	"anime-kafka/api/database"
	"anime-kafka/api/database/postgres"
	"anime-kafka/api/http"
	"anime-kafka/api/message_broker"
	"anime-kafka/api/message_broker/kafka"
	"context"
	lru "github.com/hashicorp/golang-lru"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go CatchTermination(cancel)

	dbURL := "postgres://localhost:5432/anime"
	library := postgres.NewDB()
	if err := library.Connect(dbURL); err != nil {
		panic(err)
	}
	defer func(store database.Library) {
		err := store.Close()
		if err != nil {
			panic(err)
		}
	}(library)

	twoQueueCache, err := lru.New2Q(6)
	if err != nil {
		panic(err)
	}

	brokers := []string{"localhost:29092"}
	broker := kafka.NewBroker(brokers, twoQueueCache, "peer2")
	if err := broker.Connect(ctx); err != nil {
		panic(err)
	}
	defer func(broker message_broker.MessageBroker) {
		err := broker.Close()
		if err != nil {

		}
	}(broker)

	server := http.NewServer(
		ctx,
		http.WithAddress(":8081"),
		http.WithLibrary(library),
		http.WithCache(twoQueueCache),
		http.WithBroker(broker),
	)

	if err := server.Run(); err != nil {
		log.Println(err)
	}

	server.WaitForGracefulTermination()
}

func CatchTermination(cancel context.CancelFunc) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Print("[WARN] caught termination signal")
	cancel()
}
