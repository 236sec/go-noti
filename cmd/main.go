package main

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"goboilerplate.com/config"
	"goboilerplate.com/src/rest"
)

func main() {
	cfg := config.GetConfig()

	app := fiber.New()

	rest.RegisterMiddleware(app)

	rest.RouteRegisterHandlers(app)

	err := app.Listen(cfg.YMLConfig.Server.Port)
	if err != nil {
		slog.Error("Failed to start Fiber Server")
	}
}
