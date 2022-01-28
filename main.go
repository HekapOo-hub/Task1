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
	h := handlers.NewHumanHandler(service.NewService(repo))

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
	repo2 := repository.NewMongoUserRepository(conn)
	h2 := handlers.NewUserHandler(service.NewUserService(repo2))

	e := echo.New()
	e.POST("/create", h.Create)
	e.GET("/get/:name", h.Get)
	e.PATCH("/update", h.Update)
	e.DELETE("/delete/:id", h.Delete)
	e.GET("/getUser/:login", h2.Get)
	e.POST("/createUser/", h2.Create)
	e.PATCH("/updateUser/", h2.Update)
	e.GET("/enter", h2.Authenticate)
	e.DELETE("/deleteUser/:login", h2.Delete)
	err = e.Start(":1323")
	if err != nil {
		log.WithField("error", err).Warn("error with starting an echo server")
		return
	}
}
