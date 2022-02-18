package handlers

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/config"
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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"os"
	"os/exec"
	"testing"
	"time"
)

const (
	url = "http://localhost:1323/"
)

var (
	postgresClient *pgxpool.Pool
	mongoClient    *mongo.Client
	redisClient    *redis.Client

	accessToken  string
	refreshToken string
)

func TestMain(m *testing.M) {
	postgresPool, postgresResource := startPostgres()
	mongoPool, mongoResource := startMongo()
	redisPool, redisResource := startRedis()

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
	humanHandler := NewHumanHandler(service.NewHumanService(repository.NewHumanRepository(postgresClient), redisCacheHumanRepository),
		service.NewAuthService(repository.NewMongoTokenRepository(mongoClient)))
	userHandler := NewUserHandler(service.NewUserService(userRepo),
		service.NewAuthService(repository.NewMongoTokenRepository(mongoClient)))
	fileHandler := &FileHandler{}
	echoServer := echo.New()
	validator, err := validation.NewValidator()
	if err != nil {
		log.Warnf("echo validator error %v", err)
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
	go func() {
		err = echoServer.Start(":1323")
		if err != nil {
			log.Warnf("error with starting an echo server: %v", err)
			return
		}
	}()
	time.Sleep(time.Second)
	code := m.Run()
	if err := redisPool.Purge(redisResource); err != nil {
		log.Errorf("Could not purge resource: %s", err)
	}
	// You can't defer this because os.Exit doesn't care for defer
	if err := postgresPool.Purge(postgresResource); err != nil {
		log.Errorf("Could not purge resource: %s", err)
	}

	// disconnect mongodb client
	if err := mongoPool.Purge(mongoResource); err != nil {
		log.Errorf("Could not purge resource: %s", err)
	}

	// disconnect mongodb client
	if err := mongoClient.Disconnect(context.TODO()); err != nil {
		log.Errorf("mongo disconnection error %v", err)
	}
	os.Exit(code)

}

func startPostgres() (*dockertest.Pool, *dockertest.Resource) {
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
	err = postgresResource.Expire(60) // Tell docker to hard kill the container in 120 seconds
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
	return pool, postgresResource
}

func startMongo() (*dockertest.Pool, *dockertest.Resource) {
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
	err = mongoResource.Expire(60) // Tell docker to hard kill the container in 120 seconds
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	return pool, mongoResource
}

func startRedis() (*dockertest.Pool, *dockertest.Resource) {
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
	err = resource.Expire(60) // Tell docker to hard kill the container in 120 seconds
	if err != nil {
		log.Errorf("Could not start resource: %s", err)
	}
	return pool, resource
}
