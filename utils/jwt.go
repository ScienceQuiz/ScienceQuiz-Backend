package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type CustomClaims map[string]interface{}

func (cc CustomClaims) Valid() error {
	return jwt.MapClaims(cc).Valid()
}

var JWTSecret = []byte("SECRET_KEY")

func GenerateJWT(id uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	t, _ := token.SignedString(JWTSecret)
	return t
}

func DecodeJWT(token string) uint {
	t, _ := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})
	return uint((*t.Claims.(*CustomClaims))["id"].(float64))
}