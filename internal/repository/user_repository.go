package repository

import (
	"context"
	"fmt"

	"github.com/HekapOo-hub/Task1/internal/model"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository is a crud interface for working with db where user info stored
type UserRepository interface {
	Create(ctx context.Context, user model.User) error
	Get(ctx context.Context, login string) (*model.User, error)
	Delete(ctx context.Context, login string) error
	Update(ctx context.Context, login string, newUser model.User) error
}

// MongoUserRepository implement UserRepository interface
type MongoUserRepository struct {
	collection *mongo.Collection
}

// NewMongoUserRepository returns instance of MongoUserRepository
func NewMongoUserRepository(c *mongo.Client) *MongoUserRepository {
	collection := c.Database("myDatabase").Collection("users")
	return &MongoUserRepository{collection: collection}
}

// Create is used for creating human info in db
func (m *MongoUserRepository) Create(ctx context.Context, user model.User) error {
	user.ID = uuid.NewV4().String()
	user.Role = "admin"
	_, err := m.collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("mongo creation user error %w", err)
	}
	return nil
}

// Get is used for getting user info from db
func (m *MongoUserRepository) Get(ctx context.Context, login string) (*model.User, error) {
	var res model.User
	filter := bson.D{primitive.E{Key: "login", Value: login}}
	err := m.collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return nil, fmt.Errorf("mongo get user error %w", err)
	}
	return &res, nil
}

// Delete is used for deleting user info from db
func (m *MongoUserRepository) Delete(ctx context.Context, login string) error {
	filter := bson.D{primitive.E{Key: "login", Value: login}}
	delRes, err := m.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("mongo delete user error %w", err)
	}
	if delRes.DeletedCount == 0 {
		return fmt.Errorf("no such user in db")
	}
	return nil
}

// Update is used for updating user info in db
func (m *MongoUserRepository) Update(ctx context.Context, login string, newUser model.User) error {
	filter := bson.D{primitive.E{Key: "login", Value: login}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "login", Value: newUser.Login}, primitive.E{Key: "password", Value: newUser.Password},
		},
		}}
	_, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("mongo update user error %w", err)
	}
	return nil
}
