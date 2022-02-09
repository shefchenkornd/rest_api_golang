package api

import (
	"encoding/json"
	"net/http"
)

func (app *App) GetAllArticles(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)

	// Логируем запрос
	app.logger.Infoln("Get all articles")

	articles, err := app.storage.Article().SelectAll()
	if err != nil {
		// Логируем запрос
		app.logger.Infoln("Error while Articles.SelectAll:", err)

		msg := Message{
			StatusCode: http.StatusInternalServerError,
			Message:    "We have some troubles to accessing DB. Try again later",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(articles)
}

func (app *App) GetArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)
}

func (app *App) CreateArticle(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)
}

func (app *App) UpdateArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)
}

func (app *App) DeleteArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)
}
