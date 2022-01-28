package handlers

import (
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/service"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}
func (h *Handler) CreateHuman(c echo.Context) error {
	human := new(model.Human)
	if err := c.Bind(human); err != nil {
		log.WithField("error", err).Warn("error in binding structure with env variables in create")
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err := h.service.CreateHuman(*human)
	if err != nil {
		log.WithField("error", err).Warn()
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "human info was created")
}
func (h *Handler) UpdateHuman(c echo.Context) error {
	human := new(model.Human)
	if err := c.Bind(human); err != nil {
		log.WithField("error", err).Warn("error in binding structure with env variables in update")
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err := h.service.UpdateHuman(human.Id, *human)
	if err != nil {
		log.WithField("error", err).Warn()
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "human info was updated")
}
func (h *Handler) GetHuman(c echo.Context) error {
	name := new(model.Name)
	if err := c.Bind(name); err != nil {
		log.WithField("error", err).Warn("error in binding structure with env variables in get")
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	human, err := h.service.GetHumanInfo(name.Name)
	if err != nil {
		log.WithField("error", err).Warn()
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, human.String())
}
func (h *Handler) DeleteHuman(c echo.Context) error {
	id := new(model.Id)
	if err := c.Bind(id); err != nil {
		log.WithField("error", err).Warn("error in binding structure with env variables in delete")
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err := h.service.DeleteHuman(id.Id)
	if err != nil {
		log.WithField("error", err).Warn()
		return echo.NewHTTPError(http.StatusNoContent, err.Error())
	}
	return c.String(http.StatusOK, "human's info was deleted")
}
