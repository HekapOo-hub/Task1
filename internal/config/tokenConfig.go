package config

import (
	"github.com/HekapOo-hub/Task1/internal/service"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

const (
	AccessKey  = "superSecretKey"
	RefreshKey = "wgnbwglwrgnl"
	AccessTTL  = time.Minute * 15
	RefreshTTL = time.Hour * 24 * 7
)

// GetAccessTokenConfig returns access jwt config for middleware authentication
func GetAccessTokenConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &service.TokenClaims{},
		SigningKey: []byte(AccessKey),
	}
}

// GetRefreshTokenConfig returns refresh jwt config for middleware authentication
func GetRefreshTokenConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &service.TokenClaims{},
		SigningKey: []byte(RefreshKey),
	}
}
