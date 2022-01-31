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

type HumanHandler struct {
	humanService *service.HumanService
	authService  *service.AuthService
}

func NewHumanHandler(hs *service.HumanService, as *service.AuthService) *HumanHandler {
	return &HumanHandler{humanService: hs, authService: as}
}
func (h *HumanHandler) Create(c echo.Context) error {
	req := new(request.CreateHumanRequest)
	if err := c.Bind(req); err != nil {
		log.WithField("error", err).Warn("error in binding structure with env variables in create")
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}
	_, role, err := h.authService.Authorize(req.Token)
	if err != nil {
		log.WithField("error", err).Warn("error in authorization in create human")
		return fmt.Errorf("error in authorization in create human %w", err)
	}
	err = h.humanService.Create(role, model.Human{Name: req.Name, Male: req.Male, Age: req.Age})
	if err != nil {
		log.WithField("error", err).Warn()
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "human info was created")
}
func (h *HumanHandler) Update(c echo.Context) error {

	req := new(request.UpdateHumanRequest)
	if err := c.Bind(req); err != nil {
		log.WithField("error", err).Warn("error in binding structure with env variables in update")
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}
	_, role, err := h.authService.Authorize(req.Token)
	if err != nil {
		log.WithField("error", err).Warn("error in authorization in update human")
		return fmt.Errorf("error in authorization in update human %w", err)
	}
	err = h.humanService.Update(role, req.Id, model.Human{Name: req.NewName,
		Male: req.NewMale, Age: req.NewAge})
	if err != nil {
		log.WithField("error", err).Warn()
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "human info was updated")
}
func (h *HumanHandler) Get(c echo.Context) error {
	token := c.QueryParam("token")
	name := c.QueryParam("name")
	_, role, err := h.authService.Authorize(token)
	if err != nil {
		log.WithField("error", err).Warn("error in authorization in get human")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	human, err := h.humanService.Get(role, name)
	if err != nil {
		log.WithField("error", err).Warn()
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, human.String())
}
func (h *HumanHandler) Delete(c echo.Context) error {
	token := c.QueryParam("token")
	id := c.QueryParam("id")
	_, role, err := h.authService.Authorize(token)
	if err != nil {
		log.WithField("error", err).Warn("error in authorization in delete human")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = h.humanService.Delete(role, id)
	if err != nil {
		log.WithField("error", err).Warn()
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "human's info was deleted")
}
