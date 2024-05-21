package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Request struct {
	CEP string `json:"cep"`
}

func main() {
	initTracer()
	http.HandleFunc("/cep", handleCEP)
	http.ListenAndServe(":8080", nil)
}

func handleCEP(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || !isValidCEP(req.CEP) {
		http.Error(w, `{"message":"invalid zipcode"}`, http.StatusUnprocessableEntity)
		return
	}

	// Tracing
	tracer := otel.Tracer("service-A")
	_, span := tracer.Start(r.Context(), "handleCEP")
	defer span.End()

	// Call Service B
	resp, err := http.Post("http://service-b:8081/weather", "application/json", strings.NewReader(fmt.Sprintf(`{"cep":"%s"}`, req.CEP)))
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, `{"message":"invalid zipcode"}`, http.StatusUnprocessableEntity)
		return
	}
	defer resp.Body.Close()

	// Copy the response from Service B
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp.Body); err != nil {
		http.Error(w, `{"message":"internal error"}`, http.StatusInternalServerError)
	}
}

func isValidCEP(cep string) bool {
	return len(cep) == 8
}

func initTracer() {
	exporter, err := zipkin.New(
		"http://zipkin:9411/api/v2/spans",
		zipkin.WithLogger(nil),
	)
	if err != nil {
		panic(err)
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String("service-A"),
		)),
	)
	otel.SetTracerProvider(tp)
}
