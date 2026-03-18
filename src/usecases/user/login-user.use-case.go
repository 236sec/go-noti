package user

import (
	"context"

	"goboilerplate.com/src/domain"
	"goboilerplate.com/src/repo"
	"goboilerplate.com/src/usecases"
)

type ILoginUserUseCase interface {
	Apply(ctx context.Context, req LoginUserRequest) (LoginUserResponse, error)
}

type LoginUserUseCase struct {
	userRepo repo.IUserRepo
}

func NewLoginUserUseCase(userRepo repo.IUserRepo) *LoginUserUseCase {
	return &LoginUserUseCase{userRepo: userRepo}
}

func (u *LoginUserUseCase) Apply(ctx context.Context, req LoginUserRequest) (LoginUserResponse, error) {
	modelUser, err := u.userRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return LoginUserResponse{}, usecases.ErrUserNotFound
	}
	
	domainUser := domain.FromModel(modelUser)
	
	if !domainUser.IsAbleToLogin() {
		return LoginUserResponse{}, usecases.ErrUserNotAbleToLogin
	}

	// TODO: Add proper password verification here
	// Example: bcrypt.CompareHashAndPassword([]byte(domainUser.Password), []byte(req.Password))
	// if domainUser.Password != req.Password {
	// 	return LoginUserResponse{}, usecases.ErrorInvalidCredentials
	// }
	
	// TODO: Generate proper JWT token here
	// For now, returning a placeholder token
	token := generateToken(domainUser)
	
	return LoginUserResponse{
		Token: token,
	}, nil
}

// generateToken creates a token for the user (placeholder implementation)
func generateToken(user domain.User) string {
	// TODO: Implement proper JWT token generation
	// This is just a placeholder
	return "jwt_token_for_" + user.GetFullName()
}