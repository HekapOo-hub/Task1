package main

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/config"
	"github.com/HekapOo-hub/Task1/internal/handlers"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository"
	"github.com/HekapOo-hub/Task1/internal/service"
	"github.com/HekapOo-hub/Task1/internal/validation"
	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	log "github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"os"
	"os/exec"
	"testing"
	"time"
)

var (
	postgresClient *pgxpool.Pool
	mongoClient    *mongo.Client
	redisClient    *redis.Client
)

func TestMain(m *testing.M) {
	startPostgres()
	startMongo()
	startRedis()
	hash, err := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("can't create hash password for user")
	}
	_, err = mongoClient.Database("myDatabase").Collection("users").InsertOne(context.Background(),
		model.User{Login: "admin", Password: string(hash), Role: "admin"})
	if err != nil {
		log.Errorf("can't create user in mongodb")
	}
	ctx, cancel := context.WithCancel(context.Background())
	redisCacheHumanRepository := repository.NewRedisHumanCacheRepository(ctx, redisClient)
	defer cancel()
	userRepo := repository.NewMongoUserRepository(mongoClient)
	h := handlers.NewHumanHandler(service.NewHumanService(repository.NewHumanRepository(postgresClient), redisCacheHumanRepository),
		service.NewAuthService(repository.NewMongoTokenRepository(mongoClient)))
	h2 := handlers.NewUserHandler(service.NewUserService(userRepo),
		service.NewAuthService(repository.NewMongoTokenRepository(mongoClient)))
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
	e.GET("/file/download/:fileName", h3.Download)
	e.GET("/file/upload/:fileName", h3.Upload)
	accessGroup1.DELETE("delete/:login", h2.Delete)
	refreshGroup.GET("update", h2.Refresh)
	refreshGroup.DELETE("logOut", h2.LogOut)
	err = e.Start(":1324")
	if err != nil {
		log.Warnf("error with starting an echo server: %v", err)
		return
	}
}

func startPostgres() {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Errorf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	postgresResource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env: []string{
			"POSTGRES_PASSWORD=1234",
			"POSTGRES_USER=test",
			"POSTGRES_DB=testDB",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Errorf("Could not start resource: %s", err)
	}

	hostAndPort := postgresResource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgresql://test:1234@%s/testDB?sslmode=disable", hostAndPort)
	flywayURL := fmt.Sprintf("jdbc:postgresql://%s/testDB", hostAndPort)
	log.Info("Connecting to database on url: ", databaseUrl)
	log.Info("flyway url: ", flywayURL)
	err = postgresResource.Expire(180) // Tell docker to hard kill the container in 120 seconds
	if err != nil {
		log.Errorf("Could not start resource: %s", err)
	}
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 60 * time.Second
	if err = pool.Retry(func() error {
		postgresClient, err = pgxpool.Connect(context.Background(), databaseUrl)
		if err != nil {
			return err
		}
		return postgresClient.Ping(context.Background())
	}); err != nil {
		log.Errorf("Could not connect to docker: %s", err)
	}

	cmd := exec.Command("flyway", "-user=test", "-password=1234",
		"-locations=filesystem:./../../migrations", fmt.Sprintf("-url=%s", flywayURL), "migrate")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		postgresClient.Close()
		log.Errorf("flyway execution error %v", err)
	}
}

func startMongo() {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pull mongodb docker image for version 5.0
	mongoResource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "latest",
		Env: []string{
			// username and password for mongodb superuser
			"MONGO_INITDB_ROOT_USERNAME=root",
			"MONGO_INITDB_ROOT_PASSWORD=password",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	err = pool.Retry(func() error {
		var err error
		mongoClient, err = mongo.Connect(
			context.TODO(),
			options.Client().ApplyURI(
				fmt.Sprintf("mongodb://root:password@localhost:%s", mongoResource.GetPort("27017/tcp")),
			),
		)
		if err != nil {
			return err
		}
		return mongoClient.Ping(context.TODO(), nil)
	})
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	err = mongoResource.Expire(180) // Tell docker to hard kill the container in 120 seconds
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
}

func startRedis() {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Errorf("Could not connect to docker: %s", err)
	}

	resource, err := pool.Run("redis", "latest", nil)
	if err != nil {
		log.Errorf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		redisClient = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("localhost:%s", resource.GetPort("6379/tcp")),
		})

		return redisClient.Ping().Err()
	}); err != nil {
		log.Errorf("Could not connect to docker: %s", err)
	}
	err = resource.Expire(180)
	if err != nil {
		log.Errorf("Could not start resource: %s", err)
	}
}
