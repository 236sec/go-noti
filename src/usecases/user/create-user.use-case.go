package user

import (
	"context"
	"errors"

	"goboilerplate.com/src/models"
	"goboilerplate.com/src/pkg/database"
	"goboilerplate.com/src/repo"
	"goboilerplate.com/src/usecases"
)



type ICreateUserUseCase interface {
	Apply(ctx context.Context, req CreateUserRequest) (CreateUserResponse, error)
}

type CreateUserUseCase struct {
	userRepo repo.IUserRepo
}

func NewCreateUserUseCase(userRepo repo.IUserRepo) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: userRepo}
}

func (u *CreateUserUseCase) Apply(ctx context.Context, req CreateUserRequest) (CreateUserResponse, error) {
	existingUser, err := u.userRepo.GetUserByUsername(ctx, req.Username)
	if err == nil && existingUser.ID != 0 {
		return CreateUserResponse{}, usecases.ErrUserAlreadyExists
	}
	if err != nil && !errors.Is(err, database.ErrRecordNotFound) {
		return CreateUserResponse{}, usecases.ErrInternalServerError
	}
	
	newUser, err := u.userRepo.CreateUser(ctx, models.User{
		Username:    req.Username,
		Password:    req.Password,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		DateOfBirth: req.DateOfBirth,
		Role:        "user", // Default role
	})
	if err != nil {
		return CreateUserResponse{}, usecases.ErrCannotCreateUser
	}
	
	return CreateUserResponse{
		ID: newUser.ID,
	}, nil
}