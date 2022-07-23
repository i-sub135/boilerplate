package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

// CreateJwt -- create jwt
func CreateJwt(param map[string]interface{}) (string, error) {
	keys := []byte(os.Getenv("SECRET_KEY"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	// Set some claims

	for i, m := range param {
		claims[i] = m
	}
	claims["exp"] = time.Now().Add(time.Hour * 24).Local()
	token.Claims = claims

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(keys)
	return tokenString, err
}

// ValidJwt -- validate
func ValidJwt(MyToken string) (jwt.MapClaims, error) {
	keys := os.Getenv("SECRET_KEY")
	token, err := jwt.Parse(MyToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(keys), nil
	})
	if token == nil {
		return nil, errors.New("not format token")

	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err

}
