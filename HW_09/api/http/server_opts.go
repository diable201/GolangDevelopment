package http

import (
	"anime-redis/api/cache"
	"anime-redis/api/database"
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

func WithCache(cacheAnime cache.AnimeCache, cacheManga cache.MangaCache) ServerOption {
	return func(server *Server) {
		server.cacheAnime = cacheAnime
		server.cacheManga = cacheManga
	}
}
