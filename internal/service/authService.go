package service

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var accessKey = []byte("superSecretKey")
var refreshKey = []byte("wgnbwglwrgnl")

type TokenClaims struct {
	Login string
	Role  string
	ID    string
	jwt.StandardClaims
}

type AuthService struct {
	r repository.TokenRepository
}

func GetAccessTokenConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &TokenClaims{},
		SigningKey: accessKey,
	}
}

func GetRefreshTokenConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &TokenClaims{},
		SigningKey: refreshKey,
	}
}

func NewAuthService(r repository.TokenRepository) *AuthService {
	return &AuthService{r: r}
}
func (a *AuthService) encodeToken(user *model.User, expiresAt int64, style string) (*model.Token, error) {
	var key []byte
	if style == "access" {
		key = accessKey
	} else if style == "refresh" {
		key = refreshKey
	}
	claims := TokenClaims{
		user.Login,
		user.Role,
		uuid.NewV4().String(),
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	// Sign token and return
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	val, err := token.SignedString(key)
	if err != nil {
		return nil, fmt.Errorf("encode token error with signing %w", err)
	}
	return &model.Token{Value: val, ExpiresAt: expiresAt, Login: user.Login}, nil
}

func (a *AuthService) Create(token model.Token) error {

	err := a.r.Create(context.Background(), token)
	if err != nil {
		return fmt.Errorf("service layer create token error %w", err)
	}
	return nil
}
func (a *AuthService) Get(token string) (*model.Token, error) {
	tokenFromDB, err := a.r.Get(context.Background(), token)
	if err != nil {
		return nil, fmt.Errorf("service layer get token error %w", err)
	}
	return tokenFromDB, nil
}
func (a *AuthService) Delete(token string) error {
	err := a.r.Delete(context.Background(), token)
	if err != nil {
		return fmt.Errorf("authentication layer delete token error %w", err)
	}
	return nil
}
func (a *AuthService) Authenticate(user *model.User, password string) (string, string, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", fmt.Errorf("authentication comparing passwords error %w", err)
	}
	accessToken, err := a.encodeToken(user, time.Now().Add(time.Minute*15).Unix(), "access")
	if err != nil {
		return "", "", fmt.Errorf("service layer authentication encode access token error %w", err)
	}
	refreshToken, err := a.encodeToken(user, time.Now().Add(time.Hour*24*7).Unix(), "refresh")
	if err != nil {
		return "", "", fmt.Errorf("service layer authentication encode refresh token error %w", err)
	}

	err = a.r.Create(context.Background(), *refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("service layer  mongo create token error %w", err)
	}
	return accessToken.Value, refreshToken.Value, nil
}

func (a *AuthService) Refresh(claims *TokenClaims, token string) (string, string, error) {

	role := claims.Role
	login := claims.Login
	err := a.r.Delete(context.Background(), token)
	if err != nil {
		return "", "", fmt.Errorf("service layer  mongo delete token error %w", err)
	}
	accessToken, err := a.encodeToken(&model.User{Role: role, Login: login}, time.Now().Add(time.Minute*15).Unix(), "access")
	if err != nil {
		return "", "", fmt.Errorf("service layer  encode access token error %w", err)
	}
	refreshToken, err := a.encodeToken(&model.User{Role: role, Login: login}, time.Now().Add(time.Hour*24*7).Unix(), "refresh")
	if err != nil {
		return "", "", fmt.Errorf("service layer  encode refresh token error %w", err)
	}
	err = a.r.Create(context.Background(), *refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("service layer  mongo create token error %w", err)
	}
	return accessToken.Value, refreshToken.Value, nil
}
