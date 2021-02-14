package repository

import (
	"backend/models"
	"time"
)

type ObjectRepository struct {
	objects []models.Object
}

func NewObjectRepository() *ObjectRepository {
	var objects []models.Object
	for i := 0; i < 150; i++ {
		objects = append(objects, models.Object{
			Id: time.Now().Nanosecond(),
			Text: time.Now().String(),
		})
	}
	return &ObjectRepository{
		objects: objects,
	}
}

func (r *ObjectRepository) GetObjects(firstNumber, count int) (*[]models.Object, error) {
	return &r.objects, nil
}