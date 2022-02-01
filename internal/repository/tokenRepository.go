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
	Get(ctx context.Context, login string) ([]model.Token, error)
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

func (m *MongoTokenRepository) Get(ctx context.Context, login string) ([]model.Token, error) {
	var res []model.Token
	filter := bson.D{primitive.E{Key: "login", Value: login}}
	cur, err := m.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("mongo get token error %w", err)
	}
	for cur.Next(ctx) {
		var token model.Token
		err := cur.Decode(&token)
		if err != nil {
			return nil, fmt.Errorf("cursor decode error in get %w", err)
		}
		res = append(res, token)
	}
	return res, nil
}

func (m *MongoTokenRepository) Delete(ctx context.Context, token string) error {
	filter := bson.D{primitive.E{Key: "value", Value: token}}
	_, err := m.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("mongo delete token error %w", err)
	}
	return nil
}
