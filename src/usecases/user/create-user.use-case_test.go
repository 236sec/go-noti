package user

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"goboilerplate.com/src/models"
	"goboilerplate.com/src/repo/mocks"
	"goboilerplate.com/src/usecases"
	"gorm.io/gorm"
)

type assertions struct {
	name          string
	request       CreateUserRequest
	mockSetup     func(mockRepo *mocks.MockIUserRepo)
	expectedError error
	expectedResult CreateUserResponse
}
	
func TestCreateUserUseCase_Apply(t *testing.T) {
	// Define test cases
	testCases := []assertions{
		{
			name: "Error at query user - should not run domain logic",
			request: CreateUserRequest{
				Username:    "testuser",
				Password:    "password123",
				FirstName:   "Test",
				LastName:    "User",
				DateOfBirth: "2000-01-01",
			},
			mockSetup: func(mockRepo *mocks.MockIUserRepo) {
				// DB returns unexpected error → CreateUser should NOT be called
				mockRepo.EXPECT().GetUserByUsername(gomock.Any(), "testuser").Return(models.User{}, errors.New("db connection error")).Times(1)
				mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Times(0)
			},
			expectedError: usecases.ErrInternalServerError,
		},
		{
			name: "Pass the db - run user domain logic once",
			request: CreateUserRequest{
				Username:    "newuser",
				Password:    "password123",
				FirstName:   "New",
				LastName:    "User",
				DateOfBirth: "2000-01-01",
			},
			mockSetup: func(mockRepo *mocks.MockIUserRepo) {
				// User not found → CreateUser should be called exactly once
				mockRepo.EXPECT().GetUserByUsername(gomock.Any(), "newuser").Return(models.User{}, gorm.ErrRecordNotFound).Times(1)
				mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(models.User{
					ID:          1,
					FirstName:   "New",
					LastName:    "User",
					Username:    "newuser",
					Password:    "password123",
					Role:        "user",
					DateOfBirth: "2000-01-01",
				}, nil).Times(1)
			},
			expectedError: nil,
			expectedResult: CreateUserResponse{
				ID:          1,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockIUserRepo(ctrl)
			tt.mockSetup(mockRepo)

			useCase := NewCreateUserUseCase(mockRepo)
			result, err := useCase.Apply(context.Background(), tt.request)

			assert.ErrorIs(t,tt.expectedError,err)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}