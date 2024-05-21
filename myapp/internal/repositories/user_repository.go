package repositories

import (
    "myapp/internal/models"
)

type UserRepository struct {
    // This can be a database connection in real implementation
}

func NewUserRepository() *UserRepository {
    return &UserRepository{}
}

func (r *UserRepository) FindAll() ([]*models.User, error) {
    // This is a mock implementation
    return []*models.User{
        {ID: 1, Name: "John Doe"},
    }, nil
}

