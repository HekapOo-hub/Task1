package jwtToken

import (
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/golang-jwt/jwt"
)

var key = []byte("superSecretKey")

type CustomClaims struct {
	Login string
	Role  string
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
	tokenType, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return "", "", fmt.Errorf("token parsing error %w", err)
	}
	// Validate the token and return the custom claims
	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {

		return claims.Login, claims.Role, nil
	} else {
		return "", "", fmt.Errorf("token valid error %w", err)
	}
}
