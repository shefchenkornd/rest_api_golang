package rest_api

// App Base App instance description
type App struct {
	// UNEXPORTED FIELD!
	config *Config
}

// NewApp  App constructor: build base App instance
func NewApp(config *Config) *App {
	return &App{config: config}
}

// Start http server/configure Loggers, router, db and etc...
func (app *App) Start() error {
	return nil
}

