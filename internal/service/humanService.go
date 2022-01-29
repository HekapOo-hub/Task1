package service

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/jwtToken"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository"
)

type HumanService struct {
	r repository.Repository
}

func NewService(r repository.Repository) *HumanService {
	return &HumanService{r: r}
}
func (s *HumanService) Create(token string, h model.Human) error {
	_, role, err := jwtToken.DecodeToken(token)
	if err != nil {
		return fmt.Errorf("humanService layer create function %w", err)
	}
	if role == "admin" {
		return s.r.Create(context.Background(), h)
	} else if role == "user" {
		return fmt.Errorf("access denied")
	} else {
		return fmt.Errorf("please authenticate in system to work with human data")
	}
}
func (s *HumanService) Delete(token string, id string) error {
	_, role, err := jwtToken.DecodeToken(token)
	if err != nil {
		return fmt.Errorf("humanService layer delete function %w", err)
	}
	if role == "admin" {
		return s.r.Delete(context.Background(), id)
	} else if role == "user" {
		return fmt.Errorf("access denied")
	} else {
		return fmt.Errorf("please authenticate in system to work with human data")
	}
}
func (s *HumanService) Update(token string, id string, h model.Human) error {
	_, role, err := jwtToken.DecodeToken(token)
	if err != nil {
		return fmt.Errorf("humanService layer update function %w", err)
	}
	if role == "admin" {
		return s.r.Update(context.Background(), id, h)
	} else if role == "user" {
		return fmt.Errorf("access denied")
	} else {
		return fmt.Errorf("please authenticate in system to work with human data")
	}
}
func (s *HumanService) Get(token string, name string) (*model.Human, error) {
	_, role, err := jwtToken.DecodeToken(token)
	if err != nil {
		return nil, fmt.Errorf("humanService layer create function %w", err)
	}
	if role == "admin" || role == "user" {
		return s.r.Get(context.Background(), name)
	} else {
		return nil, fmt.Errorf("please authenticate in system to work with human data")
	}
}
