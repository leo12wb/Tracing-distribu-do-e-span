package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Request struct {
	CEP string `json:"cep"`
}

type Response struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type ViaCEPResponse struct {
	Localidade string `json:"localidade"`
}

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func main() {
	initTracer()
	http.HandleFunc("/weather", handleWeather)
	http.ListenAndServe(":8081", nil)
}

func handleWeather(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || !isValidCEP(req.CEP) {
		http.Error(w, `{"message":"invalid zipcode"}`, http.StatusUnprocessableEntity)
		return
	}

	// Tracing
	tracer := otel.Tracer("service-B")
	ctx, span := tracer.Start(r.Context(), "handleWeather")
	defer span.End()

	city, err := getCity(ctx, req.CEP)
	if err != nil {
		http.Error(w, `{"message":"can not find zipcode"}`, http.StatusNotFound)
		return
	}

	tempC, err := getTemperature(ctx, city)
	if err != nil {
		http.Error(w, `{"message":"internal error"}`, http.StatusInternalServerError)
		return
	}

	resp := Response{
		City:  city,
		TempC: tempC,
		TempF: tempC*1.8 + 32,
		TempK: tempC + 273.15,
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func getCity(ctx context.Context, cep string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var viaCEPResponse ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&viaCEPResponse); err != nil {
		return "", err
	}

	return viaCEPResponse.Localidade, nil
}

func getTemperature(ctx context.Context, city string) (float64, error) {
	resp, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=YOUR_API_KEY&q=%s", city))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var weatherAPIResponse WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherAPIResponse); err != nil {
		return 0, err
	}

	return weatherAPIResponse.Current.TempC, nil
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
			semconv.ServiceNameKey.String("service-B"),
		)),
	)
	otel.SetTracerProvider(tp)
}
