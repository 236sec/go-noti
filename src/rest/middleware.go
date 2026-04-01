package rest

import (
	"context"
	"log"

	fiberotel "github.com/gofiber/contrib/v3/otel"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/google/uuid"
	"goboilerplate.com/config"
	ctxkey "goboilerplate.com/src/pkg/ctx"
	"goboilerplate.com/src/pkg/swagger"
	"goboilerplate.com/src/rest/response"
)

func RegisterMiddleware(app *fiber.App) {
	cfg := config.GetConfig()
	if cfg.YMLConfig.Telemetry.Enabled {
		app.Use(fiberotel.Middleware())
	}
	app.Use(cors.New())

	swagger := swagger.GetSwagger()
	app.Use(swagger)

	app.Use(func(c fiber.Ctx) error {
		ctx := c.Context()
		ctx = context.WithValue(ctx, ctxkey.RequestID{}, GetRequestID(c))
		c.SetContext(ctx)
		return c.Next()
	})

	app.Use(logger.New(logger.Config{
		CustomTags: map[string]logger.LogFunc{
			"requestid": func(output logger.Buffer, c fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				return output.WriteString(c.Get("X-Request-ID"))
			},
			"requestbody": func(output logger.Buffer, c fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				return output.Write(c.Body())
			},
		},
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\nRequestID: ${requestid}\nRequestBody: ${requestbody}\nResponse: ${resBody}\n",
	}))

	app.Use(RecoveryMiddleware)
}

func RecoveryMiddleware(c fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic recovered: %v\n", r)

			res := response.Responses[response.InternalServerErrorResponse]
			if err := c.Status(res.HttpStatus).JSON(res); err != nil {
				log.Printf("Failed to write recovery response: %v\n", err)
			}
		}
	}()
	return c.Next()
}

func GetRequestID(c fiber.Ctx) string {
	requestID := c.Get("X-Request-ID")
	if requestID == "" {
		requestID = uuid.New().String()
		c.Set("X-Request-ID", requestID)
	}
	return requestID
}
