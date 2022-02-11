package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/shefchenkornd/rest_api/internal/models"
	"log"
)

// ArticleRepository Хотим, чтобы наше приложение общалось с моделью Article через репозиторий ArticleRepository
type ArticleRepository struct {
	storage *Storage
}

var (
	tableArticle = "articles"
)

// Create new article
func (ar *ArticleRepository) Create(model *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, author, content) VALUES ($1, $2, $3) RETURNING id", tableArticle)
	if err := ar.storage.db.QueryRow(query, model.Title, model.Author, model.Content).Scan(&model.Id); err != nil {
		return nil, err
	}

	return model, nil
}

// FindById find article by id
func (ar *ArticleRepository) FindById(id int) (*models.Article, bool, error) {
	founded := false

	query := fmt.Sprintf("SELECT title, author, content FROM %s WHERE id=$1", tableArticle)
	var title string
	var author string
	var content string
	if err := ar.storage.db.QueryRow(query, id).Scan(&title, &author, &content); err != nil {
		// если нет результатов из БД, то это не ошибка!
		if errors.Is(err, sql.ErrNoRows) {
			return nil, founded, nil
		}

		return nil, founded, err
	}
	founded = true

	article := &models.Article{
		Id:      id,
		Title:   title,
		Author:  author,
		Content: content,
	}

	return article, founded, nil
}

// SelectAll select all articles
func (ar *ArticleRepository) SelectAll() ([]*models.Article, error) {
	query := fmt.Sprintf("SELECT id, title, author, content FROM %s", tableArticle)
	rows, err := ar.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Подготовим куда, будем читать данные
	articles := make([]*models.Article, 0)
	for rows.Next() {
		var id int
		var title string
		var author string
		var content string
		if err := rows.Scan(&id, &title, &author, &content); err != nil {
			log.Println(err)
			continue
		}
		article := &models.Article{
			Id:      id,
			Title:   title,
			Author:  author,
			Content: content,
		}
		articles = append(articles, article)
	}

	return articles, nil
}

// DeleteById delete article by id
func (ar *ArticleRepository) DeleteById(id int) (*models.Article, bool, error) {
	article, found, err := ar.FindById(id)
	if err != nil {
		return nil, false, err
	}

	if !found {
		return nil, false, nil
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1;", tableArticle)
	_, err = ar.storage.db.Query(query, id)
	if err != nil {
		return nil, false, err
	}

	return article, found,  nil
}