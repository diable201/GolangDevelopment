package http

import (
	"anime-redis/api/cache"
	"anime-redis/api/database"
	"context"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"time"
)

type Server struct {
	ctx        context.Context
	idleConsCh chan struct{}
	library    database.Library
	cacheAnime cache.AnimeCache
	cacheManga cache.MangaCache
	address    string
}

func NewServer(ctx context.Context, opts ...ServerOption) *Server {
	server := &Server{
		ctx:        ctx,
		idleConsCh: make(chan struct{}),
	}

	for _, opt := range opts {
		opt(server)
	}

	return server
}

func (s *Server) basicHandler() chi.Router {
	r := chi.NewRouter()
	animeResource := NewAnimeResource(s.library, s.cacheAnime)
	mangaResource := NewMangaResource(s.library, s.cacheManga)
	genresResource := NewGenresResource(s.library)
	r.Mount("/anime", animeResource.Routes())
	r.Mount("/manga", mangaResource.Routes())
	r.Mount("/genres", genresResource.Routes())
	return r
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:         s.address,
		Handler:      s.basicHandler(),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 30,
	}
	go s.ListenCtxForGt(server)
	log.Println("[HTTP] Server running on", s.address)
	return server.ListenAndServe()
}

func (s *Server) ListenCtxForGt(srv *http.Server) {
	<-s.ctx.Done()
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("[HTTP] Got error while shutting down %v", err)
	}
	log.Println("[HTTP] Processed all idle connections")
	close(s.idleConsCh)
}

func (s *Server) WaitForGracefulTermination() {
	<-s.idleConsCh
}
