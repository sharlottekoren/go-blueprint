package service

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/sharlottekoren/go-blueprint/internal/service/mocks"
	"github.com/sharlottekoren/go-blueprint/internal/domain/users"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService(t *testing.T) {
	test := assert.New(t)

	// Define test cases
	type TestCase struct {
		testName	   string
		userID         string
		errContains    string
		mockUserRepoFn func() UserRepository
	}

	// Generate an ID and a new user for testing
	id := uuid.New()
	newUser, _ := users.NewUser("John Doe", "abc@def.com", id.String())

	testCases := []TestCase{
		{
			testName: "Given valid user ID, when GetUserByID is called, then return the user",
			userID:   id.String(),
			mockUserRepoFn: func() UserRepository {
				ctrl := gomock.NewController(t)
				mockUserRepo := mocks.NewMockUserRepository(ctrl)
				mockUserRepo.EXPECT().GetUserByID(id.String()).Return(newUser, nil)
				return mockUserRepo
			},
		},
		{
			testName:    "Given the repository returns an error, when GetUserByID is called, then return the error",
			userID:      id.String(),
			errContains: "repository returned an error: blah",
			mockUserRepoFn: func() UserRepository {
				ctrl := gomock.NewController(t)
				mockUserRepo := mocks.NewMockUserRepository(ctrl)
				mockUserRepo.EXPECT().GetUserByID(id.String()).Return(nil, fmt.Errorf("blah"))
				return mockUserRepo
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			mockUserRepo := tc.mockUserRepoFn()
			svc := NewService(mockUserRepo)
			user, err := svc.GetUserByID(tc.userID)

			if err != nil {
				if tc.errContains != "" {
					test.Contains(err.Error(), tc.errContains)
				} else {
					t.Fatalf("unexpected error: %v", err)
				}
			} else {
				if tc.errContains != "" {
					t.Fatalf("expected error containing %q, but got nil", tc.errContains)
				}
				test.Equal(newUser, user)
			}
		})
	}
}