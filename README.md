# Jenkins Pipeline OpenTelemetry Exporter

> A hands-on DevOps portfolio project demonstrating OpenTelemetry instrumentation for CI/CD pipelines. Built for learning and GSoC preparation.

Proof-of-concept: Instrument a CI/CD pipeline with OpenTelemetry traces for observability.

**Relevant for:** Jenkins GSoC 2026 "Use OpenTelemetry for Jenkins Jobs on ci.jenkins.io"

## Concept

This project demonstrates how to add OpenTelemetry instrumentation to a pipeline:
- Trace pipeline stages (build, test, deploy)
- Export spans to Jaeger/OTLP collector
- Visualize pipeline execution

## Structure

```
├── pipeline/          # Sample pipeline with OTel instrumentation
├── collector/         # OTLP collector config
└── docker-compose.yml # Run locally with Jaeger
```

## Quick Start

```bash
docker compose up -d
# Jaeger UI: http://localhost:16686
```

## Tech Stack

- OpenTelemetry Go SDK
- Jaeger (trace backend)
- Docker Compose

## License

MIT
