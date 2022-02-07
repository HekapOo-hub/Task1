// Package repository contains structures which implement crud interface on different databases and tables
package repository

import (
	"context"
	"fmt"

	"github.com/HekapOo-hub/Task1/internal/model"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoRepository is a struct for working with mongoDB
type MongoRepository struct {
	collection *mongo.Collection
}

// MongoDisconnect is a function to close connection with mongoDB
func MongoDisconnect(ctx context.Context, m *mongo.Client) {
	if err := m.Disconnect(ctx); err != nil {
		log.WithField("error", err).Errorf("mongo disconnect error")
	}
}

// NewMongoRepository creates new mongo repository with human collection in it
func NewMongoRepository(c *mongo.Client) *MongoRepository {
	collection := c.Database("myDatabase").Collection("human")
	return &MongoRepository{collection: collection}
}

// Create is used for creating human info in db
func (m *MongoRepository) Create(ctx context.Context, h model.Human) error {
	h.ID = uuid.NewV4().String()
	_, err := m.collection.InsertOne(ctx, h)
	if err != nil {
		return fmt.Errorf("mongo creation human error %w", err)
	}
	return nil
}

// Get is used for getting human info in db
func (m *MongoRepository) Get(ctx context.Context, name string) (*model.Human, error) {
	var res model.Human
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	err := m.collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return nil, fmt.Errorf("mongo get human error %w", err)
	}
	return &res, nil
}

// Update is used for updating human info in db
func (m *MongoRepository) Update(ctx context.Context, id string, h model.Human) error {
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "name", Value: h.Name}, primitive.E{Key: "male", Value: h.Male},
			primitive.E{Key: "age", Value: h.Age},
		},
		}}
	_, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("mongo update human error %w", err)
	}
	return nil
}

// Delete is used for deleting human info from db
func (m *MongoRepository) Delete(ctx context.Context, id string) error {
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	_, err := m.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("mongo delete human error %w", err)
	}
	return nil
}
