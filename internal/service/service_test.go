package service

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/sharlottekoren/go-blueprint/internal/domain/users"
	"github.com/sharlottekoren/go-blueprint/internal/service/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// TestService tests the Service struct's methods.
func TestService_GetUserByID(t *testing.T) {
	test := assert.New(t)

	// Define test cases
	type testCase struct {
		id         string
		errContains    string
		mockUserRepoFn func() UserRepository
	}

	// Generate an ID and a new user for testing
	id := uuid.New()
	newUser, _ := users.NewUser("John Doe", "abc@def.com", id.String())

	testCases := map[string]testCase{
		"Given valid inputs, when getting a user by ID, then no error should be returned":
		{
			id:          id.String(),
			errContains: "",
			mockUserRepoFn: func() UserRepository {
				ctrl := gomock.NewController(t)
				mockUserRepo := mocks.NewMockUserRepository(ctrl)
				mockUserRepo.EXPECT().GetUserByID(id.String()).Return(newUser, nil)
				return mockUserRepo
			},
		},
		"Given repository returns an error, when getting a user by ID, then an error should be returned":
		{
			id:          id.String(),
			errContains: "repository returned an error: blah",
			mockUserRepoFn: func() UserRepository {
				ctrl := gomock.NewController(t)
				mockUserRepo := mocks.NewMockUserRepository(ctrl)
				mockUserRepo.EXPECT().GetUserByID(id.String()).Return(nil, fmt.Errorf("blah"))
				return mockUserRepo
			},
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			// Initialise service with mock repository
			mockUserRepo := tc.mockUserRepoFn()
			svc := NewService(mockUserRepo)
			// Call GetUserByID
			user, err := svc.GetUserByID(tc.id)
			if tc.errContains == "" {
				// No error expected
				test.NoError(err)
				test.EqualValues(tc.id, user.GetID())
			} else {
				// Error expected
				test.Error(err)
				test.ErrorContains(err, tc.errContains)
			}
		})
	}
}
