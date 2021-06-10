package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fransempe/backendGo/models"
)

//GenerateJWT is a function that generate a new user token
func GenerateJWT(user models.User) (string, error) {

	sigKey := []byte("backendGo")

	payload := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(sigKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
