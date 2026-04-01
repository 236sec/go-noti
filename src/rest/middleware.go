package rest

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"goboilerplate.com/src/pkg/swagger"
	"goboilerplate.com/src/rest/response"
)

type contextKey string

const requestIDKey contextKey = "requestId"

func RegisterMiddleware(app *fiber.App) {
	app.Use(cors.New())

	swagger := swagger.GetSwagger()
	app.Use(swagger)

	app.Use(func(c fiber.Ctx) error {
		ctx := c.Context()
		ctx = context.WithValue(ctx, requestIDKey, c.Get("X-Request-ID"))
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

// fiber:context-methods migrated
