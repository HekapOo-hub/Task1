package main

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/config"
	"github.com/HekapOo-hub/Task1/internal/handlers"
	"github.com/HekapOo-hub/Task1/internal/repository"
	"github.com/HekapOo-hub/Task1/internal/service"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	pool, err := pgxpool.Connect(context.Background(), cfg.GetURL())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer pool.Close()
	repo := repository.NewRepository(pool)
	h := handlers.NewHandler(service.Service{Repo: repo})
	e := echo.New()
	e.POST("/create", h.CreateHuman)
	e.GET("/get", h.GetHuman)
	e.GET("/update", h.UpdateHuman)
	e.GET("/delete", h.DeleteHuman)
	err = e.Start(":1323")
	if err != nil {
		fmt.Println(err)
	}
}
