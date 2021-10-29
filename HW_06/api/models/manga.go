package models

type Manga struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	TitleJapanese string  `json:"title_japanese"`
	Volumes       int     `json:"volumes"`
	Chapters      int     `json:"chapters"`
	Score         float64 `json:"score"`
	Status        string  `json:"status"`
	Synopsis      string  `json:"synopsis"`
}
