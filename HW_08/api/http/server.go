package http

import (
	"anime-database/api/database"
	"anime-database/api/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-ozzo/ozzo-validation/v4"
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
			writer.WriteHeader(http.StatusInternalServerError)
			_, err := fmt.Fprintf(writer, "DB err: %v", err)
			if err != nil {
				return
			}
			return
		}
		render.JSON(writer, request, anime)
	})

	r.Get("/anime/{id}", func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			_, err := fmt.Fprintf(writer, "Unknown err: %v", err)
			if err != nil {
				return
			}
			return
		}

		anime, err := s.library.Anime().ByID(request.Context(), id)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, err := fmt.Fprintf(writer, "DB err: %v", err)
			if err != nil {
				return
			}
			return
		}
		render.JSON(writer, request, anime)
	})

	r.Post("/anime", func(writer http.ResponseWriter, request *http.Request) {
		anime := new(models.Anime)
		if err := json.NewDecoder(request.Body).Decode(anime); err != nil {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			_, err := fmt.Fprintf(writer, "Unknown error: %v", err)
			if err != nil {
				return
			}
			return
		}
		if err := s.library.Anime().Create(request.Context(), anime); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, err := fmt.Fprintf(writer, "DB err: %v", err)
			if err != nil {
				return
			}
			return
		}
		writer.WriteHeader(http.StatusCreated)
	})

	r.Put("/anime", func(writer http.ResponseWriter, request *http.Request) {
		anime := new(models.Anime)
		if err := json.NewDecoder(request.Body).Decode(anime); err != nil {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			_, err := fmt.Fprintf(writer, "Unknown err: %v", err)
			if err != nil {
				return
			}
			return
		}

		err := validation.ValidateStruct(
			anime,
			validation.Field(&anime.ID, validation.Required),
			validation.Field(&anime.GenreID, validation.Required),
			validation.Field(&anime.Title, validation.Required),
			validation.Field(&anime.TitleJapanese, validation.Required),
			validation.Field(&anime.Source, validation.Required),
			validation.Field(&anime.Episodes, validation.Required),
			validation.Field(&anime.Kind, validation.Required),
			validation.Field(&anime.Score, validation.Required),
			validation.Field(&anime.Status, validation.Required),
			validation.Field(&anime.Synopsis, validation.Required),
		)

		if err != nil {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			_, err := fmt.Fprintf(writer, "Unknown err: %v", err)
			if err != nil {
				return
			}
			return
		}

		if err := s.library.Anime().Update(request.Context(), anime); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, err := fmt.Fprintf(writer, "DB err: %v", err)
			if err != nil {
				return
			}
			return
		}
	})

	r.Delete("/anime/{id}", func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			_, err := fmt.Fprintf(writer, "Unknown err: %v", err)
			if err != nil {
				return
			}
			return
		}
		if err := s.library.Anime().Delete(request.Context(), id); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, err := fmt.Fprintf(writer, "DB err: %v", err)
			if err != nil {
				return
			}
			return
		}
	})

	r.Get("/manga", func(writer http.ResponseWriter, request *http.Request) {
		manga, err := s.library.Manga().All(request.Context())
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, err := fmt.Fprintf(writer, "DB err: %v", err)
			if err != nil {
				return
			}
			return
		}
		render.JSON(writer, request, manga)
	})

	r.Get("/manga/{id}", func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			_, err := fmt.Fprintf(writer, "Unknown err: %v", err)
			if err != nil {
				return
			}
			return
		}

		manga, err := s.library.Manga().ByID(request.Context(), id)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, err := fmt.Fprintf(writer, "DB err: %v", err)
			if err != nil {
				return
			}
			return
		}
		render.JSON(writer, request, manga)
	})

	r.Post("/manga", func(writer http.ResponseWriter, request *http.Request) {
		manga := new(models.Manga)
		if err := json.NewDecoder(request.Body).Decode(manga); err != nil {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			_, err := fmt.Fprintf(writer, "Unknown error: %v", err)
			if err != nil {
				return
			}
			return
		}
		if err := s.library.Manga().Create(request.Context(), manga); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, err := fmt.Fprintf(writer, "DB err: %v", err)
			if err != nil {
				return
			}
			return
		}
		writer.WriteHeader(http.StatusCreated)
	})

	r.Put("/manga", func(writer http.ResponseWriter, request *http.Request) {
		manga := new(models.Manga)
		if err := json.NewDecoder(request.Body).Decode(manga); err != nil {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			_, err := fmt.Fprintf(writer, "Unknown err: %v", err)
			if err != nil {
				return
			}
			return
		}

		err := validation.ValidateStruct(
			manga,
			validation.Field(&manga.ID, validation.Required),
			validation.Field(&manga.GenreID, validation.Required),
			validation.Field(&manga.Title, validation.Required),
			validation.Field(&manga.TitleJapanese, validation.Required),
			validation.Field(&manga.Volumes, validation.Required),
			validation.Field(&manga.Chapters, validation.Required),
			validation.Field(&manga.Score, validation.Required),
			validation.Field(&manga.Status, validation.Required),
			validation.Field(&manga.Synopsis, validation.Required),
		)

		if err != nil {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			_, err := fmt.Fprintf(writer, "Unknown err: %v", err)
			if err != nil {
				return
			}
			return
		}

		if err := s.library.Manga().Update(request.Context(), manga); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, err := fmt.Fprintf(writer, "DB err: %v", err)
			if err != nil {
				return
			}
			return
		}
	})

	r.Delete("/manga/{id}", func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			_, err := fmt.Fprintf(writer, "Unknown err: %v", err)
			if err != nil {
				return
			}
			return
		}
		if err := s.library.Manga().Delete(request.Context(), id); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, err := fmt.Fprintf(writer, "DB err: %v", err)
			if err != nil {
				return
			}
			return
		}
	})

	r.Get("/genres", func(writer http.ResponseWriter, request *http.Request) {
		genres, err := s.library.Genres().All(request.Context())
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, err := fmt.Fprintf(writer, "DB err: %v", err)
			if err != nil {
				return
			}
			return
		}
		render.JSON(writer, request, genres)
	})

	r.Get("/genres/{id}", func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			_, err := fmt.Fprintf(writer, "Unknown err: %v", err)
			if err != nil {
				return
			}
			return
		}

		genre, err := s.library.Genres().ByID(request.Context(), id)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, err := fmt.Fprintf(writer, "DB err: %v", err)
			if err != nil {
				return
			}
			return
		}
		render.JSON(writer, request, genre)
	})

	r.Post("/genres", func(writer http.ResponseWriter, request *http.Request) {
		genre := new(models.Genre)
		if err := json.NewDecoder(request.Body).Decode(genre); err != nil {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			_, err := fmt.Fprintf(writer, "Unknown error: %v", err)
			if err != nil {
				return
			}
			return
		}
		if err := s.library.Genres().Create(request.Context(), genre); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, err := fmt.Fprintf(writer, "DB err: %v", err)
			if err != nil {
				return
			}
			return
		}
		writer.WriteHeader(http.StatusCreated)
	})

	r.Put("/genres", func(writer http.ResponseWriter, request *http.Request) {
		genre := new(models.Genre)
		if err := json.NewDecoder(request.Body).Decode(genre); err != nil {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			_, err := fmt.Fprintf(writer, "Unknown err: %v", err)
			if err != nil {
				return
			}
			return
		}

		err := validation.ValidateStruct(
			genre,
			validation.Field(&genre.ID, validation.Required),
			validation.Field(&genre.Name, validation.Required),
		)

		if err != nil {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			_, err := fmt.Fprintf(writer, "Unknown err: %v", err)
			if err != nil {
				return
			}
			return
		}

		if err := s.library.Genres().Update(request.Context(), genre); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, err := fmt.Fprintf(writer, "DB err: %v", err)
			if err != nil {
				return
			}
			return
		}
	})

	r.Delete("/genres/{id}", func(writer http.ResponseWriter, request *http.Request) {
		idStr := chi.URLParam(request, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			_, err := fmt.Fprintf(writer, "Unknown err: %v", err)
			if err != nil {
				return
			}
			return
		}
		if err := s.library.Genres().Delete(request.Context(), id); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, err := fmt.Fprintf(writer, "DB err: %v", err)
			if err != nil {
				return
			}
			return
		}
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
