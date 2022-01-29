package service

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/jwtToken"
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

func (u *UserService) CreateUser(login string, password string) error {
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

func (u *UserService) UpdateUser(token string, oldLogin string, newUser model.User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error with hashing user's password in update function %w", err)
	}
	newUser.Password = string(hashedPass)
	login, role, err := jwtToken.DecodeToken(token)
	if err != nil {
		return fmt.Errorf("service layer update function %w", err)
	}
	if login == oldLogin || role == "admin" {
		err = u.r.Update(context.Background(), login, newUser)
		if err != nil {
			return fmt.Errorf("service layer update function %w", err)
		}
	} else {
		return fmt.Errorf("access denied in update")
	}
	return nil
}

func (u *UserService) DeleteUser(token string, loginToDelete string) error {
	login, role, err := jwtToken.DecodeToken(token)
	if err != nil {
		return fmt.Errorf("service layer update function %w", err)
	}
	if login == loginToDelete || role == "admin" {
		err = u.r.Delete(context.Background(), loginToDelete)
		if err != nil {
			return fmt.Errorf("service layer delete function %w", err)
		}
	} else {
		return fmt.Errorf("access denied in delete")
	}
	return nil
}

func (u *UserService) Authentication(login string, password string) (string, error) {
	user, err := u.r.Get(context.Background(), login)
	if err != nil {
		return "", fmt.Errorf("service layer authenticcation function %w", err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("authentication comparing passwords error %w", err)
	}
	token, err := jwtToken.EncodeToken(user)
	if err != nil {
		return "", fmt.Errorf("service layer authentication %w", err)
	}
	return token, nil
}
func (u *UserService) Get(token, loginToGet string) (*model.User, error) {
	login, role, err := jwtToken.DecodeToken(token)
	if err != nil {
		return nil, fmt.Errorf("service layer update function %w", err)
	}
	if login == loginToGet || role == "admin" {
		user, err := u.r.Get(context.Background(), loginToGet)
		if err != nil {
			return nil, fmt.Errorf("service layer get function %w", err)
		}
		return user, nil
	} else {
		return nil, fmt.Errorf("access denied in get")
	}
}
