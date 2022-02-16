package service

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	humanCacheRepository mocks.HumanCacheRepository
	humanRepository      mocks.HumanRepository
)

func TestHumanService_Create(t *testing.T) {
	ctx := context.Background()
	human := model.Human{Name: "its me", Age: 228, Male: true}
	humanService := NewHumanService(&humanRepository, &humanCacheRepository)
	humanRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("model.Human")).Return(func(ctx context.Context, human model.Human) error {
		return nil
	})
	humanCacheRepository.On("Create", mock.AnythingOfType("model.Human")).Return(
		func(human model.Human) error {
			return nil
		})
	err := humanService.Create(ctx, human)
	require.NoError(t, err)
}

func TestHumanService_Get(t *testing.T) {
	ctx := context.Background()
	name := "qwerty"
	humanService := NewHumanService(&humanRepository, &humanCacheRepository)
	humanCacheRepository.On("Get", mock.AnythingOfType("string")).Return(
		func(name string) *model.Human {
			return &model.Human{Name: name}
		},
		func(name string) error {
			return nil
		})
	_, err := humanService.Get(ctx, name)
	require.NoError(t, err)
	humanCacheRepository.On("Get", mock.AnythingOfType("string")).Return(
		func(ctx context.Context, name string) *model.Human {
			return &model.Human{Name: name}
		},
		func(ctx context.Context, s string) error {
			return fmt.Errorf("cache error")
		})

	humanRepository.On("Get", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(func(ctx context.Context, name string) *model.Human {
		return &model.Human{Name: name}
	},
		func(ctx context.Context, s string) error {
			return nil
		})
	humanCacheRepository.On("Create", mock.AnythingOfType("model.Human")).Return(
		func(human model.Human) error {
			return nil
		})
	_, err = humanService.Get(ctx, name)
	require.NoError(t, err)
}

func TestHumanService_Update(t *testing.T) {
	ctx := context.Background()
	humanService := NewHumanService(&humanRepository, &humanCacheRepository)
	name := "old"
	human := model.Human{Name: "new", Age: 4322, Male: false}
	humanRepository.On("Update", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string"),
		mock.AnythingOfType("model.Human")).Return(func(ctx context.Context, name string, human model.Human) error {
		return nil
	})
	humanCacheRepository.On("Delete", mock.AnythingOfType("string")).Return(func(name string) error {
		return nil
	})
	err := humanService.Update(ctx, name, human)
	require.NoError(t, err)
}

func TestHumanService_Delete(t *testing.T) {
	ctx := context.Background()
	humanService := NewHumanService(&humanRepository, &humanCacheRepository)
	name := "deleted"
	humanCacheRepository.On("Delete", mock.AnythingOfType("string")).Return(func(name string) error {
		return nil
	})
	humanRepository.On("Delete", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(func(ctx context.Context, name string) error {
		return nil
	})
	err := humanService.Delete(ctx, name)
	require.NoError(t, err)
}
