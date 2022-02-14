package config

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

const (
	// AccessKey is used for signing access token
	AccessKey = "superSecretKey"
	// RefreshKey is used for signing refresh token
	RefreshKey = "wgnbwglwrgnl"
	// AccessTTL is access token's time to live
	AccessTTL = time.Minute * 15
	// RefreshTTL is refresh token's time to live
	RefreshTTL = time.Hour * 24 * 7
)

// TokenClaims describes custom token claim
type TokenClaims struct {
	Login string
	Role  string
	ID    string
	jwt.StandardClaims
}

// GetAccessTokenConfig returns access jwt config for middleware authentication
func GetAccessTokenConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &TokenClaims{},
		SigningKey: []byte(AccessKey),
	}
}

// GetRefreshTokenConfig returns refresh jwt config for middleware authentication
func GetRefreshTokenConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &TokenClaims{},
		SigningKey: []byte(RefreshKey),
	}
}
