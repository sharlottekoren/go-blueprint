package service

import (
	"fmt"

	"github.com/sharlottekoren/go-blueprint/internal/domain/users"
)

// Service struct that holds the UserRepository.
type Service struct {
	userRepo UserRepository
}

// CreateUserRequest represents the data needed to create a new user.
type CreateUserRequest struct {
	Name  string
	Email string
	ID    string
}

// NewService creates a new instance of Service with the provided UserRepository.
func NewService(userRepo UserRepository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

// GetUserByID retrieves a user by their ID using the UserRepository.
func (s *Service) GetUserByID(id string) (*users.User, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("repository returned an error: %w", err)
	}
	return user, nil
}

// CreateUser creates a new user and saves it using the UserRepository.
func (s *Service) CreateUser(req CreateUserRequest) (*users.User, error) {
	// Create a new user instance
	newUser, err := users.NewUser(req.Name, req.Email, req.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create user object: %w", err)
	}
	// Save the new user using the repository
	err = s.userRepo.CreateNewUser(newUser)
	if err != nil {
		return nil, fmt.Errorf("failed to add new user to repository: %w", err)
	}
	return newUser, nil
}