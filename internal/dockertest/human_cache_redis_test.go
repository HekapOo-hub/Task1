package dockertest

import (
	"context"
	"testing"
	"time"

	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

var redisClient *redis.Client

func TestHumanCache(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	repo := repository.NewRedisHumanCacheRepository(ctx, redisClient)
	expected := model.Human{ID: uuid.NewV4().String(), Name: "create",
		Male: false, Age: 123}
	err := repo.Create(expected)
	require.NoError(t, err)
	time.Sleep(time.Second * 1)
	actual, err := repo.Get(expected.Name)
	require.NoError(t, err)
	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, expected.Age, actual.Age)
	require.Equal(t, expected.Male, actual.Male)
	err = repo.Delete(expected.Name)
	require.NoError(t, err)
	_, err = repo.Get(expected.Name)
	require.Error(t, err)
}
