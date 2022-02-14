package dockertest

import (
	"context"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

var dbClient *mongo.Client

func TestUserCreate(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewMongoUserRepository(dbClient)
	expected := model.User{Login: "create", Password: "1234"}
	err := repo.Create(ctx, expected)
	require.NoError(t, err)
}

func TestUserGet(t *testing.T) {
	ctx := context.Background()
	collection := dbClient.Database("myDatabase").Collection("users")
	repo := repository.NewMongoUserRepository(dbClient)
	expected := model.User{Login: "get", Password: "1234"}
	_, err := collection.InsertOne(ctx, expected)
	require.NoError(t, err)
	actual, err := repo.Get(ctx, expected.Login)
	require.NoError(t, err)
	require.Equal(t, expected.Password, actual.Password)
}
