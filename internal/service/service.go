package service

import (
	"context"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository"
)

type Service struct {
	repository.Repo
}

func (s *Service) CreateHuman(h model.Human) error {
	return s.Create(context.Background(), h)
}
func (s *Service) DeleteHuman(id int) error {
	return s.Repo.Delete(context.Background(), id)
}
func (s *Service) UpdateHuman(id int, h model.Human) error {
	return s.Repo.Update(context.Background(), id, h)
}
func (s *Service) GetHumanInfo(id int) (*model.Human, error) {
	return s.Repo.Get(context.Background(), id)
}
