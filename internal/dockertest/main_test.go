package dockertest

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	postgresPool, postgresResource := startPostgres()
	mongoPool, mongoResource := startMongo()
	//Run tests
	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := postgresPool.Purge(postgresResource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	// disconnect mongodb client
	if err := mongoPool.Purge(mongoResource); err != nil {
		log.Errorf("Could not purge resource: %s", err)
	}

	// disconnect mongodb client
	if err := dbClient.Disconnect(context.TODO()); err != nil {
		log.Errorf("mongo disconnection error %v", err)
	}
	os.Exit(code)
}
func startPostgres() (*dockertest.Pool, *dockertest.Resource) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	postgresResource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
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
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := postgresResource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgresql://test:1234@%s/testDB?sslmode=disable", hostAndPort)
	flywayURL := fmt.Sprintf("jdbc:postgresql://%s/testDB", hostAndPort)
	log.Info("Connecting to database on url: ", databaseUrl)
	log.Info("flyway url: ", flywayURL)
	/*	err = postgresResource.Expire(120) // Tell docker to hard kill the container in 120 seconds
		if err != nil {
			log.Fatalf("Could not start resource: %s", err)
		}*/
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 60 * time.Second
	if err = pool.Retry(func() error {
		postgresDB, err = pgxpool.Connect(context.Background(), databaseUrl)
		if err != nil {
			return err
		}
		return postgresDB.Ping(context.Background())
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	cmd := exec.Command("flyway", "-user=test", "-password=1234",
		"-locations=filesystem:./../../migrations", fmt.Sprintf("-url=%s", flywayURL), "migrate")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		postgresDB.Close()
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
		Tag:        "5.0",
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
		dbClient, err = mongo.Connect(
			context.TODO(),
			options.Client().ApplyURI(
				fmt.Sprintf("mongodb://root:password@localhost:%s", mongoResource.GetPort("27017/tcp")),
			),
		)
		if err != nil {
			return err
		}
		return dbClient.Ping(context.TODO(), nil)
	})
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	err = mongoResource.Expire(120) // Tell docker to hard kill the container in 120 seconds
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	return pool, mongoResource
}
