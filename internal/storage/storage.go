package storage

import (
	"database/sql"
	_ "github.com/lib/pq" // Для того чтобы отработала функция init() данного пакета
)

type Storage struct {
	databaseUrl       string
	db                *sql.DB
	articleRepository *ArticleRepository
	userRepository    *UserRepository
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

func (s *Storage) User() *UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			storage: s,
		}
	}

	return s.userRepository
}

func (s *Storage) Article() *ArticleRepository {
	if s.articleRepository == nil {
		s.articleRepository = &ArticleRepository{
			storage: s,
		}
	}

	return s.articleRepository
}
