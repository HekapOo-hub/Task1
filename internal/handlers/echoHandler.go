package handlers

import (
	"context"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
		return echo.NewHTTPError(http.StatusNoContent, err)
	}
	return c.String(http.StatusCreated, "human info was created")
}
func (h *Handler) UpdateHuman(c echo.Context) error {
	name := c.QueryParam("name")
	maleStr := c.QueryParam("male")
	ageStr := c.QueryParam("age")
	idStr := c.QueryParam("id")
	var male bool
	if maleStr == "true" {
		male = true
	} else {
		male = false
	}
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	err = h.service.Update(context.Background(), id, model.Human{Name: name, Male: male, Age: age})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "human info was updated")
}
func (h *Handler) GetHuman(c echo.Context) error {
	idStr := c.QueryParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	human, err := h.service.Get(context.Background(), id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, human.String())
}
func (h *Handler) DeleteHuman(c echo.Context) error {
	idStr := c.QueryParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	err = h.service.Delete(context.Background(), id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "human's info was deleted")
}
