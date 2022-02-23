package service

import (
	"context"
	"fmt"
	"time"

	"github.com/HekapOo-hub/Task1/internal/config"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository"
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// AuthService implements authentication and refresh token functional
type AuthService struct {
	r repository.TokenRepository
}

// NewAuthService returns instance of AuthService
func NewAuthService(r repository.TokenRepository) *AuthService {
	return &AuthService{r: r}
}

func (a *AuthService) getTokens(user *model.User) (accessToken, refreshToken *model.Token, err error) {
	claims1 := config.TokenClaims{
		Login: user.Login,
		Role:  user.Role,
		ID:    uuid.NewV4().String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.AccessTTL).Unix(),
		},
	}
	token1 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims1)
	accessValue, err := token1.SignedString([]byte(config.AccessKey))
	if err != nil {
		return nil, nil, fmt.Errorf("access token error with signing %w", err)
	}
	claims2 := config.TokenClaims{
		Login: user.Login,
		Role:  user.Role,
		ID:    uuid.NewV4().String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.RefreshTTL).Unix(),
		},
	}
	// Sign token and return
	token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims2)
	refreshValue, err := token2.SignedString([]byte(config.RefreshKey))
	if err != nil {
		return nil, nil, fmt.Errorf("refresh token error with signing %w", err)
	}
	return &model.Token{Value: accessValue, ExpiresAt: time.Now().Add(config.AccessTTL).Unix(), Login: user.Login},
		&model.Token{Value: refreshValue, ExpiresAt: time.Now().Add(config.RefreshTTL).Unix()}, nil
}

// Create is used for creating human info from db
func (a *AuthService) Create(ctx context.Context, token model.Token) error {
	err := a.r.Create(ctx, token)
	if err != nil {
		return fmt.Errorf("service layer create token error %w", err)
	}
	return nil
}

// Get is used for getting human info from db
func (a *AuthService) Get(ctx context.Context, token string) (*model.Token, error) {
	tokenFromDB, err := a.r.Get(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("service layer get token error %w", err)
	}
	return tokenFromDB, nil
}

// Delete is used for deleting human info from db
func (a *AuthService) Delete(ctx context.Context, login string) error {
	err := a.r.Delete(ctx, login)
	if err != nil {
		return fmt.Errorf("authentication layer delete token error %w", err)
	}
	return nil
}

// Authenticate finds user login and password and returns appropriate tokens
func (a *AuthService) Authenticate(ctx context.Context, user *model.User, password string) (accessValue, refreshValue string, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", fmt.Errorf("authentication comparing passwords error %w", err)
	}
	accessToken, refreshToken, err := a.getTokens(user)
	if err != nil {
		return "", "", fmt.Errorf("service layer authentication get tokens error %w", err)
	}
	err = a.r.Create(ctx, *refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("service layer mongo create token error %w", err)
	}
	return accessToken.Value, refreshToken.Value, nil
}

// Refresh returns new access and refresh tokens instead of old refresh token
func (a *AuthService) Refresh(ctx context.Context, claims *config.TokenClaims, token string) (accessValue, refreshValue string, err error) {
	role := claims.Role
	login := claims.Login
	err = a.r.Delete(ctx, token)
	if err != nil {
		return "", "", fmt.Errorf("service layer mongo delete token error %w", err)
	}
	accessToken, refreshToken, err := a.getTokens(&model.User{Role: role, Login: login})
	if err != nil {
		return "", "", fmt.Errorf("service layer get tokens error %w", err)
	}
	err = a.r.Create(ctx, *refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("service layer mongo create token error %w", err)
	}
	return accessToken.Value, refreshToken.Value, nil
}
