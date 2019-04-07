package usecase

import (
	models "github.com/vasocaity/echo-mysql/models"
	"github.com/vasocaity/echo-mysql/repository"
)

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo}
}

func (s *service) GetAreas() *models.Areas {
	return s.repo.GetAreas()
}

func (s *service) GetArea() (*models.Areas, error) {
	return s.repo.GetArea()

}
func (s *service) GetAreaByID(id string) (*models.Area, error) {
	return s.repo.GetAreaByID(id)
}
