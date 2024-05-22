package di

import (
	"net/http"

	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-otel-go-expert-pos/microservices/service-b/internal/entity"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-otel-go-expert-pos/microservices/service-b/internal/infra/viacep"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-otel-go-expert-pos/microservices/service-b/internal/infra/weatherapi"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-otel-go-expert-pos/microservices/service-b/internal/infra/web"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-otel-go-expert-pos/microservices/service-b/internal/usecase"
	"github.com/go-playground/validator/v10"
)

func ConfigWebController(validator *validator.Validate) entity.WeatherController {
	httpClient := http.DefaultClient

	cepHttpClient := viacep.NewCepHTTPClient(httpClient)
	weatherHttpClient := weatherapi.NewWeatherHTTPClient(httpClient)
	weatherUseCase := usecase.NewWeatherUseCase(cepHttpClient, weatherHttpClient)
	weatherController := web.NewWebController(weatherUseCase, validator)

	return weatherController
}
