package auth

import (
	"ecom/configs"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWT(secret []byte, userId int) (string, error) {
	expirationTime := time.Second * time.Duration(configs.Envs.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": strconv.Itoa(userId),
		"exp":    expirationTime,
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
