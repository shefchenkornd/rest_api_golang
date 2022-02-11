package api

import (
	"encoding/json"
	"github.com/shefchenkornd/rest_api/internal/models"
	"net/http"
)

func initHeader(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func (app *App) UserRegister(writer http.ResponseWriter, req *http.Request) {
	initHeader(writer)
	app.logger.Infoln("Post user register", req.URL)

	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
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

	_, found, err := app.storage.User().FindByLogin(user.Login)
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

	if found {
		app.logger.Infoln("User with that login already exists.")

		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "User with that login already exists.",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	userCreated, err := app.storage.User().Create(&user)
	if err != nil {
		app.logger.Errorln("Can't create user:", err)

		msg := Message{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Can't create user",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(userCreated)
}
