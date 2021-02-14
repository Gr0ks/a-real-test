package api

import (
	"backend/models"
)

type Repository interface {
	GetAllObjects() ([]models.Object, error)
}