package http

import (
	"anime/api/database"
	"anime/api/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	ctx        context.Context
	idleConsCh chan struct{}
	library    database.Library
	address    string
}

func NewServer(ctx context.Context, library database.Library, address string) *Server {
	return &Server{
		ctx:        ctx,
		idleConsCh: make(chan struct{}),
		library:    library,
		address:    address,
	}
}

func (s *Server) basicHandler() chi.Router {
	r := chi.NewRouter()

	r.Get("/anime", func(writer http.ResponseWriter, request *http.Request) {
		anime, err := s.library.Anime().All(request.Context())
		if err != nil {
			fmt.Fprintf(writer, "Unknown error: %v", err)
			return
		}
		render.JSON(writer, request, anime)
	})

	r.Get("/anime/{id}", func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(writer, "Unknown error: %v", err)
			return
		}

		anime, err := s.library.Anime().ByID(request.Context(), id)
		if err != nil {
			fmt.Fprintf(writer, "Unknown error: %v", err)
			return
		}
		render.JSON(writer, request, anime)
	})

	r.Post("/anime", func(writer http.ResponseWriter, request *http.Request) {
		anime := new(models.Anime)
		if err := json.NewDecoder(request.Body).Decode(anime); err != nil {
			fmt.Fprintf(writer, "Unknown error: %v", err)
			return
		}
		s.library.Anime().Create(request.Context(), anime)
	})

	r.Put("/anime", func(writer http.ResponseWriter, request *http.Request) {
		anime := new(models.Anime)
		if err := json.NewDecoder(request.Body).Decode(anime); err != nil {
			fmt.Fprintf(writer, "Unknown error: %v", err)
			return
		}
		s.library.Anime().Update(request.Context(), anime)
	})

	r.Delete("/anime/{id}", func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(writer, "Unknown error: %v", err)
			return
		}
		s.library.Anime().Delete(request.Context(), id)
	})

	r.Get("/manga", func(writer http.ResponseWriter, request *http.Request) {
		manga, err := s.library.Manga().All(request.Context())
		if err != nil {
			fmt.Fprintf(writer, "Unknown error: %v", err)
			return
		}
		render.JSON(writer, request, manga)
	})

	r.Get("/manga/{id}", func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(writer, "Unknown error: %v", err)
			return
		}

		manga, err := s.library.Manga().ByID(request.Context(), id)
		if err != nil {
			fmt.Fprintf(writer, "Unknown error: %v", err)
			return
		}
		render.JSON(writer, request, manga)
	})

	r.Post("/manga", func(writer http.ResponseWriter, request *http.Request) {
		manga := new(models.Manga)
		if err := json.NewDecoder(request.Body).Decode(manga); err != nil {
			fmt.Fprintf(writer, "Unknown error: %v", err)
			return
		}
		s.library.Manga().Create(request.Context(), manga)
	})

	r.Put("/manga", func(writer http.ResponseWriter, request *http.Request) {
		manga := new(models.Manga)
		if err := json.NewDecoder(request.Body).Decode(manga); err != nil {
			fmt.Fprintf(writer, "Unknown error: %v", err)
			return
		}
		s.library.Manga().Update(request.Context(), manga)
	})

	r.Delete("/manga/{id}", func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(writer, "Unknown error: %v", err)
			return
		}
		s.library.Manga().Delete(request.Context(), id)
	})

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
