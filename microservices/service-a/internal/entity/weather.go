package entity

import (
	"context"
	"net/http"

	"github.com/leo12wb/Tracing-distribu-do-e-span/microservices/service-a/internal/dto"
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
