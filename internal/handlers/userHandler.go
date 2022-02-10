package handlers

import (
	"fmt"
	"net/http"

	"github.com/HekapOo-hub/Task1/internal/config"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/request"
	"github.com/HekapOo-hub/Task1/internal/service"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

const (
	admin = "admin"
)

// UserHandler implements crud interface for working with db and is used to define echo server's handler functions
type UserHandler struct {
	userService *service.UserService
	authService *service.AuthService
}

// NewUserHandler creates new user handler
func NewUserHandler(us *service.UserService, as *service.AuthService) *UserHandler {
	return &UserHandler{userService: us, authService: as}
}

// Authenticate checks if user is existing in db and in positive case returns access and refresh tokens
func (u *UserHandler) Authenticate(c echo.Context) error {
	login := c.QueryParam("login")
	password := c.QueryParam("password")
	req := request.SignInRequest{Login: login, Password: password}
	if err := c.Validate(req); err != nil {
		log.Warnf("validation sign in error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := u.userService.Get(c.Request().Context(), login)
	if err != nil {
		log.Warnf("get user error in authenticate: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	accessToken, refreshToken, err := u.authService.Authenticate(c.Request().Context(), user, password)
	if err != nil {
		log.Warnf("error with token in authentication: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, fmt.Sprintf("access token: %s\nrefresh token: %s", accessToken, refreshToken))
}

// Create is used for creating user in db
func (u *UserHandler) Create(c echo.Context) error {
	req := new(request.CreateUserRequest)
	if err := c.Bind(req); err != nil {
		log.Warnf("error in binding structure with env variables in create user: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	if err := c.Validate(req); err != nil {
		log.Warnf("validation create user error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.TokenClaims)
	role := claims.Role
	if role != admin {
		return echo.NewHTTPError(http.StatusBadRequest, "access denied")
	}
	err := u.userService.Create(c.Request().Context(), req.Login, req.Password)
	if err != nil {
		log.Warnf("error with creating user: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "user was created")
}

// Get is used for getting user info from db
func (u *UserHandler) Get(c echo.Context) error {
	loginToGet := c.Param("login")
	req := request.GetUserRequest{Login: loginToGet}
	if err := c.Validate(req); err != nil {
		log.WithField("error", err).Warn("validation get user error")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.TokenClaims)
	role := claims.Role
	login := claims.Login
	if login != loginToGet && role != admin {
		return echo.NewHTTPError(http.StatusBadRequest, "access denied")
	}
	res, err := u.userService.Get(c.Request().Context(), loginToGet)
	if err != nil {
		log.WithField("error", err).Warn("error in getting user from db")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, res.String())
}

// Update is used for updating user info in db
func (u *UserHandler) Update(c echo.Context) error {
	info := new(request.UpdateUserRequest)
	if err := c.Bind(info); err != nil {
		log.WithField("error", err).Warn("error in binding structure with env variables in update user")
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	if err := c.Validate(info); err != nil {
		log.WithField("error", err).Warn("validation update user error")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.TokenClaims)
	role := claims.Role
	login := claims.Login

	if login != info.OldLogin && role != admin {
		return echo.NewHTTPError(http.StatusBadRequest, "access denied")
	}

	err := u.userService.Update(c.Request().Context(), info.OldLogin, model.User{Login: info.NewLogin, Password: info.NewPassword})
	if err != nil {
		log.WithField("error", err).Warn("error in update user. layer:handler")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "user was updated")
}

// Delete is used for deleting user from db
func (u *UserHandler) Delete(c echo.Context) error {
	loginToDelete := c.Param("login")
	req := request.DeleteUserRequest{Login: loginToDelete}
	if err := c.Validate(req); err != nil {
		log.WithField("error", err).Warn("validation delete user error")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.TokenClaims)
	role := claims.Role
	login := claims.Login

	if login != loginToDelete && role != admin {
		return fmt.Errorf("access denied in delete")
	}
	err := u.userService.Delete(c.Request().Context(), loginToDelete)
	if err != nil {
		log.WithField("error", err).Warn("error in delete handler layer")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "user was deleted")
}

// Refresh returns new access and refresh token instead of old refresh token
func (u *UserHandler) Refresh(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.TokenClaims)
	accessToken, refreshToken, err := u.authService.Refresh(c.Request().Context(), claims, user.Raw)
	if err != nil {
		log.WithField("error", err).Warn("refresh token error")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, fmt.Sprintf("access token: %s\nrefresh token: %s", accessToken, refreshToken))
}

// LogOut deletes refresh token from db
func (u *UserHandler) LogOut(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	err := u.authService.Delete(c.Request().Context(), user.Raw)
	if err != nil {
		log.WithField("error", err).Warn("log out: delete token error")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "user logged out")
}
