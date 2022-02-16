package repository

import (
	"context"
	"testing"

	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbClient *mongo.Client

func TestUserCreate(t *testing.T) {
	ctx := context.Background()
	collection := dbClient.Database("myDatabase").Collection("users")
	repo := NewMongoUserRepository(dbClient)
	expected := model.User{Login: "create", Password: "1234"}
	err := repo.Create(ctx, expected)
	require.NoError(t, err)
	var actual model.User
	err = collection.FindOne(ctx, bson.D{primitive.E{Key: "login", Value: expected.Login}}).Decode(&actual)
	require.NoError(t, err)
	require.Equal(t, expected.Password, actual.Password)
	err = collection.Drop(ctx)
	require.NoError(t, err)
}

func TestUserGet(t *testing.T) {
	ctx := context.Background()
	collection := dbClient.Database("myDatabase").Collection("users")
	repo := NewMongoUserRepository(dbClient)
	expected := model.User{Login: "get", Password: "1234"}
	_, err := collection.InsertOne(ctx, expected)
	require.NoError(t, err)
	actual, err := repo.Get(ctx, expected.Login)
	require.NoError(t, err)
	require.Equal(t, expected.Password, actual.Password)
	err = collection.Drop(ctx)
	require.NoError(t, err)
}

func TestUserUpdate(t *testing.T) {
	ctx := context.Background()
	collection := dbClient.Database("myDatabase").Collection("users")
	repo := NewMongoUserRepository(dbClient)
	expected := model.User{Login: "update", Password: "12345"}
	_, err := collection.InsertOne(ctx, expected)
	require.NoError(t, err)
	expected.Login = "updated"
	err = repo.Update(ctx, "update", expected)
	require.NoError(t, err)
	var actual model.User
	err = collection.FindOne(ctx, bson.D{primitive.E{Key: "login", Value: expected.Login}}).Decode(&actual)
	require.NoError(t, err)
	require.Equal(t, expected.Password, actual.Password)
	require.Equal(t, expected.Role, actual.Role)
	err = collection.Drop(ctx)
	require.NoError(t, err)
}

func TestUserDelete(t *testing.T) {
	ctx := context.Background()
	collection := dbClient.Database("myDatabase").Collection("users")
	repo := NewMongoUserRepository(dbClient)
	expected := model.User{Login: "delete", Password: "12345"}
	_, err := collection.InsertOne(ctx, expected)
	require.NoError(t, err)
	err = repo.Delete(ctx, expected.Login)
	require.NoError(t, err)
	var deleted model.User
	err = collection.FindOne(ctx, bson.D{primitive.E{Key: "login", Value: expected.Login}}).Decode(&deleted)
	expectedError := collection.FindOne(ctx, bson.D{primitive.E{Key: "login", Value: "not exising"}}).Decode(&deleted)
	require.Equal(t, expectedError, err)
	err = collection.Drop(ctx)
	require.NoError(t, err)
}
