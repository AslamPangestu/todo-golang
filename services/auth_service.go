package services

import (
	"errors"
	"todo-be/dtos"
	"todo-be/entities"
	"todo-be/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthInteractor interface {
	GetUserByID(id int) (entities.User, error)
	GetUserByUserID(payload dtos.LoginRequest) (entities.User, error)
	AddUser(payload dtos.RegisterRequest) (entities.User, error)
}

type authService struct {
	repository repositories.AuthInteractor
}

func NewAuthService(repository repositories.AuthInteractor) *authService {
	return &authService{repository}
}

func (s *authService) GetUserByID(id int) (entities.User, error) {
	model, err := s.repository.FindOne(id)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (s *authService) GetUserByUserID(payload dtos.LoginRequest) (entities.User, error) {
	model, err := s.repository.FindOneByUserID(payload.UserID)
	if err != nil {
		return model, err
	}

	if model.ID == 0 {
		return model, errors.New("USER NOT FOUND")
	}

	err = bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(payload.Password))
	if err != nil {
		return model, errors.New("PASSWORD INCORRECT")
	}
	return model, nil
}

func (s *authService) AddUser(payload dtos.RegisterRequest) (entities.User, error) {
	var model entities.User

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.MinCost)
	if err != nil {
		return model, err
	}

	model = entities.User{
		Name:     payload.Name,
		UserID:   payload.UserID,
		Password: string(passwordHash),
	}

	newUser, err := s.repository.Create(model)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
