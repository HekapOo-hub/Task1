package service

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository"
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var key = []byte("superSecretKey")

type tokenClaims struct {
	Login string
	Role  string
	Uuid  string
	jwt.StandardClaims
}

type AuthService struct {
	r repository.TokenRepository
}

func NewAuthService(r repository.TokenRepository) *AuthService {
	return &AuthService{r: r}
}
func (a *AuthService) encodeToken(user *model.User, expiresAt int64) (*model.Token, error) {
	claims := tokenClaims{
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
	return &model.Token{Value: val, ExpiresAt: expiresAt}, nil
}
func (a *AuthService) decodeToken(token string) (string, string, error) {
	// Parse the token
	tokenType, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return "", "", fmt.Errorf("token parsing error %w", err)
	}
	// Validate the token and return the custom claims
	claims, ok := tokenType.Claims.(*tokenClaims)
	if !ok {
		return "", "", fmt.Errorf("token parsing error %w", err)
	}
	if !tokenType.Valid {
		return "", "", fmt.Errorf("token expiration is over  %w", err)
	}
	return claims.Login, claims.Role, nil
}
func (a *AuthService) Create(token model.Token) error {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token.Value), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error with hashing token in create %w", err)
	}
	token.Value = string(hashedToken)
	err = a.r.Create(context.Background(), token)
	if err != nil {
		return fmt.Errorf("authentication layer create token error %w", err)
	}
	return nil
}
func (a *AuthService) Get(token string) (*time.Time, error) {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error with hashing token in get %w", err)
	}
	token = string(hashedToken)
	exAt, err := a.r.Get(context.Background(), token)
	if err != nil {
		return nil, fmt.Errorf("authentication layer get token error %w", err)
	}
	expireAt := time.Unix(exAt, 0)
	return &expireAt, nil
}
func (a *AuthService) Delete(token string) error {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error with hashing token in delete %w", err)
	}
	token = string(hashedToken)
	err = a.r.Delete(context.Background(), token)
	if err != nil {
		return fmt.Errorf("authentication layer delete token error %w", err)
	}
	return nil
}
func (a *AuthService) Authenticate(user *model.User, password string) (string, string, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", fmt.Errorf("authentication comparing passwords error %w", err)
	}
	accessToken, err := a.encodeToken(user, time.Now().Add(time.Minute*1).Unix())
	if err != nil {
		return "", "", fmt.Errorf("service layer authentication encode access token error %w", err)
	}
	refreshToken, err := a.encodeToken(user, time.Now().Add(time.Hour*24*7).Unix())
	if err != nil {
		return "", "", fmt.Errorf("service layer authentication encode refresh token error %w", err)
	}
	err = a.Create(*refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("service layer authentication mongo create token error %w", err)
	}
	return accessToken.Value, refreshToken.Value, nil
}
func (a *AuthService) Refresh(token string) (string, string, error) {
	login, role, err := a.decodeToken(token)
	if err != nil {
		return "", "", fmt.Errorf("service layer authentication decode token error in refresh %w", err)
	}
	_, err = a.Get(token)
	if err != nil {
		return "", "", fmt.Errorf("service layer authentication mongo get  token error %w\n", err)
	}
	err = a.r.Delete(context.Background(), token)
	if err != nil {
		return "", "", fmt.Errorf("service layer authentication mongo delete token error %w", err)
	}
	accessToken, err := a.encodeToken(&model.User{Role: role, Login: login}, time.Now().Add(time.Minute*15).Unix())
	if err != nil {
		return "", "", fmt.Errorf("service layer authentication encode access token error %w", err)
	}
	refreshToken, err := a.encodeToken(&model.User{Role: role, Login: login}, time.Now().Add(time.Hour*24*7).Unix())
	if err != nil {
		return "", "", fmt.Errorf("service layer authentication encode refresh token error %w", err)
	}
	err = a.r.Create(context.Background(), *refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("service layer authentication mongo create token error %w", err)
	}
	return accessToken.Value, refreshToken.Value, nil
}
func (a *AuthService) Authorize(token string) (string, string, error) {
	login, role, err := a.decodeToken(token)
	if err != nil {
		return "", "", fmt.Errorf("error with decodeing token in authorization %w", err)
	}
	return login, role, nil
}
