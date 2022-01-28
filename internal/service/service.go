package service

import (
	"context"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository"
)

type Service struct {
	r repository.Repo
}

func NewService(r repository.Repo) *Service {
	return &Service{r: r}
}
func (s *Service) CreateHuman(h model.Human) error {
	return s.r.Create(context.Background(), h)
}
func (s *Service) DeleteHuman(id string) error {
	return s.r.Delete(context.Background(), id)
}
func (s *Service) UpdateHuman(id string, h model.Human) error {
	return s.r.Update(context.Background(), id, h)
}
func (s *Service) GetHumanInfo(name string) (*model.Human, error) {
	return s.r.Get(context.Background(), name)
}
