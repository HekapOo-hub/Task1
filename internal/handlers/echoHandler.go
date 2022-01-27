package handlers

import (
	"context"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	service service.Service
}

func NewHandler(s service.Service) *Handler {
	return &Handler{service: s}
}
func (h *Handler) CreateHuman(c echo.Context) error {
	human := new(model.Human)
	if err := c.Bind(human); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err := h.service.Create(context.Background(), *human)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "human info was created")
}
func (h *Handler) UpdateHuman(c echo.Context) error {
	human := new(model.Human)
	if err := c.Bind(human); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err := h.service.Update(context.Background(), human.Id, *human)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "human info was updated")
}
func (h *Handler) GetHuman(c echo.Context) error {
	id := new(model.Id)
	if err := c.Bind(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	human, err := h.service.Get(context.Background(), id.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, human.String())
}
func (h *Handler) DeleteHuman(c echo.Context) error {
	id := new(model.Id)
	if err := c.Bind(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err := h.service.Delete(context.Background(), id.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNoContent, err.Error())
	}
	return c.String(http.StatusOK, "human's info was deleted")
}
