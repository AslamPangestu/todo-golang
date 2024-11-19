package repositories

import (
	"todo-be/entities"

	"gorm.io/gorm"
)

type AuthInteractor interface {
	FindOne(id int) (entities.User, error)
	FindOneByUserID(userID string) (entities.User, error)
	Create(payload entities.User) (entities.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) FindOne(id int) (entities.User, error) {
	var model entities.User
	err := r.db.First(&model, id).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *authRepository) FindOneByUserID(userID string) (entities.User, error) {
	var model entities.User
	err := r.db.Where("user_id = ?", userID).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *authRepository) Create(payload entities.User) (entities.User, error) {
	err := r.db.Create(&payload).Error
	if err != nil {
		return payload, err
	}
	return payload, nil
}
