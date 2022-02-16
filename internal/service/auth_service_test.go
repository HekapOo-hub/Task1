package service

import (
	"context"
	"fmt"
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
	// no error
	tokenRepository = mocks.TokenRepository{}
	ctx := context.Background()
	token := model.Token{Value: "wsworgweeog[fjwq", ExpiresAt: 12341235, Login: "login"}
	authService := NewAuthService(&tokenRepository)
	tokenRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("model.Token")).Return(func(ctx context.Context, token model.Token) error {
		return nil
	})
	err := authService.Create(ctx, token)
	require.NoError(t, err)
	// create error
	tokenRepository = mocks.TokenRepository{}
	tokenRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("model.Token")).Return(func(ctx context.Context, token model.Token) error {
		return fmt.Errorf("this must be wrapped")
	})
	err = authService.Create(ctx, token)
	require.Equal(t, fmt.Errorf("service layer create token error %w",
		fmt.Errorf("this must be wrapped")), err)
}

func TestAuthService_Get(t *testing.T) {
	// no error
	tokenRepository = mocks.TokenRepository{}
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
	// get error
	tokenRepository = mocks.TokenRepository{}
	tokenRepository.On("Get", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(func(ctx context.Context, token string) *model.Token {
		return nil
	}, func(ctx context.Context, token string) error {
		return fmt.Errorf("this must be wrapped")
	})
	_, err = authService.Get(ctx, tokenStr)
	require.Equal(t, fmt.Errorf("service layer get token error %w",
		fmt.Errorf("this must be wrapped")), err)
}

func TestAuthService_Delete(t *testing.T) {
	// no error
	tokenRepository = mocks.TokenRepository{}
	ctx := context.Background()
	tokenStr := "tokenValue"
	authService := NewAuthService(&tokenRepository)
	tokenRepository.On("Delete", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(func(ctx context.Context, token string) error {
		return nil
	})
	err := authService.Delete(ctx, tokenStr)
	require.NoError(t, err)
	// delete error
	tokenRepository = mocks.TokenRepository{}
	tokenRepository.On("Delete", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(func(ctx context.Context, token string) error {
		return fmt.Errorf("this must be wrapped")
	})
	err = authService.Delete(ctx, tokenStr)
	require.Equal(t, fmt.Errorf("authentication layer delete token error %w",
		fmt.Errorf("this must be wrapped")), err)
}

func TestAuthService_Authenticate(t *testing.T) {
	// no error
	tokenRepository = mocks.TokenRepository{}
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
	// create error
	tokenRepository = mocks.TokenRepository{}
	tokenRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("model.Token")).Return(func(ctx context.Context, token model.Token) error {
		return fmt.Errorf("this must be wrapped")
	})
	_, _, err = authService.Authenticate(ctx, &user, password)
	require.Equal(t, fmt.Errorf("service layer mongo create token error %w",
		fmt.Errorf("this must be wrapped")), err)
}

func TestAuthService_Refresh(t *testing.T) {
	// no error
	tokenRepository = mocks.TokenRepository{}
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
	// delete error
	tokenRepository = mocks.TokenRepository{}
	tokenRepository.On("Delete", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(func(ctx context.Context, token string) error {
		return fmt.Errorf("this must be wrapped")
	})
	_, _, err = authService.Refresh(ctx, claims, tokenStr)
	require.Equal(t, fmt.Errorf("service layer mongo delete token error %w",
		fmt.Errorf("this must be wrapped")), err)
	// create error
	tokenRepository = mocks.TokenRepository{}
	tokenRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("model.Token")).Return(func(ctx context.Context, token model.Token) error {
		return fmt.Errorf("this must be wrapped")
	})
	tokenRepository.On("Delete", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(func(ctx context.Context, token string) error {
		return nil
	})
	_, _, err = authService.Refresh(ctx, claims, tokenStr)
	require.Equal(t, fmt.Errorf("service layer mongo create token error %w",
		fmt.Errorf("this must be wrapped")), err)

}
