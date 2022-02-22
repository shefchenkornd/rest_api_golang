package middleware

import (
	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

var (
	SecretKey      []byte      = []byte("secret-token-3000")
	emptyValidFunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	}
)

var JwtMiddleware = jwtMiddleware.New(
	jwtMiddleware.Options{
		ValidationKeyGetter: emptyValidFunc,
		SigningMethod:       jwt.SigningMethodHS256,
	},
)
