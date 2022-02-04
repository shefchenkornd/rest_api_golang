package rest_api

import "github.com/sirupsen/logrus"

// App Base App instance description
type App struct {
	// UNEXPORTED FIELD!
	config *Config
	logger *logrus.Logger
}

// NewApp  App constructor: build base App instance
func NewApp(config *Config) *App {
	return &App{
		config: config,
		logger: logrus.New(),
	}
}

// Start http server/configure Loggers, router, db and etc...
func (app *App) Start() error {
	if err := app.configureLoggerField(); err != nil {
		return err
	}

	app.logger.Infoln("Starting api server at port:", app.config.BindAddr)
	return nil
}
