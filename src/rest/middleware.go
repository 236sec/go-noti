package rest

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"
	"goboilerplate.com/src/pkg/swagger"
	"goboilerplate.com/src/rest/response"
)

type contextKey string

const requestIDKey contextKey = "requestId"

func RegisterMiddleware(app *fiber.App) {
	app.Use(cors.New())

	swagger := swagger.GetSwagger()
	app.Use(swagger)
	
	app.Use(func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		ctx = context.WithValue(ctx, requestIDKey, c.Get("X-Request-ID"))
		c.SetUserContext(ctx)
		return c.Next()
	})

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger,
	}))

	app.Use(RecoveryMiddleware)
}

func RecoveryMiddleware(c *fiber.Ctx) error {
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