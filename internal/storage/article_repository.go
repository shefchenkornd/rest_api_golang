package storage

// ArticleRepository Хотим, чтобы наше приложение общалось с моделью Article через репозиторий ArticleRepository
type ArticleRepository struct {
	storage *Storage
}
