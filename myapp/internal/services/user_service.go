package services

import (
    "myapp/internal/models"
    "myapp/internal/repositories"
)

type UserService struct {
    userRepository *repositories.UserRepository
}

func NewUserService() *UserService {
    userRepository := repositories.NewUserRepository()
    return &UserService{userRepository: userRepository}
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
    return s.userRepository.FindAll()
}

