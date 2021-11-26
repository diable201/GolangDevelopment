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

type AnimeResource struct {
	library database.Library
	broker  message_broker.MessageBroker
	cache   *lru.TwoQueueCache
}

func NewAnimeResource(library database.Library, broker message_broker.MessageBroker, cache *lru.TwoQueueCache) *AnimeResource {
	return &AnimeResource{
		library: library,
		broker:  broker,
		cache:   cache,
	}
}

func (ar *AnimeResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", ar.CreateAnime)
	r.Get("/", ar.AllAnime)
	r.Get("/{id}", ar.AnimeByID)
	r.Put("/", ar.UpdateAnime)
	r.Delete("/{id}", ar.DeleteAnime)

	return r
}

func (ar *AnimeResource) CreateAnime(writer http.ResponseWriter, request *http.Request) {
	anime := new(models.Anime)
	if err := json.NewDecoder(request.Body).Decode(anime); err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		_, err := fmt.Fprintf(writer, "Unknown error: %v", err)
		if err != nil {
			return
		}
		return
	}
	if err := ar.library.Anime().Create(request.Context(), anime); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
	err := ar.broker.Cache().Purge()
	if err != nil {
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (ar *AnimeResource) AllAnime(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()
	filter := &models.AnimeFilter{}
	searchQuery := queryValues.Get("query")
	if searchQuery != "" {
		animeFromCache, ok := ar.cache.Get(searchQuery)
		if ok {
			render.JSON(writer, request, animeFromCache)
			return
		}
		filter.Query = &searchQuery
	}
	anime, err := ar.library.Anime().All(request.Context(), filter)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
	if searchQuery != "" {
		ar.cache.Add(searchQuery, anime)
	}
	render.JSON(writer, request, anime)
}

func (ar *AnimeResource) AnimeByID(writer http.ResponseWriter, request *http.Request) {
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

	animeFromCache, ok := ar.cache.Get(id)
	if ok {
		render.JSON(writer, request, animeFromCache)
		return
	}

	anime, err := ar.library.Anime().ByID(request.Context(), id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
	ar.cache.Add(id, anime)
	render.JSON(writer, request, anime)
}

func (ar *AnimeResource) UpdateAnime(writer http.ResponseWriter, request *http.Request) {
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

	if err := ar.library.Anime().Update(request.Context(), anime); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
	err = ar.broker.Cache().Remove(anime.ID)
	if err != nil {
		return
	}
}

func (ar *AnimeResource) DeleteAnime(writer http.ResponseWriter, request *http.Request) {
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

	if err := ar.library.Anime().Delete(request.Context(), id); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(writer, "DB err: %v", err)
		if err != nil {
			return
		}
		return
	}
	err = ar.broker.Cache().Remove(id)
	if err != nil {
		return
	}
}
