package main

import (
	"fmt"
	"log"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
	"go.opentelemetry.io/otel/trace"
)

func main() {
    // Create Zipkin exporter
    endpoint := "http://localhost:9411/api/v2/spans"
    exporter, err := zipkin.NewRawExporter(endpoint)
    if err != nil {
        log.Fatal(err)
    }

    // Create trace provider with Zipkin exporter
    tp := trace.NewTracerProvider(
        trace.WithConfig(trace.Config{DefaultSampler: trace.AlwaysSample()}),
        trace.WithBatcher(exporter),
        trace.WithResource(resource.NewWithAttributes(
            semconv.ServiceNameKey.String("ServiceA"),
        )),
    )

    // Initialize the global trace provider
    otel.SetTracerProvider(tp)

    // Initialize the propagators
    propagators := propagation.TraceContext{}
    otel.SetTextMapPropagator(propagators)

    // Handle incoming requests
    http.HandleFunc("/process-cep", ProcessCEP)
    http.ListenAndServe(":8080", nil)
}

func ProcessCEP(w http.ResponseWriter, r *http.Request) {
    // Start a span
    ctx, span := otel.GetTracerProvider().Tracer("ServiceA").Start(r.Context(), "ProcessCEP")
    defer span.End()

    // Here you can add tracing to specific operations within this handler
    // For example, you can create child spans for each operation

    // Example:
    // ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "OperationName")
    // defer span.End()

    // Your existing handler logic goes here

    fmt.Fprintf(w, "CEP received and forwarded to Service B")
}
