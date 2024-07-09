package service

import (
	"errors"
	"pet-matching-service/internal/model"
	"pet-matching-service/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) Authenticate(username, password string) (*model.User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAllUsers()
}
