package repository

import (
	"context"
	"testing"

	"github.com/HekapOo-hub/Task1/internal/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestHumanMongoCreate(t *testing.T) {
	ctx := context.Background()
	collection := dbClient.Database("myDatabase").Collection("human")
	repo := NewHumanMongoRepository(dbClient)
	expected := model.Human{Name: "create", Age: 11, Male: false}
	expected.ID = uuid.NewV1().String()
	err := repo.Create(ctx, expected)
	require.NoError(t, err)
	var actual model.Human
	err = collection.FindOne(ctx,
		bson.D{primitive.E{Key: "name", Value: expected.Name}}).Decode(&actual)
	require.NoError(t, err)
	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, expected.Male, actual.Male)
	require.Equal(t, expected.Age, actual.Age)
	err = collection.Drop(ctx)
	require.NoError(t, err)
}

func TestHumanMongoGet(t *testing.T) {
	ctx := context.Background()
	collection := dbClient.Database("myDatabase").Collection("human")
	repo := NewHumanMongoRepository(dbClient)
	expected := model.Human{Name: "get", Age: 11, Male: false}
	expected.ID = uuid.NewV1().String()
	_, err := collection.InsertOne(ctx, expected)
	require.NoError(t, err)
	actual, err := repo.Get(ctx, expected.Name)
	require.NoError(t, err)
	require.Equal(t, expected.Age, actual.Age)
	require.Equal(t, expected.Male, actual.Male)
	err = collection.Drop(ctx)
	require.NoError(t, err)
}

func TestHumanMongoUpdate(t *testing.T) {
	ctx := context.Background()
	collection := dbClient.Database("myDatabase").Collection("human")
	repo := NewHumanMongoRepository(dbClient)
	expected := model.Human{Name: "update", Age: 11, Male: false}
	_, err := collection.InsertOne(ctx, expected)
	require.NoError(t, err)
	expected.Name = "updated"
	expected.Age = 123
	err = repo.Update(ctx, "update", expected)
	require.NoError(t, err)
	var actual model.Human
	err = collection.FindOne(ctx,
		bson.D{primitive.E{Key: "name", Value: expected.Name}}).Decode(&actual)
	require.NoError(t, err)
	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, expected.Male, actual.Male)
	require.Equal(t, expected.Age, actual.Age)
	err = collection.Drop(ctx)
	require.NoError(t, err)
}

func TestHumanMongoDelete(t *testing.T) {
	ctx := context.Background()
	collection := dbClient.Database("myDatabase").Collection("human")
	repo := NewHumanMongoRepository(dbClient)
	expected := model.Human{Name: "delete", Age: 11, Male: false}
	_, err := collection.InsertOne(ctx, expected)
	require.NoError(t, err)
	err = repo.Delete(ctx, expected.Name)
	require.NoError(t, err)
	var deleted model.Human
	err = collection.FindOne(ctx, bson.D{primitive.E{Key: "name", Value: expected.Name}}).Decode(&deleted)
	expectedError := collection.FindOne(ctx, bson.D{primitive.E{Key: "name", Value: "not exising"}}).Decode(&deleted)
	require.Equal(t, expectedError, err)
}
