package users

import (
	"fmt"
	"net/mail"
	"regexp"

	"github.com/google/uuid"
)

type User struct {
	name  string
	email string
	id    string
}

func NewUser(name, email, id string) (*User, error) {
	// Validate name
	matched, err := regexp.MatchString(`^[A-Za-zÀ-ÖØ-öø-ÿ'\\- ]{2,50}$`, name)
	if err != nil {
		return nil, fmt.Errorf("error validating name: %w", err)
	}

	if !matched {
		return nil, fmt.Errorf("invalid name format")
	}

	// Validate email
	_, err = mail.ParseAddress(email)
	if err != nil {
		return nil, fmt.Errorf("invalid email format: %w", err)
	}

	// Validate id with UUID
	_, err = uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id format: %w", err)
	}

	// All validations passed, create and return the User instance
	return &User{
		name:  name,
		email: email,
		id:    id,
	}, nil
}