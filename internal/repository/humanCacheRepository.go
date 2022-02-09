package repository

import (
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/config"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/go-redis/redis"
	"strconv"
)

type HumanCacheRepository interface {
	Create(h model.Human) error
	Get(name string) (*model.Human, error)
	Delete(name string) error
}

// RedisHumanCacheRepository implements HumanCacheRepository with redis
type RedisHumanCacheRepository struct {
	client *redis.Client
}

// NewRedisHumanCacheRepository returns new instance of RedisHumanCacheRepository
func NewRedisHumanCacheRepository(c *redis.Client) *RedisHumanCacheRepository {
	return &RedisHumanCacheRepository{client: c}
}

// Create is used for creating human cache info in db
func (r *RedisHumanCacheRepository) Create(h model.Human) error {
	err := r.client.XAdd(&redis.XAddArgs{
		Stream:       config.RedisStream,
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "",
		Values: map[string]interface{}{
			"id":   h.ID,
			"name": h.Name,
			"male": h.Male,
			"age":  h.Age,
		},
	}).Err()
	if err != nil {
		return fmt.Errorf("redis stream create human cache error %w", err)
	}
	return nil
}

// Get is used for getting human cache info from database
func (r *RedisHumanCacheRepository) Get(name string) (*model.Human, error) {
	res, err := r.client.XRead(&redis.XReadArgs{
		Block:   0,
		Count:   0,
		Streams: []string{config.RedisStream, "0"},
	}).Result()
	if err != nil {
		return nil, fmt.Errorf("redis stream get human cache error %w", err)
	}
	var humanMap map[string]interface{}
	var found bool
	for i := 0; i < len(res[0].Messages); i++ {
		if res[0].Messages[i].Values["name"].(string) == name {
			humanMap = res[0].Messages[i].Values
			found = true
			break
		}
	}
	if !found {
		return nil, fmt.Errorf("human info wasn't found in redis db")
	}
	var male bool
	if humanMap["male"].(string) == "1" {
		male = true
	} else {
		male = false
	}
	age, err := strconv.Atoi(humanMap["age"].(string))
	if err != nil {
		return nil, fmt.Errorf("converting age to int error in redis get %w", err)
	}
	h := model.Human{
		ID:   humanMap["id"].(string),
		Name: humanMap["name"].(string),
		Male: male,
		Age:  age,
	}
	return &h, nil
}

// Delete is used for deleting human cache info from db
func (r *RedisHumanCacheRepository) Delete(name string) error {
	res, err := r.client.XRead(&redis.XReadArgs{
		Block:   0,
		Count:   0,
		Streams: []string{"people", "0"},
	}).Result()
	if err != nil {
		return fmt.Errorf("redis stream get human cache error in delete %w", err)
	}

	for i := 0; i < len(res[0].Messages); i++ {
		if res[0].Messages[i].Values["name"].(string) == name {
			err := r.client.XDel(res[0].Stream, res[0].Messages[i].ID).Err()
			if err != nil {
				return fmt.Errorf("redis stream delete human cache error %w", err)
			}
			return nil
		}
	}
	return fmt.Errorf("human info wasn't found in redis db")
}
