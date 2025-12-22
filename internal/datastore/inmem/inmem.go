package inmem

import (
	"fmt"

	"github.com/sharlottekoren/go-blueprint/internal/domain/users"
)

// We can define some common errors for the in-memory datastore so that they can be checked against in tests and higher-level logic.
var UserNotFoundError = fmt.Errorf("user not found")
var UserAlreadyExistsError = fmt.Errorf("user already exists")

// InMemUserStore is an in-memory implementation of a user data store.
type InMemUserStore struct {
	users map[string]*users.User
}

// NewInMemUserStore creates a new instance of InMemUserStore.
func NewInMemUserStore() *InMemUserStore {
	return &InMemUserStore{
		users: make(map[string]*users.User),
	}
}

// GetUserByID retrieves a user by their ID.
func (store *InMemUserStore) GetUserByID(id string) (*users.User, error) {
	item, found := store.users[id]
	if !found {
		return nil, UserNotFoundError
	}
	return item, nil
}

// CreateNewUser adds a new user to the store.
func (store *InMemUserStore) CreateNewUser(user *users.User) error {
	_, found := store.users[user.GetID()]
	if found {
		return UserAlreadyExistsError
	}
	store.users[user.GetID()] = user
	return nil
}