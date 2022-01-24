package handlers

import (
	"Task1/internal/model"
	"Task1/internal/service"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type EchoHandler struct {
	*echo.Echo
	Service service.Service
}

func (e *EchoHandler) Register() {
	e.GET("create", e.createHuman)
	e.GET("get", e.getHuman)
	e.GET("update", e.updateHuman)
	e.GET("delete", e.deleteHuman)
}
func (e *EchoHandler) createHuman(c echo.Context) error {
	name := c.QueryParam("name")
	maleStr := c.QueryParam("male")
	ageStr := c.QueryParam("age")
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
	err = e.Service.Create(context.Background(), model.Human{Name: name, Male: male, Age: age})
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "human info was created")
}
func (e *EchoHandler) updateHuman(c echo.Context) error {
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
	err = e.Service.Update(context.Background(), id, model.Human{Name: name, Male: male, Age: age})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "human info was updated")
}
func (e *EchoHandler) getHuman(c echo.Context) error {
	idStr := c.QueryParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	h, err := e.Service.Get(context.Background(), id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, h.String())
}
func (e *EchoHandler) deleteHuman(c echo.Context) error {
	idStr := c.QueryParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	err = e.Service.Delete(context.Background(), id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "human's info was deleted")
}
