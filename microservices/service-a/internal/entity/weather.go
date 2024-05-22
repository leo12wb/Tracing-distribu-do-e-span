package entity

import (
	"context"
	"net/http"

	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-otel-go-expert-pos/microservices/service-a/internal/dto"
)

type WeatherHTTPClient interface {
	Get(context.Context, string) (*dto.WeatherOutput, error)
}

type WeatherUseCase interface {
	Get(context.Context, string) (*dto.WeatherOutput, error)
}

type WeatherController interface {
	Get(http.ResponseWriter, *http.Request)
}
