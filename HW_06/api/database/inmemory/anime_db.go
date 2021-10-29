package inmemory

import (
	"anime/api/models"
	"context"
	"fmt"
	"sync"
)

type AnimeRepository struct {
	data map[int]*models.Anime
	mu   *sync.RWMutex
}

func (d *AnimeRepository) Create(_ context.Context, anime *models.Anime) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[anime.ID] = anime
	return nil
}

func (d *AnimeRepository) All(context.Context) ([]*models.Anime, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	animeList := make([]*models.Anime, 0, len(d.data))
	for _, anime := range d.data {
		animeList = append(animeList, anime)
	}
	return animeList, nil
}

func (d *AnimeRepository) Update(_ context.Context, anime *models.Anime) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[anime.ID] = anime
	return nil
}

func (d *AnimeRepository) Delete(_ context.Context, id int) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.data, id)
	return nil
}

func (d *AnimeRepository) ByID(_ context.Context, id int) (*models.Anime, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	anime, ok := d.data[id]
	if !ok {
		return nil, fmt.Errorf("NO ANIME WITH ID %d", id)
	}
	return anime, nil
}
