package repository

import (
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/config"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
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
	cache  map[string]model.Human
	lastID string
	cancel chan bool
}

// NewRedisHumanCacheRepository returns new instance of RedisHumanCacheRepository
func NewRedisHumanCacheRepository(c *redis.Client) *RedisHumanCacheRepository {
	r := &RedisHumanCacheRepository{
		client: c,
		cache:  make(map[string]model.Human),
		lastID: "0-0",
		cancel: make(chan bool),
	}
	go r.listen()
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
	log.Warn("sending to chan!")
	r.cancel <- false
	return nil
}

// Get is used for getting human cache info
func (r *RedisHumanCacheRepository) Get(name string) (*model.Human, error) {
	h, ok := r.cache[name]
	if !ok {
		return nil, fmt.Errorf("human not found in cache")
	}
	return &h, nil
}

// Delete is used for deleting human cache info
func (r *RedisHumanCacheRepository) Delete(name string) error {
	delete(r.cache, name)
	return nil
}

func (r *RedisHumanCacheRepository) listen() {
	for {
		c := <-r.cancel
		switch c {
		case true:
			log.Warn("go routing finished!")
			return
		case false:
			res, err := r.client.XRead(&redis.XReadArgs{
				Block:   0,
				Count:   1,
				Streams: []string{config.RedisStream, r.lastID},
			}).Result()
			r.lastID = res[0].Messages[0].ID
			if err != nil {
				log.WithField("error", err).Warn("read stream error")
				return
			}
			humanMap := res[0].Messages[0].Values

			var male bool
			if humanMap["male"].(string) == "1" {
				male = true
			} else {
				male = false
			}
			age, err := strconv.Atoi(humanMap["age"].(string))
			if err != nil {
				log.WithField("error", err).Warn("converting age to int error in redis listen")
			}
			r.cache[humanMap["name"].(string)] = model.Human{
				Name: humanMap["name"].(string),
				ID:   humanMap["id"].(string),
				Male: male,
				Age:  age,
			}
		}
	}
}

// Cancel is used for finishing listen go routine
func (r *RedisHumanCacheRepository) Cancel() {
	r.cancel <- true
}
