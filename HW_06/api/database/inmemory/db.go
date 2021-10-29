package inmemory

import (
	"anime/api/database"
	"anime/api/models"
	"sync"
)

type Database struct {
	animeRepository database.AnimeRepository
	mangaRepository database.MangaRepository
	mu              *sync.RWMutex
}

func NewDB() database.Library {
	return &Database{
		mu: new(sync.RWMutex),
	}
}

func (db *Database) Anime() database.AnimeRepository {
	if db.animeRepository == nil {
		db.animeRepository = &AnimeRepository{
			data: make(map[int]*models.Anime),
			mu:   new(sync.RWMutex),
		}
	}
	return db.animeRepository
}

func (db *Database) Manga() database.MangaRepository {
	if db.mangaRepository == nil {
		db.mangaRepository = &MangaRepository{
			data: make(map[int]*models.Manga),
			mu:   new(sync.RWMutex),
		}
	}
	return db.mangaRepository
}
