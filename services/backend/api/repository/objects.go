package repository

import (
	"backend/models"
	"time"
)

type ObjectRepository struct {
	objects []models.Object
}

func NewObjectRepository() *ObjectRepository {
	objects := make([]models.Object, 150)
	for i := 0; i < 150; i++ {
		objects[i] = models.Object{
			Id: time.Now().Nanosecond(),
			Num: i,
			Text: time.Now().String(),
		}
	}
	return &ObjectRepository{
		objects: objects,
	}
}

func (r *ObjectRepository) GetAllObjects() ([]models.Object, error) {
	return r.objects, nil
}