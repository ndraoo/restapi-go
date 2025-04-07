package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ndraoo/restapi-go/config"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":     strconv.Itoa(userID),
		"expired_At": time.Now().Add(expiration).Unix(), // Token expiration time
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
