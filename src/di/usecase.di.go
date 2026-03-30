package di


import (
	"sync"

	"goboilerplate.com/src/usecases"
	"goboilerplate.com/src/usecases/email"
	"goboilerplate.com/src/usecases/user"
)

var GetHealthUseCase = sync.OnceValue(func() *usecases.HealthUseCase {
	return usecases.NewHealthUseCase()
})

var GetGetUserUseCase = sync.OnceValue(func() *user.GetUserUseCase {
	return user.NewGetUserUseCase(getUserRepo())
})

var GetLoginUserUseCase = sync.OnceValue(func() *user.LoginUserUseCase {
	return user.NewLoginUserUseCase(getUserRepo())
})

var GetCreateUserUseCase = sync.OnceValue(func() *user.CreateUserUseCase {
	return user.NewCreateUserUseCase(getUserRepo())
})

var GetSendEmailUseCase = sync.OnceValue(func() *email.SendEmailUseCase {
	return email.NewSendEmailUseCase(getEmailService())
})