package main

import (
	"context"
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
		log.WithField("error", err).Warn("postgres config error")
		return
	}
	pool, err := pgxpool.Connect(context.Background(), cfg.GetURL())
	if err != nil {
		log.WithField("error", err).Warn("postgres connect error")
		return
	}
	defer pool.Close()
	repo := repository.NewRepository(pool)

	uri, err := config.GetMongoURI()
	if err != nil {
		return
	}
	conn, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.WithFields(log.Fields{"error": err, "uri": uri}).Warn("error with connecting to mongodb")
		return
	}
	defer repository.MongoDisconnect(context.Background(), conn)
	userRepo := repository.NewMongoUserRepository(conn)
	h := handlers.NewHumanHandler(service.NewService(repo),
		service.NewAuthService(repository.NewMongoTokenRepository(conn)))
	h2 := handlers.NewUserHandler(service.NewUserService(userRepo),
		service.NewAuthService(repository.NewMongoTokenRepository(conn)))

	e := echo.New()
	e.POST("/human/create", h.Create)
	e.GET("/human/get", h.Get)
	e.PATCH("/human/update", h.Update)
	e.DELETE("/human/delete", h.Delete)
	e.GET("/user/get", h2.Get)
	e.POST("/user/create", h2.Create)
	e.PATCH("/user/update", h2.Update)
	e.GET("/signIn", h2.Authenticate)
	e.DELETE("/user/delete", h2.Delete)
	e.GET("/refresh/:token", h2.Refresh)
	e.DELETE("/logOut/:token", h2.LogOut)
	err = e.Start(":1323")
	if err != nil {
		log.WithField("error", err).Warn("error with starting an echo server")
		return
	}
}
