package rest_api

import "github.com/sirupsen/logrus"

// configureLoggerField пытаемся конфигурировать наш App инстанс (а конкретнее поле - Logger)
func (app *App) configureLoggerField() error {
	logLevel, err := logrus.ParseLevel(app.config.LoggerLevel)
	if err != nil {
		return err
	}

	app.logger.SetLevel(logLevel)
	return nil
}
