// Package service contains services which wrap repository and implement business logic
package service

import (
	"context"
	"fmt"

	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository"
)

// HumanService wraps human repository implementing business logic of app
type HumanService struct {
	r repository.Repository
}

// NewHumanService returns instance of HumanService
func NewHumanService(r repository.Repository) *HumanService {
	return &HumanService{r: r}
}

// Create is used for creating human info from db
func (s *HumanService) Create(ctx context.Context, h model.Human) error {
	err := s.r.Create(ctx, h)
	if err != nil {
		return fmt.Errorf("human service %w", err)
	}
	return nil
}

// Delete is used for deleting human info from db
func (s *HumanService) Delete(ctx context.Context, id string) error {
	err := s.r.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("human service %w", err)
	}
	return nil
}

// Update is used for updating human info in db
func (s *HumanService) Update(ctx context.Context, id string, h model.Human) error {
	err := s.r.Update(ctx, id, h)
	if err != nil {
		return fmt.Errorf("human service %w", err)
	}
	return nil
}

// Get is used for getting human info from db
func (s *HumanService) Get(ctx context.Context, name string) (*model.Human, error) {
	h, err := s.r.Get(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("human service %w", err)
	}
	return h, nil
}
