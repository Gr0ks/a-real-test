package usecase

import (
	"backend/models"
	"backend/api"
	"errors"
)

type ObjectCreator struct {
	repo api.Repository
}

func NewObjectCreator(repo api.Repository) *ObjectCreator {
	return &ObjectCreator {
		repo: repo,
	}
}

func (o *ObjectCreator) GetObjects(firstNumber, count int) ([]models.Object, error) {
	
	objects, _ := o.repo.GetAllObjects()
	if firstNumber > len(objects) {
		return []models.Object{}, errors.New("No more objects in repo")
	}

	lastObjectNumber := firstNumber + count
	if lastObjectNumber > len(objects) {
		lastObjectNumber = len(objects)
		count = lastObjectNumber - firstNumber
	}

	result := make([]models.Object, count)
	for i := firstNumber; i < lastObjectNumber; i++ {
		result[i] = objects[i]
	}

	return result, nil
}