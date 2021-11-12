package postgres

import (
	"anime-database/api/database"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	connection *sqlx.DB
	genres     database.GenresRepository
	anime      database.AnimeRepository
	manga      database.MangaRepository
}

func NewDB() database.Library {
	return &DB{}
}

func (db *DB) Connect(url string) error {
	connection, err := sqlx.Connect("pgx", url)
	if err != nil {
		return err
	}

	if err := connection.Ping(); err != nil {
		return err
	}

	db.connection = connection
	return nil
}

func (db *DB) Close() error {
	return db.connection.Close()
}
