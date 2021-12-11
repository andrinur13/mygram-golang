package repository

import (
	"mygram-golang/model/entity"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Save(photo entity.Photo) (entity.Photo, error)
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{db}
}
