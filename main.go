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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
	fmt.Println("Enter 1 to create Postgres db, 2- to create mongo db")
	n := os.Getenv("CASE")
	h := &handlers.Handler{}

	switch n {
	case "1":
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
		h = handlers.NewHandler(service.NewService(repo))

	case "2":
		uri, err := config.GetMongoURI()
		if err != nil {
			return
		}
		conn, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
		if err != nil {
			log.WithFields(log.Fields{"error": err, "uri": uri}).Warn("error with connecting to mongodb")
			return
		}
		defer repository.Disconnect(context.Background(), conn)
		repo := repository.NewMongoRepository(conn)
		h = handlers.NewHandler(service.NewService(repo))
	default:
		log.WithField("number", n).Warn("you entered inappropriate number ")
		return
	}
	e := echo.New()
	e.POST("/create", h.CreateHuman)
	e.GET("/get", h.GetHuman)
	e.PATCH("/update", h.UpdateHuman)
	e.DELETE("/delete", h.DeleteHuman)
	err := e.Start(":1323")
	if err != nil {
		log.WithField("error", err).Warn("error with starting an echo server")
		return
	}
}
