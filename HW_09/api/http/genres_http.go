package http

import (
	"anime-redis/api/database"
	"anime-redis/api/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
	"strconv"
)

type GenresResource struct {
	library database.Library
}

func NewGenresResource(library database.Library) *GenresResource {
	return &GenresResource{
		library: library,
	}
}

func (gr *GenresResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", gr.CreateGenre)
	r.Get("/", gr.AllGenres)
	r.Get("/{id}", gr.GenreByID)
	r.Put("/", gr.UpdateGenre)
	r.Delete("/{id}", gr.DeleteGenre)

	return r
}

func (gr *GenresResource) CreateGenre(writer http.ResponseWriter, request *http.Request) {
	genre := new(models.Genre)
	if err := json.NewDecoder(request.Body).Decode(genre); err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		_, err := fmt.Fprintf(writer, "Unknown error: %v", err)
		if err != nil {
			return
		}
		return
	}
	if err := gr.library.Genres().Create(request.Context(), genre); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (gr *GenresResource) AllGenres(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()
	filter := &models.GenreFilter{}

	if searchQuery := queryValues.Get("query"); searchQuery != "" {
		filter.Query = &searchQuery
	}
	genres, err := gr.library.Genres().All(request.Context(), filter)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
	render.JSON(writer, request, genres)
}

func (gr *GenresResource) GenreByID(writer http.ResponseWriter, request *http.Request) {
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

	genre, err := gr.library.Genres().ByID(request.Context(), id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
	render.JSON(writer, request, genre)
}

func (gr *GenresResource) UpdateGenre(writer http.ResponseWriter, request *http.Request) {
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

	if err := gr.library.Genres().Update(request.Context(), genre); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
}

func (gr *GenresResource) DeleteGenre(writer http.ResponseWriter, request *http.Request) {
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
	if err := gr.library.Genres().Delete(request.Context(), id); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
}
