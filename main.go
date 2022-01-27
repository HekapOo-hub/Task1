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
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
	var n int
	fmt.Println("Enter 1 to create Postgres db, 2- to create mongo db")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Errorf("scan error in main. %s", err.Error())
		return
	}
	switch n {
	case 1:
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
		e.PATCH("/update", h.UpdateHuman)
		e.DELETE("/delete", h.DeleteHuman)
		err = e.Start(":1323")
		if err != nil {
			fmt.Println(err)
		}
	case 2:

	default:
		log.WithField("number", n).Warn("you entered inappropriate number ")
		return
	}
}
