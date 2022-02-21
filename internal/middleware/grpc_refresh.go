package middleware

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/config"
	"github.com/HekapOo-hub/Task1/internal/repository"
	"github.com/HekapOo-hub/Task1/internal/service"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"time"
)

var (
	needToRefresh = time.Minute * 10
	mongoClient   *mongo.Client
)

func unaryServerRefreshInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	if mongoClient == nil {
		uri, err := config.GetMongoURI()
		if err != nil {
			log.Warnf("error: %v", err)
			return nil, fmt.Errorf("unary server refresh interceptor error %v", err)
		}
		mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
		if err != nil {
			log.WithField("uri", uri).Warnf("error with connecting to mongodb: %v", err)
			return nil, fmt.Errorf("unary server refresh inerceptor error %v", err)
		}

	}
	claims, err := authFunc(ctx)
	if err != nil {
		return nil, fmt.Errorf("unary server refresh interceptor %v", err)
	}
	expireAt := claims.StandardClaims.ExpiresAt
	beforeExpiration := time.Unix(expireAt, 0).Sub(time.Now())
	if beforeExpiration.Minutes() < needToRefresh.Minutes() && beforeExpiration.Minutes() > 0 {
		tokenService := service.NewAuthService(repository.NewMongoTokenRepository(mongoClient))
		access, refresh, err := tokenService.Refresh(ctx, claims)
	}
}
