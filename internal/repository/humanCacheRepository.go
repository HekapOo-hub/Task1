package repository

import (
	"encoding/json"
	"fmt"

	"github.com/HekapOo-hub/Task1/internal/config"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/go-redis/redis"
)

type HumanCacheRepository interface {
	Create(h model.Human) error
	Get(name string) (*model.Human, error)
	Update(name string, h model.Human) error
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

	jsonHuman, err := json.Marshal(h)
	if err != nil {
		return fmt.Errorf("redis create json marshal error %w", err)
	}
	if err := r.client.Set(h.Name, jsonHuman, config.HumanCacheTTL).Err(); err != nil {
		return fmt.Errorf("redis create human cache error %w", err)
	}
	return nil
}

// Get is used for getting human cache info from database
func (r *RedisHumanCacheRepository) Get(name string) (*model.Human, error) {
	res, err := r.client.Get(name).Result()
	if err != nil {
		return nil, fmt.Errorf("redis get human cache info error %w", err)
	}
	var human model.Human
	err = json.Unmarshal([]byte(res), &human)
	if err != nil {
		return nil, fmt.Errorf("redis get human json unmarshal error %w", err)
	}
	return &human, nil
}

// Update is used for updating human cache info in db
func (r *RedisHumanCacheRepository) Update(name string, h model.Human) error {
	res, err := r.client.Get(name).Result()
	if err != nil {
		return fmt.Errorf("redis get human cache info error in update %w", err)
	}
	var human model.Human
	err = json.Unmarshal([]byte(res), &human)
	if err != nil {
		return fmt.Errorf("redis get human json unmarshal error in update %w", err)
	}
	h.ID = human.ID
	if err := r.client.Del(name).Err(); err != nil {
		return fmt.Errorf("redis delete human cache info error in update %w", err)
	}
	jsonHuman, err := json.Marshal(h)
	if err != nil {
		return fmt.Errorf("redis update human json marshal error %w", err)
	}
	if err := r.client.Set(h.Name, jsonHuman, config.HumanCacheTTL).Err(); err != nil {
		return fmt.Errorf("redis update human cache info error %w", err)
	}
	return nil
}

// Delete is used for deleting human cache info from db
func (r *RedisHumanCacheRepository) Delete(name string) error {
	if err := r.client.Del(name).Err(); err != nil {
		return fmt.Errorf("redis delete human cache info error %w", err)
	}
	return nil
}
