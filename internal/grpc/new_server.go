package grpc

import (
	"context"
	"github.com/HekapOo-hub/Task1/internal/config"
	"github.com/HekapOo-hub/Task1/internal/repository"
	"github.com/HekapOo-hub/Task1/internal/service"
	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetInstance() (server *HumanServer, cancel context.CancelFunc, err error) {
	cfg, err := config.NewPostgresConfig()
	if err != nil {
		log.Warnf("postgres config error: %v", err)
		return nil, nil, err
	}
	postgresClient, err := pgxpool.Connect(context.Background(), cfg.GetURL())
	if err != nil {
		log.Warnf("postgres connect error: %v", err)
		return nil, nil, err
	}

	uri, err := config.GetMongoURI()
	if err != nil {
		log.Warnf("error: %v", err)
		return nil, nil, err
	}
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.WithField("uri", uri).Warnf("error with connecting to mongodb: %v", err)
		return nil, nil, err
	}

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
	return &HumanServer{
		humanService: service.NewHumanService(repository.NewHumanRepository(postgresClient), redisCacheHumanRepository),
		userService:  service.NewUserService(userRepo),
		authService:  service.NewAuthService(repository.NewMongoTokenRepository(mongoClient)),
		fileService:  &service.FileService{},
	}, cancel, nil
}
