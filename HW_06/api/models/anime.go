package models

type Anime struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	TitleJapanese string  `json:"title_japanese"`
	Source        string  `json:"source"`
	Episodes      int     `json:"episodes"`
	Kind          string  `json:"kind"`
	Score         float64 `json:"score"`
	Status        string  `json:"status"`
	Synopsis      string  `json:"synopsis"`
}
