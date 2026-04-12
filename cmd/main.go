package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"goboilerplate.com/config"
	"goboilerplate.com/src/consumers"
	"goboilerplate.com/src/di"
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

	if cfg.YMLConfig.Broker.Enabled {
		slog.Info("Kafka broker is enabled. Testing connection...")
		producer := di.GetKafkaProducer()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		if err := producer.Ping(ctx); err != nil {
			cancel()
			log.Fatalf("Failed to connect to Kafka: %v", err)
		}
		cancel()
		slog.Info("Successfully connected to Kafka.")
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if cfg.YMLConfig.Broker.Enabled {
		consumers.RegisterConsumers(ctx)
	}

	app := fiber.New()

	rest.RegisterMiddleware(app)

	rest.RouteRegisterHandlers(app)

	go func() {
		if err := app.Listen(cfg.YMLConfig.Server.Port); err != nil {
			slog.Error("Failed to start Fiber Server", "error", err)
		}
	}()

	<-ctx.Done()
	slog.Info("Shutdown signal received. Commencing graceful shutdown...")
	stop()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(shutdownCtx); err != nil {
		slog.Error("Fiber Shutdown Error", "error", err)
	}

	slog.Info("Server exited properly")
}
