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
	humanRepository = mocks.HumanRepository{}
	humanCacheRepository = mocks.HumanCacheRepository{}
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
	humanRepository = mocks.HumanRepository{}
	humanRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("model.Human")).Return(func(ctx context.Context, human model.Human) error {
		return fmt.Errorf("this must be wrapped")
	})
	err = humanService.Create(ctx, human)
	require.Equal(t, fmt.Errorf("human service %w",
		fmt.Errorf("this must be wrapped")), err)
	humanRepository = mocks.HumanRepository{}
	humanCacheRepository = mocks.HumanCacheRepository{}
	humanRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("model.Human")).Return(func(ctx context.Context, human model.Human) error {
		return nil
	})
	humanCacheRepository.On("Create", mock.AnythingOfType("model.Human")).Return(
		func(human model.Human) error {
			return fmt.Errorf("cache must be wrapped")
		})
	err = humanService.Create(ctx, human)
	require.Equal(t, fmt.Errorf("human service %w",
		fmt.Errorf("cache must be wrapped")), err)
}

func TestHumanService_Get(t *testing.T) {
	// no error
	humanRepository = mocks.HumanRepository{}
	humanCacheRepository = mocks.HumanCacheRepository{}
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
	// no error
	humanCacheRepository = mocks.HumanCacheRepository{}
	humanCacheRepository.On("Get", mock.AnythingOfType("string")).Return(
		func(name string) *model.Human {
			return &model.Human{Name: name}
		},
		func(name string) error {
			return fmt.Errorf("cache error")
		})

	humanRepository.On("Get", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(func(ctx context.Context, name string) *model.Human {
		return &model.Human{Name: name}
	},
		func(ctx context.Context, name string) error {
			return nil
		})
	humanCacheRepository.On("Create", mock.AnythingOfType("model.Human")).Return(
		func(human model.Human) error {
			return nil
		})
	_, err = humanService.Get(ctx, name)
	require.NoError(t, err)
	// get error
	humanRepository = mocks.HumanRepository{}
	humanRepository.On("Get", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(func(ctx context.Context, name string) *model.Human {
		return &model.Human{Name: name}
	},
		func(ctx context.Context, name string) error {
			return fmt.Errorf("must be wrapped")
		})
	_, err = humanService.Get(ctx, name)
	require.Equal(t, fmt.Errorf("human service %w", fmt.Errorf("must be wrapped")), err)
	// create error
	humanRepository = mocks.HumanRepository{}
	humanCacheRepository = mocks.HumanCacheRepository{}
	humanCacheRepository.On("Get", mock.AnythingOfType("string")).Return(
		func(name string) *model.Human {
			return &model.Human{Name: name}
		},
		func(name string) error {
			return fmt.Errorf("cache error")
		})
	humanRepository.On("Get", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(func(ctx context.Context, name string) *model.Human {
		return &model.Human{Name: name}
	},
		func(ctx context.Context, name string) error {
			return nil
		})
	humanCacheRepository.On("Create", mock.AnythingOfType("model.Human")).Return(
		func(human model.Human) error {
			return fmt.Errorf("create must be wrapped")
		})
	_, err = humanService.Get(ctx, name)
	require.Equal(t, fmt.Errorf("human service get func %w",
		fmt.Errorf("create must be wrapped")), err)
}

func TestHumanService_Update(t *testing.T) {
	// no error
	humanRepository = mocks.HumanRepository{}
	humanCacheRepository = mocks.HumanCacheRepository{}
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
	// update error
	humanRepository = mocks.HumanRepository{}
	humanRepository.On("Update", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string"),
		mock.AnythingOfType("model.Human")).Return(func(ctx context.Context, name string, human model.Human) error {
		return fmt.Errorf("update must be wrapped")
	})
	err = humanService.Update(ctx, name, human)
	require.Equal(t, fmt.Errorf("human service %w", fmt.Errorf("update must be wrapped")), err)
	// delete error
	humanRepository = mocks.HumanRepository{}
	humanCacheRepository = mocks.HumanCacheRepository{}
	humanRepository.On("Update", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string"),
		mock.AnythingOfType("model.Human")).Return(func(ctx context.Context, name string, human model.Human) error {
		return nil
	})
	humanCacheRepository.On("Delete", mock.AnythingOfType("string")).Return(func(name string) error {
		return fmt.Errorf("delete must be wrapped")
	})
	err = humanService.Update(ctx, name, human)
	require.Equal(t, fmt.Errorf("human service update func %w", fmt.Errorf("delete must be wrapped")), err)

}

func TestHumanService_Delete(t *testing.T) {
	// no error
	humanRepository = mocks.HumanRepository{}
	humanCacheRepository = mocks.HumanCacheRepository{}
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
	// delete human error
	humanRepository = mocks.HumanRepository{}
	humanRepository.On("Delete", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(func(ctx context.Context, name string) error {
		return fmt.Errorf("delete must be wrapped")
	})
	err = humanService.Delete(ctx, name)
	require.Equal(t, fmt.Errorf("human service %w", fmt.Errorf("delete must be wrapped")), err)
	// delete cache error
	humanRepository = mocks.HumanRepository{}
	humanCacheRepository = mocks.HumanCacheRepository{}
	humanRepository.On("Delete", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(func(ctx context.Context, name string) error {
		return nil
	})
	humanCacheRepository.On("Delete", mock.AnythingOfType("string")).Return(func(name string) error {
		return fmt.Errorf("cache must be wrapped")
	})
	err = humanService.Delete(ctx, name)
	require.Equal(t, fmt.Errorf("human service %w", fmt.Errorf("cache must be wrapped")), err)
}
