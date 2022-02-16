package service

import (
	"context"
	"github.com/HekapOo-hub/Task1/internal/config"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

var tokenRepository mocks.TokenRepository

func TestAuthService_Create(t *testing.T) {
	ctx := context.Background()
	token := model.Token{Value: "wsworgweeog[fjwq", ExpiresAt: 12341235, Login: "login"}
	authService := NewAuthService(&tokenRepository)
	tokenRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("model.Token")).Return(func(ctx context.Context, token model.Token) error {
		return nil
	})
	err := authService.Create(ctx, token)
	require.NoError(t, err)
}

func TestAuthService_Get(t *testing.T) {
	ctx := context.Background()
	tokenStr := "tokenValue"
	authService := NewAuthService(&tokenRepository)
	tokenRepository.On("Get", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(func(ctx context.Context, token string) *model.Token {
		return &model.Token{}
	}, func(ctx context.Context, token string) error {
		return nil
	})
	_, err := authService.Get(ctx, tokenStr)
	require.NoError(t, err)
}

func TestAuthService_Delete(t *testing.T) {
	ctx := context.Background()
	tokenStr := "tokenValue"
	authService := NewAuthService(&tokenRepository)
	tokenRepository.On("Delete", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(func(ctx context.Context, token string) error {
		return nil
	})
	err := authService.Delete(ctx, tokenStr)
	require.NoError(t, err)
}

func TestAuthService_Authenticate(t *testing.T) {
	ctx := context.Background()
	authService := NewAuthService(&tokenRepository)
	password := "123214"
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	require.NoError(t, err)
	user := model.User{Login: "admin",
		Password: string(hashedPwd)}
	tokenRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("model.Token")).Return(func(ctx context.Context, token model.Token) error {
		return nil
	})
	_, _, err = authService.Authenticate(ctx, &user, password)
	require.NoError(t, err)
}

func TestAuthService_Refresh(t *testing.T) {
	ctx := context.Background()
	authService := NewAuthService(&tokenRepository)
	tokenRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("model.Token")).Return(func(ctx context.Context, token model.Token) error {
		return nil
	})
	tokenRepository.On("Delete", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(func(ctx context.Context, token string) error {
		return nil
	})
	claims := &config.TokenClaims{Role: "user", Login: "login"}
	tokenStr := "wojg09wugqqerfwefg"
	_, _, err := authService.Refresh(ctx, claims, tokenStr)
	require.NoError(t, err)
}
