package repository

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/HekapOo-hub/Task1/internal/config"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

// HumanCacheRepository is model for working with human cache
type HumanCacheRepository interface {
	Create(h model.Human) error
	Get(name string) (*model.Human, error)
	Delete(name string) error
}

// RedisHumanCacheRepository implements HumanCacheRepository with redis
type RedisHumanCacheRepository struct {
	client *redis.Client
	cache  map[string]model.Human
	lastID string
	mu     sync.RWMutex
}

// NewRedisHumanCacheRepository returns new instance of RedisHumanCacheRepository
func NewRedisHumanCacheRepository(ctx context.Context, c *redis.Client) *RedisHumanCacheRepository {
	r := &RedisHumanCacheRepository{
		client: c,
		cache:  make(map[string]model.Human),
		lastID: "0-0",
		mu:     sync.RWMutex{},
	}
	go r.listen(ctx)
	return r
}

// Create is used for sending create request to stream
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
		return fmt.Errorf("redis stream sending create human request error %w", err)
	}
	return nil
}

// Get is used for getting human cache info
func (r *RedisHumanCacheRepository) Get(name string) (*model.Human, error) {
	r.mu.RLock()
	h, ok := r.cache[name]
	r.mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("human not found in cache")
	}
	return &h, nil
}

// Delete is used for deleting human cache info
func (r *RedisHumanCacheRepository) Delete(name string) error {
	r.mu.Lock()
	delete(r.cache, name)
	r.mu.Unlock()
	return nil
}

func (r *RedisHumanCacheRepository) listen(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			res, err := r.client.XRead(&redis.XReadArgs{
				Block:   0,
				Count:   1,
				Streams: []string{config.RedisStream, r.lastID},
			}).Result()
			if err != nil {
				log.Warnf("read stream error:%v", err)
				continue
			}
			if res[0].Messages == nil {
				log.Warn("message is empty")
				continue
			}
			r.lastID = res[0].Messages[0].ID
			humanMap := res[0].Messages[0].Values
			var male bool
			maleStr := humanMap["male"]
			switch maleStr {
			case "1":
				male = true
			case "0":
				male = false
			default:
				log.Warnf("parsing bool male from stream message in listen error: %v", err)
				continue
			}
			age, err := strconv.Atoi(humanMap["age"].(string))
			if err != nil {
				log.Warnf("converting age to int error in redis listen: %v", err)
				continue
			}
			r.mu.Lock()
			r.cache[humanMap["name"].(string)] = model.Human{
				Name: humanMap["name"].(string),
				ID:   humanMap["id"].(string),
				Male: male,
				Age:  age,
			}
			r.mu.Unlock()
		}
	}
}
