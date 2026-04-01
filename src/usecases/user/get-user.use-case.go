package user

import (
	"context"

	"go.opentelemetry.io/otel"
	"goboilerplate.com/src/repo"
	"goboilerplate.com/src/usecases"
)

var getUserTracer = otel.Tracer("usecase.getuser")

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
	ctx, span := getUserTracer.Start(ctx, "GetUserUseCase.Apply")
	defer span.End()

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
