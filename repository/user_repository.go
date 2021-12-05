package repository

import (
	"mygram-golang/model/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user entity.User) (entity.User, error)
	// Update(user entity.User) (entity.User, error)
	// Delete(ID int) (entity.User, error)
	// Get(ID int) (entity.User, error)
	// GetAll() ([]entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
