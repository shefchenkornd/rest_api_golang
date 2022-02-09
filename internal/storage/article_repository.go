package storage

import (
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

// Create article
func (ar *ArticleRepository) Create(model *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, author, content) VALUES ($1, $2, $3) RETURNING id", tableUser)
	if err := ar.storage.db.QueryRow(query, model.Title, model.Author, model.Content).Scan(&model.Id); err != nil {
		return nil, err
	}

	return model, nil
}

// DeleteById ...
func (ar *ArticleRepository) DeleteById(id int) (*models.Article, error) {
	article, found, err := ar.FindById(id)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("article not found")
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1;", tableUser)
	_, err = ar.storage.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	return article, nil
}

// FindById ...
func (ar *ArticleRepository) FindById(id int) (*models.Article, bool, error) {
	var article *models.Article
	founded := false

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", tableUser)
	if err := ar.storage.db.QueryRow(query, id).Scan(&article); err != nil {
		return nil, false, err
	}
	founded = true

	return article, founded, nil
}

// SelectAll all articles
func (ar *ArticleRepository) SelectAll() ([]*models.Article, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableUser)
	rows, err := ar.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := make([]*models.Article, 100)
	for rows.Next() {
		article := &models.Article{}
		if err := rows.Scan(article.Id, article.Title, article.Author, article.Content); err != nil {
			log.Println(err)
			continue
		}
		articles = append(articles, article)
	}

	return articles, nil
}
