package otel

import (
	"context"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

// InitOTel sets up tracing and metrics pushing to the OTel Collector
func InitOTel(serviceName string) (func(context.Context) error, error) {
	ctx := context.Background()
	instanceID, err := os.Hostname()
	if err != nil || instanceID == "" {
		// Fallback if Hostname fails for some reason
		instanceID = "unknown-instance"
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
			semconv.ServiceInstanceID(instanceID),
		),
	)
	if err != nil {
		return nil, err
	}

	// 1. Setup Tracing Exporter (Push to Collector via gRPC)
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tracerProvider)

	// 2. Setup Metrics Exporter (Push to Collector via gRPC)
	metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExporter, sdkmetric.WithInterval(3*time.Second))),
		sdkmetric.WithResource(res),
	)
	otel.SetMeterProvider(meterProvider)

	// 3. Setup W3C Trace Context Propagation (The "Trace-ID" headers)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Return a shutdown function to clean up on exit
	shutdown := func(ctx context.Context) error {
		_ = tracerProvider.Shutdown(ctx)
		return meterProvider.Shutdown(ctx)
	}
	return shutdown, nil
}
