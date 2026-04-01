package handlers

import (
	"github.com/gofiber/fiber/v3"
	"goboilerplate.com/src/usecases"
)

type HealthHandler struct {
	healthUseCase usecases.IHealthUseCase
}

func NewHealthHandler(healthUseCase usecases.IHealthUseCase) *HealthHandler {
	return &HealthHandler{
		healthUseCase: healthUseCase,
	}
}

func (h *HealthHandler) CheckHealth(c fiber.Ctx) error {
	ctx := c.Context()
	if err := h.healthUseCase.Apply(ctx); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Service Unhealthy",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Service Healthy",
	})
}
