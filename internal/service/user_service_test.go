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

var userRepository mocks.UserRepository

func TestUserService_Create(t *testing.T) {
	// no error
	userRepository = mocks.UserRepository{}
	ctx := context.Background()
	login := "login"
	password := "pwd"
	userService := NewUserService(&userRepository)
	userRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("model.User")).Return(func(ctx context.Context, u model.User) error {
		return nil
	})
	err := userService.Create(ctx, login, password)
	require.NoError(t, err)
	// create error
	userRepository = mocks.UserRepository{}
	userRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("model.User")).Return(func(ctx context.Context, u model.User) error {
		return fmt.Errorf("create must be wrapped")
	})
	err = userService.Create(ctx, login, password)
	require.Equal(t, fmt.Errorf("service layer create function %w", fmt.Errorf("create must be wrapped")), err)
}

func TestUserService_Get(t *testing.T) {
	// no error
	userRepository = mocks.UserRepository{}
	ctx := context.Background()
	login := "login"
	userService := NewUserService(&userRepository)
	userRepository.On("Get", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(
		func(ctx context.Context, login string) *model.User {
			return &model.User{Login: login}
		},
		func(ctx context.Context, login string) error {
			return nil
		})
	_, err := userService.Get(ctx, login)
	require.NoError(t, err)
	// get error
	userRepository = mocks.UserRepository{}
	userRepository.On("Get", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(
		func(ctx context.Context, login string) *model.User {
			return nil
		},
		func(ctx context.Context, login string) error {
			return fmt.Errorf("get must be wrapped")
		})
	_, err = userService.Get(ctx, login)
	require.Equal(t, fmt.Errorf("service layer get function %w", fmt.Errorf("get must be wrapped")), err)

}

func TestUserService_Update(t *testing.T) {
	// no error
	userRepository = mocks.UserRepository{}
	ctx := context.Background()
	oldLogin := "old"
	userService := NewUserService(&userRepository)
	user := model.User{Login: "new", Password: "qerfwegf"}
	userRepository.On("Update", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string"),
		mock.AnythingOfType("model.User")).Return(func(ctx context.Context, login string, user model.User) error {
		return nil
	})
	err := userService.Update(ctx, oldLogin, user)
	require.NoError(t, err)
	// update error
	userRepository = mocks.UserRepository{}
	userRepository.On("Update", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string"),
		mock.AnythingOfType("model.User")).Return(func(ctx context.Context, login string, user model.User) error {
		return fmt.Errorf("update must be wrapped")
	})
	err = userService.Update(ctx, oldLogin, user)
	require.Equal(t, fmt.Errorf("service layer update function %w", fmt.Errorf("update must be wrapped")), err)

}

func TestUserService_Delete(t *testing.T) {
	// no error
	userRepository = mocks.UserRepository{}
	ctx := context.Background()
	login := "delete"
	userService := NewUserService(&userRepository)
	userRepository.On("Delete", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(
		func(ctx context.Context, s string) error {
			return nil
		})
	err := userService.Delete(ctx, login)
	require.NoError(t, err)
	// delete error
	userRepository = mocks.UserRepository{}
	userRepository.On("Delete", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(
		func(ctx context.Context, s string) error {
			return fmt.Errorf("delete must be wrapped")
		})
	err = userService.Delete(ctx, login)
	require.Equal(t, fmt.Errorf("service layer delete function %w",
		fmt.Errorf("delete must be wrapped")), err)
}
