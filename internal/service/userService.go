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

func (u *UserService) Update(login string, role string, oldLogin string, newUser model.User) error {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error with hashing user's password in update function %w", err)
	}
	newUser.Password = string(hashedPass)
	if login == oldLogin || role == admin {
		err = u.r.Update(context.Background(), login, newUser)
		if err != nil {
			return fmt.Errorf("service layer update function %w", err)
		}
	} else {
		return fmt.Errorf("access denied in update")
	}
	return nil
}

func (u *UserService) Delete(login string, role string, loginToDelete string) error {
	if login == loginToDelete || role == admin {
		err := u.r.Delete(context.Background(), loginToDelete)
		if err != nil {
			return fmt.Errorf("service layer delete function %w", err)
		}
	} else {
		return fmt.Errorf("access denied in delete")
	}
	return nil
}

func (u *UserService) Get(login string, role string, loginToGet string) (*model.User, error) {

	if login == loginToGet || role == admin {
		user, err := u.r.Get(context.Background(), loginToGet)
		if err != nil {
			return nil, fmt.Errorf("service layer get function %w", err)
		}
		return user, nil
	} else {
		return nil, fmt.Errorf("access denied in get")
	}
}
