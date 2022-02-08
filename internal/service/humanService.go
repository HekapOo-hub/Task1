// Package service contains services which wrap repository and implement business logic
package service

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"

	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository"
)

// HumanService wraps human repository implementing business logic of app
type HumanService struct {
	r     repository.Repository
	cache repository.HumanCacheRepository
}

// NewHumanService returns instance of HumanService
func NewHumanService(r repository.Repository, cache repository.HumanCacheRepository) *HumanService {
	return &HumanService{r: r, cache: cache}
}

// Create is used for creating human info from db
func (s *HumanService) Create(ctx context.Context, h model.Human) error {
	h.ID = uuid.NewV1().String()
	err := s.r.Create(ctx, h)
	if err != nil {
		return fmt.Errorf("human service %w", err)
	}
	err = s.cache.Create(h)
	if err != nil {
		return fmt.Errorf("human service %w", err)
	}
	return nil
}

// Delete is used for deleting human info from db
func (s *HumanService) Delete(ctx context.Context, name string) error {
	err := s.r.Delete(ctx, name)
	if err != nil {
		return fmt.Errorf("human service %w", err)
	}
	err = s.cache.Delete(name)
	if err != nil {
		return fmt.Errorf("human service %w", err)
	}
	return nil
}

// Update is used for updating human info in db
func (s *HumanService) Update(ctx context.Context, name string, h model.Human) error {
	err := s.r.Update(ctx, name, h)
	if err != nil {
		return fmt.Errorf("human service %w", err)
	}
	err = s.cache.Update(name, h)
	if err != nil {
		return fmt.Errorf("human service %w", err)
	}
	return nil
}

// Get is used for getting human info from db
func (s *HumanService) Get(ctx context.Context, name string) (*model.Human, error) {
	res, err := s.cache.Get(name)
	if err == nil {

		return res, nil
	}
	log.WithField("error", err).Warn("redis")
	h, err := s.r.Get(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("human service %w", err)
	}
	return h, nil
}
