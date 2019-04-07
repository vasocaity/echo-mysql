package usecase

import "github.com/vasocaity/echo-mysql/models"

type Service interface {
	GetAreas() *models.Areas
	GetArea() (*models.Areas, error)
	GetAreaByID(id string) (*models.Area, error)
}
