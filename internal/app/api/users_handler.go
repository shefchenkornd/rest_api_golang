package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func initHeader(writer http.ResponseWriter)  {
	writer.Header().Set("Content-Type", "application/json")
}

func (app *App) UserRegister(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)

	mux.Vars(req)
}
