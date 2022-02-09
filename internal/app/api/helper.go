package api

import (
	"github.com/sirupsen/logrus"
)

var (
	prefixApi = "/api/v1"
)

// configureLoggerField пытаемся конфигурировать наш App инстанс (а конкретнее поле - Logger)
func (app *App) configureLoggerField() error {
	logLevel, err := logrus.ParseLevel(app.config.LoggerLevel)
	if err != nil {
		return err
	}

	app.logger.SetLevel(logLevel)
	return nil
}

// configureRouterField Конфигурируем маршрутизатор
func (app *App) configureRouterField() {
	app.router.HandleFunc(prefixApi+"/articles", app.GetAllArticles).Methods("GET")
	app.router.HandleFunc(prefixApi+"/articles/{id}", app.GetArticleById).Methods("GET")
	app.router.HandleFunc(prefixApi+"/articles", app.CreateArticle).Methods("POST")
	app.router.HandleFunc(prefixApi+"/articles/{id}", app.UpdateArticleById).Methods("PUT")
	app.router.HandleFunc(prefixApi+"/articles/{id}", app.DeleteArticleById).Methods("DELETE")

	app.router.HandleFunc(prefixApi+"/user/register", app.UserRegister).Methods("POST")
}

// configureStorageField Конфигурируем хранилище
func (app *App) configureStorageField() error {
	// пытаемся соединиться с БД
	if err := app.storage.Open(); err != nil {
		return err
	}
	app.logger.Infoln("DB connection created!")

	return nil
}
