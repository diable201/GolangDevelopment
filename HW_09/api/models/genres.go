package models

type (
	Genre struct {
		ID   int    `json:"id" db:"id"`
		Name string `json:"name" db:"name"`
	}

	GenreFilter struct {
		Query *string `json:"query"`
	}
)
