module github.com/ChitranshuTiwari/jenkins-pipeline-otel-exporter/pipeline

go 1.21

require (
	go.opentelemetry.io/otel v1.21.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.21.0
	go.opentelemetry.io/otel/sdk v1.21.0
	go.opentelemetry.io/otel/semconv/v1.21.0 v1.21.0
	go.opentelemetry.io/otel/trace v1.21.0
)
