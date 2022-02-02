package handlers

import (
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/request"
	"github.com/HekapOo-hub/Task1/internal/service"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type HumanHandler struct {
	humanService *service.HumanService
	authService  *service.AuthService
}

func NewHumanHandler(hs *service.HumanService, as *service.AuthService) *HumanHandler {
	return &HumanHandler{humanService: hs, authService: as}
}
func (h *HumanHandler) Create(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*service.TokenClaims)
	role := claims.Role

	req := new(request.CreateHumanRequest)
	if err := c.Bind(req); err != nil {
		log.WithField("error", err).Warn("error in binding structure with env variables in create")
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}
	if role != admin {
		log.Warn("access denied")
		return echo.NewHTTPError(http.StatusNotAcceptable, "access denied")
	}

	err := h.humanService.Create(model.Human{Name: req.Name, Male: req.Male, Age: req.Age})
	if err != nil {
		log.WithField("error", err).Warn()
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "human info was created")
}
func (h *HumanHandler) Update(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*service.TokenClaims)
	role := claims.Role

	req := new(request.UpdateHumanRequest)
	if err := c.Bind(req); err != nil {
		log.WithField("error", err).Warn("error in binding structure with env variables in update")
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}
	if role != admin {
		log.Warn("access denied")
		return echo.NewHTTPError(http.StatusNotAcceptable, "access denied")
	}
	err := h.humanService.Update(req.Id, model.Human{Name: req.NewName,
		Male: req.NewMale, Age: req.NewAge})
	if err != nil {
		log.WithField("error", err).Warn()
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "human info was updated")
}
func (h *HumanHandler) Get(c echo.Context) error {
	name := c.Param("name")

	human, err := h.humanService.Get(name)
	if err != nil {
		log.WithField("error", err).Warn()
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, human.String())
}
func (h *HumanHandler) Delete(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*service.TokenClaims)
	role := claims.Role

	id := c.Param("id")
	if role != admin {
		log.Warn("access denied")
		return echo.NewHTTPError(http.StatusNotAcceptable, "access denied")
	}
	err := h.humanService.Delete(id)
	if err != nil {
		log.WithField("error", err).Warn()
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "human's info was deleted")
}
