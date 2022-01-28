package jwtToken

import (
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/golang-jwt/jwt"
)

var key = []byte("superSecretKey")

type CustomClaims struct {
	login string
	role  string
	jwt.StandardClaims
}

func EncodeToken(user *model.User) (string, error) {
	claims := CustomClaims{
		user.Login,
		user.Role,
		jwt.StandardClaims{},
	}
	// Sign token and return
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
func DecodeToken(token string) (string, string, error) {
	// Parse the token
	tokenType, err := jwt.ParseWithClaims(string(key), &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	// Validate the token and return the custom claims
	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		return claims.login, claims.role, nil
	} else {
		return "", "", fmt.Errorf("token parsing error %w", err)
	}
}
