package models

type Manga struct {
	ID            int     `json:"id" db:"id"`
	GenreID       int     `json:"genre_id" db:"genre_id"`
	Title         string  `json:"title" db:"title"`
	TitleJapanese string  `json:"title_japanese" db:"titlejapanese"`
	Volumes       int     `json:"volumes" db:"volumes"`
	Chapters      int     `json:"chapters" db:"chapters"`
	Score         float64 `json:"score" db:"score"`
	Status        string  `json:"status" db:"status"`
	Synopsis      string  `json:"synopsis" db:"synopsis"`
}
