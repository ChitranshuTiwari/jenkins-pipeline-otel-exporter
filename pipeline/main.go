package main

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func main() {
	ctx := context.Background()

	// OTLP exporter (connects to Jaeger)
	exporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint("localhost:4317"),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("jenkins-pipeline-demo"),
		)),
	)
	defer tp.Shutdown(ctx)
	otel.SetTracerProvider(tp)

	tracer := otel.Tracer("pipeline")
	ctx, span := tracer.Start(ctx, "pipeline-run")
	defer span.End()

	// Simulate pipeline stages
	_, buildSpan := tracer.Start(ctx, "build")
	time.Sleep(100 * time.Millisecond)
	buildSpan.End()

	_, testSpan := tracer.Start(ctx, "test")
	time.Sleep(150 * time.Millisecond)
	testSpan.End()

	_, deploySpan := tracer.Start(ctx, "deploy")
	time.Sleep(80 * time.Millisecond)
	deploySpan.End()

	log.Println("Pipeline complete - check Jaeger at http://localhost:16686")
}
