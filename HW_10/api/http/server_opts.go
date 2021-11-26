package http

import (
	_ "anime-kafka/api/cache"
	"anime-kafka/api/database"
	"anime-kafka/api/message_broker"
	lru "github.com/hashicorp/golang-lru"
)

type ServerOption func(server *Server)

func WithAddress(address string) ServerOption {
	return func(server *Server) {
		server.address = address
	}
}

func WithLibrary(library database.Library) ServerOption {
	return func(server *Server) {
		server.library = library
	}
}

func WithCache(cache *lru.TwoQueueCache) ServerOption {
	return func(server *Server) {
		server.cache = cache
	}
}

func WithBroker(broker message_broker.MessageBroker) ServerOption {
	return func(server *Server) {
		server.broker = broker
	}
}
