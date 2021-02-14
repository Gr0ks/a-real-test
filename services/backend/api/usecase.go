package api

import (
	"backend/models"
)

type UseCase interface {
	GetObjects(firstNumber, count int) (*[]models.Object, error)
}