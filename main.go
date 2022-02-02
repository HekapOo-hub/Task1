package main

import (
	"context"
	"github.com/HekapOo-hub/Task1/internal/config"
	"github.com/HekapOo-hub/Task1/internal/handlers"
	"github.com/HekapOo-hub/Task1/internal/repository"
	"github.com/HekapOo-hub/Task1/internal/service"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	accessGroup1 := e.Group("/user/", middleware.JWTWithConfig(service.GetAccessTokenConfig()))
	accessGroup2 := e.Group("/human/", middleware.JWTWithConfig(service.GetAccessTokenConfig()))
	refreshGroup := e.Group("/refresh/", middleware.JWTWithConfig(service.GetRefreshTokenConfig()))

	accessGroup2.POST("create", h.Create)
	accessGroup2.GET("get/:name", h.Get)
	accessGroup2.PATCH("update", h.Update)
	accessGroup2.DELETE("delete/:id", h.Delete)
	accessGroup1.GET("get/:login", h2.Get)
	accessGroup1.POST("create", h2.Create)
	accessGroup1.PATCH("update", h2.Update)
	e.GET("/signIn", h2.Authenticate)
	accessGroup1.DELETE("delete/:login", h2.Delete)
	refreshGroup.GET("update", h2.Refresh)
	refreshGroup.DELETE("logOut", h2.LogOut)
	err = e.Start(":1323")
	if err != nil {
		log.WithField("error", err).Warn("error with starting an echo server")
		return
	}
}
