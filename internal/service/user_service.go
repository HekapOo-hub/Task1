package service

import (
	"context"
	"fmt"

	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserService implements business logic of our app using UserRepository
type UserService struct {
	r repository.UserRepository
}

// NewUserService returns new user service struct
func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{r: r}
}

// Create is used for encrypting password and  creating new user info  in db
func (u *UserService) Create(ctx context.Context, login, password string) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error with hashing user's password %w", err)
	}
	user := model.User{Login: login, Password: string(hashedPass)}
	err = u.r.Create(ctx, user)
	if err != nil {
		return fmt.Errorf("service layer create function %w", err)
	}
	return nil
}

// Update is used for updating user info in db
func (u *UserService) Update(ctx context.Context, oldLogin string, newUser model.User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error with hashing user's password in update function %w", err)
	}
	newUser.Password = string(hashedPass)

	err = u.r.Update(ctx, oldLogin, newUser)
	if err != nil {
		return fmt.Errorf("service layer update function %w", err)
	}

	return nil
}

// Delete is used for deleting user info from db
func (u *UserService) Delete(ctx context.Context, loginToDelete string) error {
	err := u.r.Delete(ctx, loginToDelete)
	if err != nil {
		return fmt.Errorf("service layer delete function %w", err)
	}

	return nil
}

// Get is used for getting user info from db
func (u *UserService) Get(ctx context.Context, login string) (*model.User, error) {
	user, err := u.r.Get(ctx, login)
	if err != nil {
		return nil, fmt.Errorf("service layer get function %w", err)
	}
	return user, nil
}
