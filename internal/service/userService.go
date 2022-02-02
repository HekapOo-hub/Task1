package service

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	r repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{r: r}
}

func (u *UserService) Create(login string, password string) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error with hashing user's password %w", err)
	}
	user := model.User{Login: login, Password: string(hashedPass)}
	err = u.r.Create(context.Background(), user)
	if err != nil {
		return fmt.Errorf("service layer create function %w", err)
	}
	return nil
}

func (u *UserService) Update(oldLogin string, newUser model.User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error with hashing user's password in update function %w", err)
	}
	newUser.Password = string(hashedPass)

	err = u.r.Update(context.Background(), oldLogin, newUser)
	if err != nil {
		return fmt.Errorf("service layer update function %w", err)
	}

	return nil
}

func (u *UserService) Delete(loginToDelete string) error {
	err := u.r.Delete(context.Background(), loginToDelete)
	if err != nil {
		return fmt.Errorf("service layer delete function %w", err)
	}

	return nil
}

func (u *UserService) Get(login string) (*model.User, error) {
	user, err := u.r.Get(context.Background(), login)
	if err != nil {
		return nil, fmt.Errorf("service layer get function %w", err)
	}
	return user, nil

}
