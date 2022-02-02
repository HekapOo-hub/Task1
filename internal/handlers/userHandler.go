package handlers

import (
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/request"
	"github.com/HekapOo-hub/Task1/internal/service"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UserHandler struct {
	userService *service.UserService
	authService *service.AuthService
}

const (
	admin = "admin"
)

func NewUserHandler(us *service.UserService, as *service.AuthService) *UserHandler {
	return &UserHandler{userService: us, authService: as}
}

func (u *UserHandler) Authenticate(c echo.Context) error {
	login := c.QueryParam("login")
	password := c.QueryParam("password")
	user, err := u.userService.Get(login)
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
	login := c.QueryParam("login")
	password := c.QueryParam("password")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*service.TokenClaims)
	role := claims.Role
	if role != admin {
		return echo.NewHTTPError(http.StatusBadRequest, "access denied")
	}
	err := u.userService.Create(login, password)
	if err != nil {
		log.WithField("error", err).Warn("error with creating user")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "user was created")
}

func (u *UserHandler) Get(c echo.Context) error {
	loginToGet := c.Param("login")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*service.TokenClaims)
	role := claims.Role
	login := claims.Login
	if login != loginToGet && role != admin {
		return echo.NewHTTPError(http.StatusBadRequest, "access denied")
	}
	res, err := u.userService.Get(loginToGet)
	if err != nil {
		log.WithField("error", err).Warn("error in getting user from db")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, res.String())

}

func (u *UserHandler) Update(c echo.Context) error {
	info := new(request.UpdateUserRequest)
	if err := c.Bind(info); err != nil {
		log.WithField("error", err).Warn("error in binding structure with env variables in update user")
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*service.TokenClaims)
	role := claims.Role
	login := claims.Login

	if login != info.OldLogin && role != admin {
		return echo.NewHTTPError(http.StatusBadRequest, "access denied")
	}

	err := u.userService.Update(info.OldLogin, model.User{Login: info.NewLogin, Password: info.NewPassword})
	if err != nil {
		log.WithField("error", err).Warn("error in update user. layer:handler")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "user was updated")
}

func (u *UserHandler) Delete(c echo.Context) error {
	loginToDelete := c.Param("login")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*service.TokenClaims)
	role := claims.Role
	login := claims.Login

	if login != loginToDelete && role != admin {
		return fmt.Errorf("access denied in delete")
	}
	err := u.userService.Delete(loginToDelete)
	if err != nil {
		log.WithField("error", err).Warn("error in delete handler layer")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "user was deleted")
}

func (u *UserHandler) Refresh(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*service.TokenClaims)
	accessToken, refreshToken, err := u.authService.Refresh(claims, user.Raw)
	if err != nil {
		log.WithField("error", err).Warn("refresh token error")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, fmt.Sprintf("access token: %s\nrefresh token: %s", accessToken, refreshToken))
}

func (u *UserHandler) LogOut(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)

	err := u.authService.Delete(user.Raw)
	if err != nil {
		log.WithField("error", err).Warn("log out: delete token error")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "user logged out")
}
