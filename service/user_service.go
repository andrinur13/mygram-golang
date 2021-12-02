package service

import (
	"mygram-golang/model/entity"
	"mygram-golang/model/input"
	"mygram-golang/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(userInput input.RegisterUserInput) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) CreateUser(input input.RegisterUserInput) (entity.User, error) {
	newUser := entity.User{}
	newUser.Username = input.Username
	newUser.Age = input.Age
	newUser.Email = input.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return entity.User{}, err
	}

	newUser.Password = string(passwordHash)

	createdUser, err := s.userRepository.Save(newUser)

	if err != nil {
		return entity.User{}, err
	}

	return createdUser, nil
}
