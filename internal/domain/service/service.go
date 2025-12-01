package service

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