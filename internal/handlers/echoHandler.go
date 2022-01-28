package handlers

import (
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/request"
	"github.com/HekapOo-hub/Task1/internal/service"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type HumanHandler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *HumanHandler {
	return &HumanHandler{service: s}
}
func (h *HumanHandler) Create(c echo.Context) error {
	human := new(request.CreateRequest)
	if err := c.Bind(human); err != nil {
		log.WithField("error", err).Warn("error in binding structure with env variables in create")
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err := h.service.Create(model.Human{Name: human.Name, Male: human.Male, Age: human.Age})
	if err != nil {
		log.WithField("error", err).Warn()
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "human info was created")
}
func (h *HumanHandler) Update(c echo.Context) error {
	human := new(request.UpdateRequest)
	if err := c.Bind(human); err != nil {
		log.WithField("error", err).Warn("error in binding structure with env variables in update")
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err := h.service.Update(human.Id, model.Human{Name: human.NewName,
		Male: human.NewMale, Age: human.NewAge})
	if err != nil {
		log.WithField("error", err).Warn()
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "human info was updated")
}
func (h *HumanHandler) Get(c echo.Context) error {
	name := c.Param("name")
	human, err := h.service.Get(name)
	if err != nil {
		log.WithField("error", err).Warn()
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, human.String())
}
func (h *HumanHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	err := h.service.Delete(id)
	if err != nil {
		log.WithField("error", err).Warn()
		return echo.NewHTTPError(http.StatusNoContent, err.Error())
	}
	return c.String(http.StatusOK, "human's info was deleted")
}
