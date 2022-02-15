package dockertest

import (
	"context"
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTokenCreate(t *testing.T) {
	ctx := context.Background()
	collection := dbClient.Database("myDatabase").Collection("tokens")
	repo := repository.NewMongoTokenRepository(dbClient)
	expected := model.Token{Login: "create", Value: "wwefgsdgoj", ExpiresAt: 123123}

	err := repo.Create(ctx, expected)
	require.NoError(t, err)
	var actual model.Token
	err = collection.FindOne(ctx, bson.D{primitive.E{Key: "login", Value: expected.Login}}).Decode(&actual)
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("%x", sha256.Sum256([]byte(expected.Value))), actual.Value)
	require.Equal(t, expected.ExpiresAt, actual.ExpiresAt)
	err = collection.Drop(ctx)
	require.NoError(t, err)
}

func TestTokenGet(t *testing.T) {
	ctx := context.Background()
	collection := dbClient.Database("myDatabase").Collection("tokens")
	repo := repository.NewMongoTokenRepository(dbClient)
	expectedValue := "aufgbafg"
	expected := model.Token{Login: "get",
		Value: fmt.Sprintf("%x", sha256.Sum256([]byte(expectedValue))), ExpiresAt: 123123}
	_, err := collection.InsertOne(ctx, expected)
	require.NoError(t, err)
	actual, err := repo.Get(ctx, expectedValue)
	require.NoError(t, err)
	require.Equal(t, expected.Login, actual.Login)
	require.Equal(t, expected.ExpiresAt, actual.ExpiresAt)
	err = collection.Drop(ctx)
	require.NoError(t, err)
}

func TestTokenDelete(t *testing.T) {
	ctx := context.Background()
	collection := dbClient.Database("myDatabase").Collection("tokens")
	repo := repository.NewMongoTokenRepository(dbClient)
	expectedValue := "aufgbafg"
	expected := model.Token{Login: "delete",
		Value: fmt.Sprintf("%x", sha256.Sum256([]byte(expectedValue))), ExpiresAt: 123123}
	_, err := collection.InsertOne(ctx, expected)
	require.NoError(t, err)
	err = repo.Delete(ctx, expectedValue)
	require.NoError(t, err)
	var deleted model.Token
	err = collection.FindOne(ctx, bson.D{primitive.E{Key: "login", Value: expected.Login}}).Decode(&deleted)
	expectedError := collection.FindOne(ctx, bson.D{primitive.E{Key: "login", Value: "not exising"}}).Decode(&deleted)
	require.Equal(t, expectedError, err)
	err = collection.Drop(ctx)
	require.NoError(t, err)
}
