package service

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository"
)

const (
	admin = "admin"
	user  = "user"
)

type HumanService struct {
	r repository.Repository
}

func NewService(r repository.Repository) *HumanService {
	return &HumanService{r: r}
}
func (s *HumanService) Create(h model.Human) error {
	err := s.r.Create(context.Background(), h)
	if err != nil {
		return fmt.Errorf("human service %w", err)
	}
	return nil
}
func (s *HumanService) Delete(id string) error {
	err := s.r.Delete(context.Background(), id)
	if err != nil {
		return fmt.Errorf("human service %w", err)
	}
	return nil
}
func (s *HumanService) Update(id string, h model.Human) error {
	err := s.r.Update(context.Background(), id, h)
	if err != nil {
		return fmt.Errorf("human service %w", err)
	}
	return nil
}
func (s *HumanService) Get(name string) (*model.Human, error) {
	h, err := s.r.Get(context.Background(), name)
	if err != nil {
		return nil, fmt.Errorf("human service %w", err)
	}
	return h, nil
}
