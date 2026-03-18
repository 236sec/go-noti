package user

import (
	"context"

	"goboilerplate.com/src/repo"
	"goboilerplate.com/src/usecases"
)

type IGetUserUseCase interface {
	Apply(ctx context.Context, username string) (GetUserResponse, error)
}

type GetUserUseCase struct {
	userRepo repo.IUserRepo
}

func NewGetUserUseCase(userRepo repo.IUserRepo) *GetUserUseCase {
	return &GetUserUseCase{userRepo: userRepo}
}

func (u *GetUserUseCase) Apply(ctx context.Context, username string) (GetUserResponse, error) {
	user, err := u.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return GetUserResponse{}, usecases.ErrUserNotFound
	}
	return GetUserResponse{
		ID:          user.ID,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		DateOfBirth: user.DateOfBirth,
	}, nil
}