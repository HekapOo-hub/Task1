package repository

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenRepository interface {
	Create(ctx context.Context, token model.Token) error
	Get(ctx context.Context, token string) (int64, error)
	Delete(ctx context.Context, token string) error
}

type MongoTokenRepository struct {
	collection *mongo.Collection
}

func NewMongoTokenRepository(c *mongo.Client) *MongoTokenRepository {
	collection := c.Database("myDatabase").Collection("tokens")
	return &MongoTokenRepository{collection: collection}
}

func (m *MongoTokenRepository) Create(ctx context.Context, token model.Token) error {
	_, err := m.collection.InsertOne(ctx, token)
	if err != nil {
		return fmt.Errorf("mongo create token error %w", err)
	}
	return nil
}

func (m *MongoTokenRepository) Get(ctx context.Context, id string) (int64, error) {
	var res model.Token
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	err := m.collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return -1, fmt.Errorf("mongo get token error %w", err)
	}
	return res.ExpiresAt, nil
}

func (m *MongoTokenRepository) Delete(ctx context.Context, token string) error {
	filter := bson.D{primitive.E{Key: "token", Value: token}}
	_, err := m.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("mongo delete token error %w", err)
	}
	return nil
}
