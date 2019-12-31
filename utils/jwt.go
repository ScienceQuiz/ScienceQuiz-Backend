package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JWTSecret = []byte("SECRET_KEY")

func GenerateJWT(id uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	t, _ := token.SignedString(JWTSecret)
	return t
}