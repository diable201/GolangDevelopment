package inmemory

import (
	"anime/api/models"
	"context"
	"fmt"
	"sync"
)

type MangaRepository struct {
	data map[int]*models.Manga
	mu   *sync.RWMutex
}

func (d *MangaRepository) Create(_ context.Context, manga *models.Manga) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[manga.ID] = manga
	return nil
}

func (d *MangaRepository) All(context.Context) ([]*models.Manga, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	mangaList := make([]*models.Manga, 0, len(d.data))
	for _, manga := range d.data {
		mangaList = append(mangaList, manga)
	}
	return mangaList, nil
}

func (d *MangaRepository) Update(_ context.Context, manga *models.Manga) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[manga.ID] = manga
	return nil
}

func (d *MangaRepository) Delete(_ context.Context, id int) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.data, id)
	return nil
}

func (d *MangaRepository) ByID(_ context.Context, id int) (*models.Manga, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	manga, ok := d.data[id]
	if !ok {
		return nil, fmt.Errorf("NO MANGA WITH ID %d", id)
	}
	return manga, nil
}
