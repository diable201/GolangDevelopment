package models

type Anime struct {
	ID            int     `json:"id" db:"id"`
	GenreID       int     `json:"genre_id" db:"genre_id"`
	Title         string  `json:"title" db:"title"`
	TitleJapanese string  `json:"title_japanese" db:"titlejapanese"`
	Source        string  `json:"source" db:"source"`
	Episodes      int     `json:"episodes" db:"episodes"`
	Kind          string  `json:"kind" db:"kind"`
	Score         float64 `json:"score" db:"score"`
	Status        string  `json:"status" db:"status"`
	Synopsis      string  `json:"synopsis" db:"synopsis"`
}
