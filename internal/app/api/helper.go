package api

import (
	"github.com/sirupsen/logrus"
	"net/http"
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
	app.router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello!!!"))
	})
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
