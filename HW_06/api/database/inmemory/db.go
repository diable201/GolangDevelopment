package inmemory

import (
	"anime/api/database"
	"anime/api/models"
	"context"
	"fmt"
	"sync"
)

type Database struct {
	data map[int]*models.Anime
	mu   *sync.RWMutex
}

func NewDB() database.Library {
	return &Database{
		data: make(map[int]*models.Anime),
		mu:   new(sync.RWMutex),
	}
}

func (d *Database) Create(_ context.Context, anime *models.Anime) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[anime.ID] = anime
	return nil
}

func (d *Database) All(context.Context) ([]*models.Anime, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	animeList := make([]*models.Anime, 0, len(d.data))
	for _, anime := range d.data {
		animeList = append(animeList, anime)
	}
	return animeList, nil
}

func (d *Database) Update(_ context.Context, anime *models.Anime) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[anime.ID] = anime
	return nil
}

func (d *Database) Delete(_ context.Context, id int) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.data, id)
	return nil
}

func (d *Database) ByID(_ context.Context, id int) (*models.Anime, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	anime, ok := d.data[id]
	if !ok {
		return nil, fmt.Errorf("NO ANIME WITH ID %d", id)
	}
	return anime, nil
}
