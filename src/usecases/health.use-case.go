package usecases

import (
	"context"

	"go.opentelemetry.io/otel"
)

var healthTracer = otel.Tracer("usecase.health")

type IHealthUseCase interface {
	Apply(ctx context.Context) error
}

type HealthUseCase struct{}

func (u *HealthUseCase) Apply(ctx context.Context) error {
	_, span := healthTracer.Start(ctx, "HealthUseCase.CheckHealth")
	defer span.End()
	return nil
}

func NewHealthUseCase() *HealthUseCase {
	return &HealthUseCase{}
}
