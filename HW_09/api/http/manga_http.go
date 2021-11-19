package http

import (
	"anime-redis/api/cache"
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

type MangaResource struct {
	library database.Library
	cache   cache.MangaCache
}

func NewMangaResource(library database.Library, cache cache.MangaCache) *MangaResource {
	return &MangaResource{
		library: library,
		cache:   cache,
	}
}

func (gr *MangaResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", gr.CreateManga)
	r.Get("/", gr.AllManga)
	r.Get("/{id}", gr.MangaByID)
	r.Put("/", gr.UpdateManga)
	r.Delete("/{id}", gr.DeleteManga)

	return r
}

func (gr *MangaResource) CreateManga(writer http.ResponseWriter, request *http.Request) {
	manga := new(models.Manga)
	if err := json.NewDecoder(request.Body).Decode(manga); err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		_, err := fmt.Fprintf(writer, "Unknown error: %v", err)
		if err != nil {
			return
		}
		return
	}
	if err := gr.library.Manga().Create(request.Context(), manga); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (gr *MangaResource) AllManga(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()
	filter := &models.MangaFilter{}

	if searchQuery := queryValues.Get("query"); searchQuery != "" {
		filter.Query = &searchQuery
	}
	manga, err := gr.library.Manga().All(request.Context(), filter)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
	render.JSON(writer, request, manga)
}

func (gr *MangaResource) MangaByID(writer http.ResponseWriter, request *http.Request) {
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

	mangaFromCache := gr.cache.Get(request.Context(), idStr)
	if mangaFromCache != nil {
		render.JSON(writer, request, mangaFromCache)
		return
	}

	manga, err := gr.library.Manga().ByID(request.Context(), id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
	gr.cache.Set(request.Context(), idStr, manga)
	render.JSON(writer, request, manga)
}

func (gr *MangaResource) UpdateManga(writer http.ResponseWriter, request *http.Request) {
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

	if err := gr.library.Manga().Update(request.Context(), manga); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
}

func (gr *MangaResource) DeleteManga(writer http.ResponseWriter, request *http.Request) {
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
	if err := gr.library.Manga().Delete(request.Context(), id); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
}
