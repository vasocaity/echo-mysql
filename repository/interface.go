package repository

import models "github.com/vasocaity/echo-mysql/models"

type Repository interface {
	GetAreas() models.Areas
	GetArea() (*models.Areas, error)
	GetAreaByID(id string) (*models.Area, error)
}
