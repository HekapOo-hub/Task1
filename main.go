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
	postgresClient, err := pgxpool.Connect(context.Background(), cfg.GetURL())
	if err != nil {
		log.Warnf("postgres connect error: %v", err)
		return
	}
	defer postgresClient.Close()

	uri, err := config.GetMongoURI()
	if err != nil {
		log.Warnf("error: %v", err)
		return
	}
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.WithField("uri", uri).Warnf("error with connecting to mongodb: %v", err)
		return
	}
	defer repository.MongoDisconnect(context.Background(), mongoClient)
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
	userRepo := repository.NewMongoUserRepository(mongoClient)
	humanHandler := handlers.NewHumanHandler(service.NewHumanService(repository.NewHumanRepository(postgresClient), redisCacheHumanRepository),
		service.NewAuthService(repository.NewMongoTokenRepository(mongoClient)))
	userHandler := handlers.NewUserHandler(service.NewUserService(userRepo),
		service.NewAuthService(repository.NewMongoTokenRepository(mongoClient)))
	fileHandler := &handlers.FileHandler{}
	echoServer := echo.New()
	echoServer.GET("/swagger/*", echoSwagger.WrapHandler)
	validator, err := validation.NewValidator()
	if err != nil {
		log.Warnf("echo validator error %v", err)
		return
	}
	echoServer.Validator = validator

	userAccessGroup := echoServer.Group("/user/", middleware.JWTWithConfig(config.GetAccessTokenConfig()))
	humanAccessGroup := echoServer.Group("/human/", middleware.JWTWithConfig(config.GetAccessTokenConfig()))
	refreshGroup := echoServer.Group("/refresh/", middleware.JWTWithConfig(config.GetRefreshTokenConfig()))
	humanAccessGroup.POST("create", humanHandler.Create)
	humanAccessGroup.GET("get/:name", humanHandler.Get)
	humanAccessGroup.PATCH("update", humanHandler.Update)
	humanAccessGroup.DELETE("delete/:name", humanHandler.Delete)
	userAccessGroup.GET("get/:login", userHandler.Get)
	userAccessGroup.POST("create", userHandler.Create)
	userAccessGroup.PATCH("update", userHandler.Update)
	echoServer.GET("/signIn", userHandler.Authenticate)
	userAccessGroup.GET("file/download/:fileName", fileHandler.Download)
	userAccessGroup.GET("file/upload/:fileName", fileHandler.Upload)
	userAccessGroup.DELETE("delete/:login", userHandler.Delete)
	refreshGroup.GET("update", userHandler.Refresh)
	refreshGroup.DELETE("logOut", userHandler.LogOut)

}
