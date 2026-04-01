package swagger

import (
	swagger "github.com/gofiber/contrib/v3/swaggerui"
	"github.com/gofiber/fiber/v3"
	"goboilerplate.com/config"
)

func GetSwagger() fiber.Handler {
	cfg := config.GetConfig()
	swaggerConfig := swagger.Config{
		BasePath: cfg.YMLConfig.Swagger.BasePath,
		FilePath: cfg.YMLConfig.Swagger.FilePath,
		Path:     cfg.YMLConfig.Swagger.Path,
		Title:    cfg.YMLConfig.Swagger.Title,
	}

	return swagger.New(swaggerConfig)
}
