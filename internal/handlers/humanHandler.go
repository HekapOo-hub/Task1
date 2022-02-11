// Package handlers contains handlers for echo server
package handlers

import (
	"net/http"

	"github.com/HekapOo-hub/Task1/internal/config"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/request"
	"github.com/HekapOo-hub/Task1/internal/service"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// HumanHandler implements  crud interface with Human entity for echo server
type HumanHandler struct {
	humanService *service.HumanService
	authService  *service.AuthService
}

// NewHumanHandler creates new human handler
func NewHumanHandler(hs *service.HumanService, as *service.AuthService) *HumanHandler {
	return &HumanHandler{humanService: hs, authService: as}
}

// Create is used for creating human info in db
// @Summary create human
// @Security ApiKeyAuth
// @Tags human
// @Description to create new human
// @Accept json
// @Param request body request.CreateHumanRequest true "create human info"
// @Success 201 body string
// @Failure 400 body echo.NewHTTPError
// @Router /human/create [post]
func (h *HumanHandler) Create(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.TokenClaims)
	role := claims.Role
	req := new(request.CreateHumanRequest)
	if err := c.Bind(req); err != nil {
		log.Warnf("error in binding structure with env variables in create: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	if err := c.Validate(req); err != nil {
		log.Warnf("validation create human error^ %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if role != admin {
		log.Warn("access denied")
		return echo.NewHTTPError(http.StatusNotAcceptable, "access denied")
	}

	err := h.humanService.Create(c.Request().Context(), model.Human{Name: req.Name, Male: req.Male, Age: req.Age})
	if err != nil {
		log.Warnf("error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "human info was created")
}

// Update is used for updating human info from db by his ID
// @Summary update human
// @Security ApiKeyAuth
// @Tags human
// @Description to update human
// @Accept json
// @Param request body request.UpdateHumanRequest true "update human info"
// @Success 200 body string
// @Failure 400 body echo.NewHTTPError
// @Router /human/update [patch]
func (h *HumanHandler) Update(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.TokenClaims)
	role := claims.Role
	req := new(request.UpdateHumanRequest)
	if err := c.Bind(req); err != nil {
		log.Warnf("error in binding structure with env variables in update: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	if err := c.Validate(req); err != nil {
		log.Warnf("validation update human error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if role != admin {
		log.Warn("access denied")
		return echo.NewHTTPError(http.StatusNotAcceptable, "access denied")
	}
	err := h.humanService.Update(c.Request().Context(), req.OldName, model.Human{Name: req.NewName,
		Male: req.NewMale, Age: req.NewAge})
	if err != nil {
		log.WithField("error", err).Warn()
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "human info was updated")
}

// Get is used for getting human info from db by his name
// @Summary get human info
// @Security ApiKeyAuth
// @Tags human
// @Description to get human info
// @Param name path string true "get human info"
// @Success 200 body string
// @Failure 400 body echo.NewHTTPError
// @Router /human/get/{name} [get]
func (h *HumanHandler) Get(c echo.Context) error {
	name := c.Param("name")
	req := request.GetHumanRequest{Name: name}
	if err := c.Validate(req); err != nil {
		log.Warnf("validation get human error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	human, err := h.humanService.Get(c.Request().Context(), name)
	if err != nil {
		log.Warnf("error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, human.String())
}

// Delete is used for deleting human info from db by his ID
// @Summary delete human info
// @Security ApiKeyAuth
// @Tags human
// @Description to delete human
// @Param name path string true "delete human info"
// @Success 200 body string
// @Failure 400 body echo.NewHTTPError
// @Router /human/delete/{name} [delete]
func (h *HumanHandler) Delete(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.TokenClaims)
	role := claims.Role
	name := c.Param("name")
	req := request.DeleteHumanRequest{Name: name}
	if err := c.Validate(req); err != nil {
		log.Warnf("validation delete human error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if role != admin {
		log.Warn("access denied")
		return echo.NewHTTPError(http.StatusNotAcceptable, "access denied")
	}
	err := h.humanService.Delete(c.Request().Context(), req.Name)
	if err != nil {
		log.Warnf("error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "human's info was deleted")
}
