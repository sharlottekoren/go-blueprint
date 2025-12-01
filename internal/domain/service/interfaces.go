package service

import "github.com/sharlottekoren/go-blueprint/internal/domain/users"

// UserRepository defines the interface for doing user-related data operations.
type UserRepository interface {
	GetUserByID(id string) (users.User, error)
	CreateNewUser(user users.User) error
}
