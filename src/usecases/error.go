package usecases

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserNotAuthorized = errors.New("user is not authorized to login")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotAbleToLogin = errors.New("user is not able to login")
	ErrCannotCreateUser      = errors.New("cannot create user")
	ErrInternalServerError = errors.New("internal server error")
	ErrUserAlreadyExists = errors.New("user already exists")
)