package service

import (
	"fmt"

	"github.com/sharlottekoren/go-blueprint/internal/domain/users"
)

// Service struct that holds the UserRepository.
type Service struct {
	userRepo UserRepository
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