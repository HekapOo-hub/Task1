package service

import (
	"context"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository"
)

type HumanService struct {
	r repository.Repository
}

func NewService(r repository.Repository) *HumanService {
	return &HumanService{r: r}
}
func (s *HumanService) Create(h model.Human) error {
	return s.r.Create(context.Background(), h)
}
func (s *HumanService) Delete(id string) error {
	return s.r.Delete(context.Background(), id)
}
func (s *HumanService) Update(id string, h model.Human) error {
	return s.r.Update(context.Background(), id, h)
}
func (s *HumanService) Get(name string) (*model.Human, error) {
	return s.r.Get(context.Background(), name)
}
