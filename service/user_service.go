package service

import (
	"errors"
	"mygram-golang/model/entity"
	"mygram-golang/model/input"
	"mygram-golang/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(userInput input.RegisterUserInput) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	GetUserByID(ID int) (entity.User, error)
	UpdateUser(ID int, input input.UpdateUserInput) (entity.User, error)
	DeleteUser(ID int) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
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

func (s *userService) GetUserByEmail(email string) (entity.User, error) {
	user, err := s.userRepository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return entity.User{}, nil
	}

	return user, nil
}

func (s *userService) GetUserByID(ID int) (entity.User, error) {
	user, err := s.userRepository.FindByID(ID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return entity.User{}, nil
	}

	return user, nil
}

func (s *userService) UpdateUser(ID int, input input.UpdateUserInput) (entity.User, error) {
	userResult, err := s.userRepository.FindByID(ID)

	if err != nil {
		return entity.User{}, err
	}

	if userResult.ID == 0 {
		return entity.User{}, errors.New("user not found!")
	}

	updatedUser := entity.User{
		Username: input.Username,
		Email:    input.Email,
		Age:      input.Age,
	}

	if input.Passowrd != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Passowrd), bcrypt.MinCost)

		if err != nil {
			return entity.User{}, err
		}

		updatedUser.Password = string(passwordHash)
	}

	userUpdate, err := s.userRepository.Update(ID, updatedUser)

	if err != nil {
		return userUpdate, err
	}

	return userUpdate, nil
}

func (s *userService) DeleteUser(ID int) (entity.User, error) {
	userQuery, err := s.userRepository.FindByID(ID)

	if err != nil {
		return entity.User{}, err
	}

	if userQuery.ID == 0 {
		return entity.User{}, nil
	}

	userDeleted, err := s.userRepository.Delete(ID)

	if err != nil {
		return entity.User{}, err
	}

	return userDeleted, nil
}
