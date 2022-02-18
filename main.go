package main

import (
	"context"
	"os"

	_ "github.com/HekapOo-hub/Task1/docs"
	"github.com/HekapOo-hub/Task1/internal/config"
	"github.com/HekapOo-hub/Task1/internal/handlers"
	"github.com/HekapOo-hub/Task1/internal/repository"
	"github.com/HekapOo-hub/Task1/internal/service"
	"github.com/HekapOo-hub/Task1/internal/validation"
	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title Human CRUD API
// @version 1.0
// @description Human CRUD API with authorization and cache
// @licence.name Apache 2.0
// @host localhost:1323
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
	cfg, err := config.NewPostgresConfig()
	if err != nil {
		log.Warnf("postgres config error: %v", err)
		return
	}
	pool, err := pgxpool.Connect(context.Background(), cfg.GetURL())
	if err != nil {
		log.Warnf("postgres connect error: %v", err)
		return
	}
	defer pool.Close()

	repo := repository.NewHumanRepository(pool)

	uri, err := config.GetMongoURI()
	if err != nil {
		log.Warnf("error: %v", err)
		return
	}
	conn, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.WithField("uri", uri).Warnf("error with connecting to mongodb: %v", err)
		return
	}
	defer repository.MongoDisconnect(context.Background(), conn)
	redisCfg, err := config.NewRedisConfig()
	if err != nil {
		log.Warnf("redis get config error: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})

	ctx, cancel := context.WithCancel(context.Background())
	redisCacheHumanRepository := repository.NewRedisHumanCacheRepository(ctx, redisClient)
	defer cancel()
	userRepo := repository.NewMongoUserRepository(conn)
	h := handlers.NewHumanHandler(service.NewHumanService(repo, redisCacheHumanRepository),
		service.NewAuthService(repository.NewMongoTokenRepository(conn)))
	h2 := handlers.NewUserHandler(service.NewUserService(userRepo),
		service.NewAuthService(repository.NewMongoTokenRepository(conn)))
	h3 := handlers.FileHandler{}
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	validator, err := validation.NewValidator()
	if err != nil {
		log.Warnf("echo validator error %v", err)
		return
	}
	e.Validator = validator
	accessGroup1 := e.Group("/user/", middleware.JWTWithConfig(config.GetAccessTokenConfig()))
	accessGroup2 := e.Group("/human/", middleware.JWTWithConfig(config.GetAccessTokenConfig()))
	refreshGroup := e.Group("/refresh/", middleware.JWTWithConfig(config.GetRefreshTokenConfig()))
	accessGroup2.POST("create", h.Create)
	accessGroup2.GET("get/:name", h.Get)
	accessGroup2.PATCH("update", h.Update)
	accessGroup2.DELETE("delete/:name", h.Delete)
	accessGroup1.GET("get/:login", h2.Get)
	accessGroup1.POST("create", h2.Create)
	accessGroup1.PATCH("update", h2.Update)
	e.GET("/signIn", h2.Authenticate)
	accessGroup1.GET("file/download/:fileName", h3.Download)
	accessGroup1.GET("file/upload/:fileName", h3.Upload)
	accessGroup1.DELETE("delete/:login", h2.Delete)
	refreshGroup.GET("update", h2.Refresh)
	refreshGroup.DELETE("logOut", h2.LogOut)
	err = e.Start(":1323")
	if err != nil {
		log.Warnf("error with starting an echo server: %v", err)
		return
	}

}
