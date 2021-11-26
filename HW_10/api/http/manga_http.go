package http

import (
	"anime-kafka/api/database"
	"anime-kafka/api/message_broker"
	"anime-kafka/api/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	lru "github.com/hashicorp/golang-lru"
	"net/http"
	"strconv"
)

type MangaResource struct {
	library database.Library
	broker  message_broker.MessageBroker
	cache   *lru.TwoQueueCache
}

func NewMangaResource(library database.Library, broker message_broker.MessageBroker, cache *lru.TwoQueueCache) *MangaResource {
	return &MangaResource{
		library: library,
		broker:  broker,
		cache:   cache,
	}
}

func (mr *MangaResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", mr.CreateManga)
	r.Get("/", mr.AllManga)
	r.Get("/{id}", mr.MangaByID)
	r.Put("/", mr.UpdateManga)
	r.Delete("/{id}", mr.DeleteManga)

	return r
}

func (mr *MangaResource) CreateManga(writer http.ResponseWriter, request *http.Request) {
	manga := new(models.Manga)
	if err := json.NewDecoder(request.Body).Decode(manga); err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		_, err := fmt.Fprintf(writer, "Unknown error: %v", err)
		if err != nil {
			return
		}
		return
	}
	if err := mr.library.Manga().Create(request.Context(), manga); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
	err := mr.broker.Cache().Purge()
	if err != nil {
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (mr *MangaResource) AllManga(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()
	filter := &models.MangaFilter{}
	searchQuery := queryValues.Get("query")
	if searchQuery != "" {
		mangaFromCache, ok := mr.cache.Get(searchQuery)
		if ok {
			render.JSON(writer, request, mangaFromCache)
			return
		}
		filter.Query = &searchQuery
	}
	manga, err := mr.library.Manga().All(request.Context(), filter)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
	if searchQuery != "" {
		mr.cache.Add(searchQuery, manga)
	}
	render.JSON(writer, request, manga)
}

func (mr *MangaResource) MangaByID(writer http.ResponseWriter, request *http.Request) {
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

	mangaFromCache, ok := mr.cache.Get(id)
	if ok {
		render.JSON(writer, request, mangaFromCache)
		return
	}

	manga, err := mr.library.Manga().ByID(request.Context(), id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
	mr.cache.Add(id, manga)
	render.JSON(writer, request, manga)
}

func (mr *MangaResource) UpdateManga(writer http.ResponseWriter, request *http.Request) {
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

	if err := mr.library.Manga().Update(request.Context(), manga); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
	err = mr.broker.Cache().Remove(manga.ID)
	if err != nil {
		return
	}
}

func (mr *MangaResource) DeleteManga(writer http.ResponseWriter, request *http.Request) {
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
	if err := mr.library.Manga().Delete(request.Context(), id); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
	err = mr.broker.Cache().Remove(id)
	if err != nil {
		return
	}
}
