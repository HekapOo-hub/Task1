package repository

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/HekapOo-hub/Task1/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TokenRepository is a crud interface for working with db where token info stored
type TokenRepository interface {
	Create(ctx context.Context, token model.Token) error
	Get(ctx context.Context, token string) (*model.Token, error)
	Delete(ctx context.Context, token string) error
}

// MongoTokenRepository implements TokenRepository with mongoDB
type MongoTokenRepository struct {
	collection *mongo.Collection
}

// NewMongoTokenRepository returns new instance of MongoTokenRepository
func NewMongoTokenRepository(c *mongo.Client) *MongoTokenRepository {
	collection := c.Database("myDatabase").Collection("tokens")
	return &MongoTokenRepository{collection: collection}
}

// Create is used for creating token info in db
func (m *MongoTokenRepository) Create(ctx context.Context, token model.Token) error {
	hashedToken := fmt.Sprintf("%x", sha256.Sum256([]byte(token.Value)))
	token.Value = hashedToken
	_, err := m.collection.InsertOne(ctx, token)
	if err != nil {
		return fmt.Errorf("mongo create token error %w", err)
	}
	return nil
}

// Get is used for getting token info from db
func (m *MongoTokenRepository) Get(ctx context.Context, token string) (*model.Token, error) {
	var res model.Token
	hashedToken := fmt.Sprintf("%x", sha256.Sum256([]byte(token)))
	filter := bson.D{primitive.E{Key: "value", Value: hashedToken}}
	err := m.collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return nil, fmt.Errorf("mongo get token error %w", err)
	}
	return &res, nil
}

// Delete is used for deleting token info from db
func (m *MongoTokenRepository) Delete(ctx context.Context, login string) error {
	filter := bson.D{primitive.E{Key: "login", Value: login}}
	delRes, err := m.collection.DeleteMany(ctx, filter)
	if err != nil {
		return fmt.Errorf("mongo delete token error %w", err)
	}
	if delRes.DeletedCount == 0 {
		return fmt.Errorf("no such token in db")
	}
	return nil
}
