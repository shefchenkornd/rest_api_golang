package api

import (
	"github.com/gorilla/mux"
	"github.com/shefchenkornd/rest_api/internal/storage"
	"github.com/sirupsen/logrus"
	"net/http"
)

// App Base App instance description
type App struct {
	// UNEXPORTED FIELD!
	config *Config
	logger *logrus.Logger
	router *mux.Router
	storage *storage.Storage
}

// NewApp  App constructor: build base App instance
func NewApp(config *Config) *App {
	return &App{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		storage: storage.New(config),
	}
}

// Start http server/configure Loggers, router, db and etc...
func (app *App) Start() error {
	if err := app.configureLoggerField(); err != nil {
		return err
	}

	// подтверждение того, что logger сконфигурирован!
	app.logger.Infoln("Starting api server at port", app.config.BindAddr)

	// собираем маршрутизатор
	app.configureRouterField()

	// подключаемся к нашей БД
	if err := app.configureStorageField(); err != nil {
		return err
	}

	// На этапе валидного завершения стартуем http server
	return http.ListenAndServe(app.config.BindAddr, app.router)
}
