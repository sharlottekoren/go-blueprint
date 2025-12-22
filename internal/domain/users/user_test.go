package users

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------- UNIT TESTS -----------------
// This is how you can write unit tests for the NewUser function.

// TestNewUser_Valid tests the NewUser function with valid inputs.
// Given there are valid inputs, when creating a new user, then no error should be returned.
func TestNewUser_Valid(t *testing.T) {
	name := "John Doe"
	email := "correct@hello.com"
	id := "550e8400-e29b-41d4-a716-446655440000"

	user, err := NewUser(name, email, id)
	// If there is an error, there shouldn't be one so the test should fail.
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	// Check that the name, email and id are set correctly.
	if user.name != name {
		t.Errorf("expected name %s, got %s", name, user.name)
	}
	if user.email != email {
		t.Errorf("expected email %s, got %s", email, user.email)
	}
	if user.id != id {
		t.Errorf("expected id %s, got %s", id, user.id)
	}
}

// TestNewUser_Bad_Email tests the NewUser function with an invalid email format.
// Given there is an invalid email format, when creating a new user, then an error should be returned.
func TestNewUser_Bad_Email(t *testing.T) {
	name := "John Doe"
	email := "bad-email-format"
	id := "550e8400-e29b-41d4-a716-446655440000"

	_, err := NewUser(name, email, id)
	// If there is no error, the test should fail.
	if err == nil {
		t.Fatal("expected error for invalid email format, got none")
	}
	// Check that the error message is as expected.
	if !strings.Contains(err.Error(), "invalid email format") {
		t.Errorf("expected error message to contain 'invalid email format', got %v", err)
	}
}

// TestNewUser_Bad_Name tests the NewUser function with an invalid name format.
// Given there is an invalid name format, when creating a new user, then an error should be returned.
func TestNewUser_Bad_Name(t *testing.T) {
	name := "J"
	email := "correct@hello.com"
	id := "550e8400-e29b-41d4-a716-446655440000"

	_, err := NewUser(name, email, id)
	// If there is no error, the test should fail.
	if err == nil {
		t.Fatal("expected error for invalid name format, got none")
	}
	// Check that the error message is as expected.
	if !strings.Contains(err.Error(), "invalid name format") {
		t.Errorf("expected error message to contain 'invalid name format', got %v", err)
	}
}

// TestNewUser_Bad_ID tests the NewUser function with an invalid UUID format.
// Given there is an invalid UUID format, when creating a new user, then an error should be returned.
func TestNewUser_Bad_ID(t *testing.T) {
	name := "John Doe"
	email := "correct@hello.com"
	id := "not-a-uuid"

	_, err := NewUser(name, email, id)
	// If there is no error, the test should fail.
	if err == nil {
		t.Fatal("expected error for invalid id format, got none")
	}
	// Check that the error message is as expected.
	if !strings.Contains(err.Error(), "invalid id format") {
		t.Errorf("expected error message to contain 'invalid id format', got %v", err)
	}
}

// ----------------- TABLE TESTS -----------------
// This is how you can write table tests for the NewUser function.

func TestNewUser(t *testing.T) {
	// Using testify for assertions
	test := assert.New(t)

	// Define a struct that represents each test case
	type testCase struct {
		name        string
		email       string
		id          string
		errContains string
	}

	testCases := map[string]testCase{
		"Given valid inputs, when creating a new user, then no error should be returned":
			{
				"John Doe",
				"correct@hello.com",
				"550e8400-e29b-41d4-a716-446655440000",
				"",
			},
		"Given invalid email, when creating a new user, then an error should be returned":
			{
				"John Doe",
				"correcthello.com",
				"550e8400-e29b-41d4-a716-446655440000",
				"invalid email",
			},
		"Given invalid name, when creating a new user, then an error should be returned":
			{
				"John123 Doe",
				"correct@hello.com",
				"550e8400-e29b-41d4-a716-446655440000",
				"invalid name",
			},
		"Given invalid ID, when creating a new user, then an error should be returned":
			{
				"John Doe",
				"correct@hello.com",
				"55-",
				"invalid UUID",
			},
		"Given empty name, when creating a new user, then an error should be returned":
			{
				"",
				"correct@hello.com",
				"550e8400-e29b-41d4-a716-446655440000",
				"invalid name",
			},
		"Given empty email, when creating a new user, then an error should be returned":
			{
				"John Doe",
				"",
				"550e8400-e29b-41d4-a716-446655440000",
				"invalid email",
			},
	}

	// Iterate over each test case
	for testName, tc := range testCases {
		// Run each test case as a subtest
		t.Run(testName, func(t *testing.T) {
			newUser, err := NewUser(tc.name, tc.email, tc.id)
			if tc.errContains == "" {
				// No error expected
				test.NoError(err)
				// Check that the returned user has the expected values
				test.Equal(tc.name, newUser.name)
				test.Equal(tc.email, newUser.email)
				test.Equal(tc.id, newUser.id)
			} else {
				// Error expected
				test.Error(err)
				test.ErrorContains(err, tc.errContains)
			}
		})
	}
}
