package rest

import (
	"github.com/gofiber/fiber/v2"
	"goboilerplate.com/src/di"
	"goboilerplate.com/src/rest/handlers"
	"goboilerplate.com/src/rest/handlers/user"
)

func RouteRegisterHandlers(app *fiber.App) {
	registerHealthRoutes(app)
	registerUserRoutes(app)
}

func registerHealthRoutes(app *fiber.App) {
	healthHandler := handlers.NewHealthHandler(di.GetHealthUseCase())
	app.Get("/health", healthHandler.CheckHealth)
}

func registerUserRoutes(app *fiber.App) {
	loginUserHandler := user.NewLoginUserHandler(di.GetLoginUserUseCase())
	app.Post("/users/login", loginUserHandler.LoginUser)
	getUserHandler := user.NewGetUserHandler(di.GetGetUserUseCase())
	app.Get("/users/:id", getUserHandler.GetUser)
	createUserHandler := user.NewCreateUserHandler(di.GetCreateUserUseCase())
	app.Post("/users", createUserHandler.CreateUser)
}