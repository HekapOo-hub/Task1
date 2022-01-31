package handlers

import (
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/request"
	"github.com/HekapOo-hub/Task1/internal/service"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UserHandler struct {
	userService *service.UserService
	authService *service.AuthService
}

func NewUserHandler(us *service.UserService, as *service.AuthService) *UserHandler {
	return &UserHandler{userService: us, authService: as}
}

func (u *UserHandler) Authenticate(c echo.Context) error {
	login := c.QueryParam("login")
	password := c.QueryParam("password")
	user, err := u.userService.Get("", "admin", login)
	if err != nil {
		log.WithField("error", err).Warn("get user error in authenticate")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	accessToken, refreshToken, err := u.authService.Authenticate(user, password)
	if err != nil {
		log.WithField("error", err).Warn("error with token in authentication")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, fmt.Sprintf("access token: %s\nrefresh token: %s", accessToken, refreshToken))
}

func (u *UserHandler) Create(c echo.Context) error {
	user := new(request.CreateUserRequest)
	if err := c.Bind(user); err != nil {
		log.WithField("error", err).Warn("error in binding structure with env variables in create user")
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	_, role, err := u.authService.Authorize(user.Token)
	if err != nil {
		log.WithField("error", err).Warn("error in authorization in create user")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if role != "admin" {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = u.userService.Create(user.Login, user.Password)
	if err != nil {
		log.WithField("error", err).Warn("error with creating user")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "user was created")
}

func (u *UserHandler) Get(c echo.Context) error {
	loginToGet := c.QueryParam("login")
	token := c.QueryParam("token")
	login, role, err := u.authService.Authorize(token)
	if err != nil {
		log.WithField("error", err).Warn("error in authorization in get user")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := u.userService.Get(login, role, loginToGet)
	if err != nil {
		log.WithField("error", err).Warn("error in getting user from db")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, user.String())

}

func (u *UserHandler) Update(c echo.Context) error {

	info := new(request.UpdateUserRequest)
	if err := c.Bind(info); err != nil {
		log.WithField("error", err).Warn("error in binding structure with env variables in update user")
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	login, role, err := u.authService.Authorize(info.Token)
	if err != nil {
		log.WithField("error", err).Warn("error in authorization in create user")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = u.userService.Update(login, role, info.OldLogin, model.User{Login: info.NewLogin, Password: info.NewPassword})
	if err != nil {
		log.WithField("error", err).Warn("error in update user. layer:handler")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "user was updated")
}

func (u *UserHandler) Delete(c echo.Context) error {
	loginToDelete := c.QueryParam("login")
	token := c.QueryParam("token")
	login, role, err := u.authService.Authorize(token)
	if err != nil {
		log.WithField("error", err).Warn("error in authorization in create user")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = u.userService.Delete(login, role, loginToDelete)
	if err != nil {
		log.WithField("error", err).Warn("error in delete handler layer")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "user was deleted")
}

func (u *UserHandler) Refresh(c echo.Context) error {
	token := c.Param("token")
	accessToken, refreshToken, err := u.authService.Refresh(token)
	if err != nil {
		log.WithField("error", err).Warn("refresh token error")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, fmt.Sprintf("access token: %s\nrefresh token: %s", accessToken, refreshToken))
}

func (u *UserHandler) LogOut(c echo.Context) error {
	token := c.Param("token")
	err := u.authService.Delete(token)
	if err != nil {
		log.WithField("error", err).Warn("log out: delete token error")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
