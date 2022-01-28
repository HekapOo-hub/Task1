package service

import (
	"context"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository"
)

type Service struct {
	r repository.Repository
}

func NewService(r repository.Repository) *Service {
	return &Service{r: r}
}
func (s *Service) Create(h model.Human) error {
	return s.r.Create(context.Background(), h)
}
func (s *Service) Delete(id string) error {
	return s.r.Delete(context.Background(), id)
}
func (s *Service) Update(id string, h model.Human) error {
	return s.r.Update(context.Background(), id, h)
}
func (s *Service) Get(name string) (*model.Human, error) {
	return s.r.Get(context.Background(), name)
}
