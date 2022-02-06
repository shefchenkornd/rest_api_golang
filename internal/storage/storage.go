package storage

import (
	"database/sql"
	_ "github.com/lib/pq" // Для того чтобы отработала функция init() данного пакета
)

type Storage struct {
	databaseUrl string
	db     *sql.DB
}

func New(databaseUrl string) *Storage {
	return &Storage{
		databaseUrl: databaseUrl,
	}
}

// Open connection
func (s *Storage) Open() error {
	db, err := sql.Open("postgres", s.databaseUrl)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close connection
func (s *Storage) Close() {
	s.db.Close()
}
