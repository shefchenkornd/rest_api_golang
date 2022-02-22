package api

import (
	"encoding/json"
	"github.com/form3tech-oss/jwt-go"
	"github.com/shefchenkornd/rest_api/internal/app/middleware"
	"github.com/shefchenkornd/rest_api/internal/models"
	"net/http"
	"time"
)

func initHeader(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func (app *App) UserRegister(w http.ResponseWriter, req *http.Request) {
	initHeader(w)
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
		w.WriteHeader(msg.StatusCode)
		json.NewEncoder(w).Encode(msg)
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
		w.WriteHeader(msg.StatusCode)
		json.NewEncoder(w).Encode(msg)
		return
	}

	if found {
		app.logger.Infoln("User with that login already exists.")

		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "User with that login already exists.",
			IsError:    true,
		}
		w.WriteHeader(msg.StatusCode)
		json.NewEncoder(w).Encode(msg)
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
		w.WriteHeader(msg.StatusCode)
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userCreated)
}

// UserAuth user authorization
func (app *App) UserAuth(w http.ResponseWriter, req *http.Request) {
	initHeader(w)
	app.logger.Infoln("post auth")

	// пробуем получить логин/пароль
	var userFromJson models.User
	err := json.NewDecoder(req.Body).Decode(&userFromJson)
	if err != nil {
		app.logger.Errorln("Invalid json data:", err)

		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid json",
			IsError:    true,
		}
		w.WriteHeader(msg.StatusCode)
		json.NewEncoder(w).Encode(msg)
		return
	}

	// пробуем найти эту пару логин/пароль в нашей БД
	userFromDB, found, err := app.storage.User().FindByLogin(userFromJson.Login)
	if err != nil {
		app.logger.Errorln("Error:", err)

		msg := Message{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Error",
			IsError:    true,
		}
		w.WriteHeader(msg.StatusCode)
		json.NewEncoder(w).Encode(msg)
		return
	}

	// если пользователь не найден
	if !found {
		app.logger.Infoln("Not found user")

		msg := Message{
			StatusCode: http.StatusNotFound,
			Message:    "Not found user",
			IsError:    true,
		}
		w.WriteHeader(msg.StatusCode)
		json.NewEncoder(w).Encode(msg)
		return
	}

	// если пароли не совпадают
	if userFromDB.Password != userFromJson.Password {
		app.logger.Infoln("Invalid credentials")

		msg := Message{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid credentials",
			IsError:    true,
		}
		w.WriteHeader(msg.StatusCode)
		json.NewEncoder(w).Encode(msg)
		return
	}

	// авторизовать найденного пользователя и отправить ему JWT токен
	token := jwt.New(jwt.SigningMethodHS256) // ВАЖНО: метод подписания такой же, как в и в файле internal/app/middleware/middleware.go
	claims := token.Claims.(jwt.MapClaims)   // дополнительные действия для шифрования
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // время истечения JWT токена через 2 часа

	tokenStr, err := token.SignedString(middleware.SecretKey)
	if err != nil {
		app.logger.Errorln("Can't generate token")

		msg:= Message{
			StatusCode: http.StatusInternalServerError,
			Message:    "We have some troubles",
			IsError:    true,
		}
		w.WriteHeader(msg.StatusCode)
		json.NewEncoder(w).Encode(msg)
	}

	// если токен был успешно сгенерирован, то отдаём его на фронт.
	msg := Message{
		StatusCode: http.StatusOK,
		Message:    tokenStr,
		IsError:    false,
	}
	w.WriteHeader(msg.StatusCode)
	json.NewEncoder(w).Encode(msg)
}
