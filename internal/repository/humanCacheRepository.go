package repository

import (
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/go-redis/redis"
)

type HumanCacheRepository interface {
	Create() error
	Get() (*model.Human, error)
}

// RedisHumanCacheRepository implements HumanCacheRepository with redis
type RedisHumanCacheRepository struct {
	client *redis.Client
}

// NewRedisHumanCacheRepository returns new instance of RedisHumanCacheRepository
func NewRedisHumanCacheRepository(c *redis.Client) *RedisHumanCacheRepository {
	return &RedisHumanCacheRepository{client: c}
}

// Create is used for creating human ache info in db
//func (r *RedisHumanCacheRepository) Create()
