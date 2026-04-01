package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"goboilerplate.com/config"
	"goboilerplate.com/src/pkg/otel"
	"goboilerplate.com/src/rest"
)

func main() {
	cfg := config.GetConfig()

	if cfg.YMLConfig.Telemetry.Enabled {
		slog.Info("Telemetry is enabled. Initializing OpenTelemetry...")
		shutdown, err := otel.InitOTel(cfg.YMLConfig.Telemetry.ServiceName)
		if err != nil {
			log.Fatalf("Failed to initialize OpenTelemetry: %v", err)
		}
		defer func() {
			if err := shutdown(context.Background()); err != nil {
				log.Printf("Failed to shutdown OpenTelemetry: %v", err)
			}
		}()
	} else {
		slog.Info("Telemetry is disabled. Skipping OpenTelemetry initialization.")
	}

	app := fiber.New()

	rest.RegisterMiddleware(app)

	rest.RouteRegisterHandlers(app)

	err := app.Listen(cfg.YMLConfig.Server.Port)
	if err != nil {
		slog.Error("Failed to start Fiber Server")
	}
}
