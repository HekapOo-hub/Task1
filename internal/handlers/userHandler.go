package handlers

import (
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/request"
	"github.com/HekapOo-hub/Task1/internal/service"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (u *UserHandler) Authenticate(c echo.Context) error {
	login := c.QueryParam("login")
	password := c.QueryParam("password")
	token, err := u.service.Authentication(login, password)
	if err != nil {
		log.WithField("error", err).Warn("error with token in authentication")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(1 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "you were authenticated!")
}

func (u *UserHandler) Create(c echo.Context) error {
	login := c.QueryParam("login")
	password := c.QueryParam("password")
	err := u.service.CreateUser(login, password)
	if err != nil {
		log.WithField("error", err).Warn("error with creating user")
		return echo.NewHTTPError(http.StatusBadRequest, "error with creating user")
	}
	return c.String(http.StatusOK, "user was created")
}

func (u *UserHandler) Get(c echo.Context) error {
	login := c.Param("login")
	token, err := c.Cookie("token")
	if err != nil {
		log.WithField("error", err).Warn("error in fetching cookie in get")
		return echo.NewHTTPError(http.StatusBadRequest, "error in fetching cookie in get")
	}
	user, err := u.service.Get(token.Value, login)
	if err != nil {
		log.WithField("error", err).Warn("error in getting user from db")
		return echo.NewHTTPError(http.StatusBadRequest, "error in getting user from db")
	}
	return c.String(http.StatusOK, user.String())
}

func (u *UserHandler) Update(c echo.Context) error {
	token, err := c.Cookie("token")
	if err != nil {
		log.WithField("error", err).Warn("error in fetching cookie in update")
		return echo.NewHTTPError(http.StatusBadRequest, "error in fetching cookie in update")
	}
	info := new(request.UpdateUserRequest)
	if err := c.Bind(info); err != nil {
		log.WithField("error", err).Warn("error in binding structure with env variables in update")
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = u.service.UpdateUser(token.Value, info.OldLogin, model.User{Login: info.NewLogin, Password: info.NewPassword})
	if err != nil {
		log.WithField("error", err).Warn("error in update handler layer")
		return fmt.Errorf("error in update handler layer %w", err)
	}
	return c.String(http.StatusOK, "user was updated")
}

func (u *UserHandler) Delete(c echo.Context) error {
	loginToDelete := c.Param("login")
	token, err := c.Cookie("token")
	if err != nil {
		log.WithField("error", err).Warn("error in fetching cookie in delete")
		return echo.NewHTTPError(http.StatusBadRequest, "error in fetching cookie in delete")
	}
	err = u.service.DeleteUser(token.Value, loginToDelete)
	if err != nil {
		log.WithField("error", err).Warn("error in delete handler layer")
		return fmt.Errorf("error in delete handler layer %w", err)
	}
	return c.String(http.StatusOK, "user was deleted")
}
