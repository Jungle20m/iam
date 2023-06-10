package tracer

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"log"
	"runtime"
	"time"
)

type tracer struct {
	otelProvider *tracesdk.TracerProvider

	environment string
	appName     string
	serviceName string
	serverName  string
	language    string
}

func NewTracer(opts ...Option) *tracer {
	t := &tracer{}

	for _, opt := range opts {
		opt(t)
	}

	return t
}

func (t *tracer) AttachJaegerProvider(url string) error {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return err
	}
	// Init resources
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(fmt.Sprintf("%s.%s", t.appName, t.serviceName)),
		attribute.String("environment", t.environment),
		attribute.String("service", t.serviceName),
		attribute.String("application", t.appName),
		attribute.String("language", t.language),
		attribute.String("server", t.serverName),
	)

	t.otelProvider = tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(res),
	)

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(t.otelProvider)

	return nil
}

func (t *tracer) Flush() {
	if t.otelProvider == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := t.otelProvider.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

func NewSpan(ctx context.Context) (context.Context, trace.Span) {
	// The call stack of the caller of the NewSpan function is obtained
	// value return include: pc, file, lineNo, ok
	pc, f, _, _ := runtime.Caller(1)

	ctx, span := otel.Tracer("").Start(ctx, runtime.FuncForPC(pc).Name())
	span.SetAttributes(attribute.Key("filepath").String(f))

	return ctx, span
}
