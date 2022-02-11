package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/shefchenkornd/rest_api/internal/models"
	"net/http"
	"strconv"
)

func (app *App) GetAllArticles(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)

	// Логируем запрос
	app.logger.Infoln("Get all articles")

	articles, err := app.storage.Article().SelectAll()
	if err != nil {
		// Логируем запрос
		app.logger.Errorln("Error while Articles.SelectAll:", err)

		msg := Message{
			StatusCode: http.StatusInternalServerError,
			Message:    "We have some troubles to accessing DB. Try again later",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(articles)
}

func (app *App) GetArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)
	app.logger.Infoln("Get article by id")

	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.logger.Errorln("Error invalid id param")

		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid id param",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	article, found, err := app.storage.Article().FindById(id)
	if err != nil {
		app.logger.Errorln("Error in storage", err)

		msg := Message{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal error. Try later",
			IsError:    false,
		}
		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if !found {
		app.logger.Infoln("Article not found by id:", id)

		msg := Message{
			StatusCode: http.StatusNotFound,
			Message:    "Not found article by id",
			IsError:    false,
		}
		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	json.NewEncoder(writer).Encode(article)
}

func (app *App) CreateArticle(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)

	// Логируем запрос
	app.logger.Infoln("Create article")

	article := new(models.Article)
	if err := json.NewDecoder(req.Body).Decode(article); err != nil {
		app.logger.Errorln("Invalid json data:", err)

		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid json",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	articleCreated, err := app.storage.Article().Create(article)
	if err != nil {
		app.logger.Errorln("Can't create article", err)

		msg := Message{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Can't create article",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	json.NewEncoder(writer).Encode(articleCreated)
}

func (app *App) UpdateArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)
}

func (app *App) DeleteArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)
	app.logger.Infoln("Delete article by id", req.URL)

	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		app.logger.Errorln("Error invalid id param")

		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid id param",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, found, err := app.storage.Article().DeleteById(id)
	if err != nil {
		app.logger.Errorln("Error in storage", err)

		msg := Message{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal error. Try later",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if !found {
		app.logger.Infoln("article not found")

		msg := Message{
			StatusCode: http.StatusNotFound,
			Message:    "Article not found",
			IsError:    false,
		}
		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	msg := Message{
		StatusCode: http.StatusOK,
		Message:    "Article deleted successfully",
		IsError:    false,
	}
	writer.WriteHeader(msg.StatusCode)
	json.NewEncoder(writer).Encode(msg)
}
