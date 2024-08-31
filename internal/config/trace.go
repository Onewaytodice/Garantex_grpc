package config

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

type TraceConfig struct {
	Name string `env:"APP_NAME"`
	Host string `env:"TRACE_HOST"`
	Port string `env:"TRACE_PORT"`
}

func (t *TraceConfig) createEndpoint() string {
	return fmt.Sprintf("%s:%s", t.Host, t.Port)
}

func (t *TraceConfig) InitTracerProvider(ctx context.Context) (*trace.TracerProvider, error) {
	exporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint(t.createEndpoint()), otlptracehttp.WithInsecure())
	if err != nil {
		return nil, err
	}
	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(t.Name),
		),
	)
	if err != nil {
		return nil, err
	}
	provider := trace.NewTracerProvider(trace.WithBatcher(exporter), trace.WithResource(res))
	otel.SetTracerProvider(provider)
	return provider, nil
}
