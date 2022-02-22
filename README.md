## Задача: создать REST API с CRUD, который будет позволять получать информацию про пользователей и статьи из PostgreSQL

### Шаг 1. Желаемый функционал.

Хотим собрать веб-сервер, который будет взаимодействовать с окружающим миром через API, поддерживающий REST

### Шаг 1.1 Какие виды моделей есть в нашем приложении:

* user - пользователи
* article - статьи

### Шаг 1.1 Виды запросов поддерживаются API

Будет существовать следующие запроса для статьи:

* GET `http://localhost:8080/api/v1/articles` - возвращает json со всеми статьями.
* GET `http://localhost:8080/api/v1/article/{id}` - возвращает информацию о статье с `id`, если такой статьи нет, то
  сообщаем об этом.
* POST `http://localhost:8080/api/v1/article/` - создаём новую статью
* PUT `http://localhost:8080/api/v1/article/{id}` - обновляем информацию об уже существующей статье с `id`, если такой
  статьи нет, то сообщаем об этом.
* DELETE `http://localhost:8080/api/v1/article/{id}` - удаляем статью с `id` , если такой статьи нет, то сообщаем об
  этом.

Будет существовать следующие запроса для пользователей:
* POST `http://localhost:8080/api/v1/user/register` - регистрируем нового пользователя.

### Шаг 2. Реализация

### Шаг 2.1 Маршрутизатор и исполнители

***Маршрутизатор (router)*** - это экземпляр, который имеет внутренний функционал, заключающийся в следующем:

* принимает на вход адрес запроса (по сути это строка `http://localhost:8080/api/v1/articles`) и вызывает исполнителя, который
  будет ассоциирован с этим запросом.

***Исполнитель (handler)*** - это функция/метод, который вызывается маршрутизатором.

Для того чтобы удобно работать с маршрутизатором и не писать с нуля, будем использовать готовую
библиотеку `github.com/gorilla/mux`:
`go get -u github.com/gorilla/mux`

Для работы с JWT токеном будем использовать следующие пакеты:
* go get -u github.com/auth0/go-jwt-middleware
* go get -u github.com/form3tech-oss/jwt-go

### Шаг 3. Добавить JWT аутентификацию
Завернем необходимые HTTP-хендлеры в JWT-декоратор. Для того, чтобы обозначить факт необходимости использования JWT токена
перед выполнением какого-либо запроса.
Заверните HTTP-хендлер в декоратор `JWTMiddleware.Handler(h http.Handler) http.Handler`
```
app.router.Handle(
    prefixApi+"/articles/{id}",
    middleware.JwtMiddleware.Handler(
        http.HandlerFunc(app.GetArticleById),
    ),
).Methods("GET")
```


В postman'e получите JWT токен и используйте его на вкладке "Headers", добавив таким образом:<br>
`Authorization: Bearer <your_token_from_auth>`