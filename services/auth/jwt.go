package auth

import (
	"ecom/configs"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

func GetUserID(r *http.Request) int {
	// Get the Authorization header
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		fmt.Printf("No Authorization header provided")
		return 0
	}

	// The token might be in the form "Bearer <token>", so we need to split it
	tokenParts := strings.Split(tokenString, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		fmt.Printf("Invalid Authorization format")
		return 0
	}

	// Parse the JWT token
	token, err := jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// Return your secret key here (for example, HMAC secret)
		return []byte("your-secret-key"), nil
	})

	if err != nil || !token.Valid {
		fmt.Printf("Invalid token: %v\n", err)
		return 0
	}

	// Extract the user ID from the token claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if userID, ok := claims["userId"].(float64); ok {
			return int(userID)
		}
	}

	return 0
}
