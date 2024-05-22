package usecase

import (
	"context"

	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-otel-go-expert-pos/microservices/service-b/internal/entity"
)

type usecase struct {
	cepHTTPClient     entity.CepHTTPClient
	weatherHTTPClient entity.WeatherHTTPClient
}

func NewWeatherUseCase(
	cepHTTPClient entity.CepHTTPClient,
	weatherHTTPClient entity.WeatherHTTPClient,
) entity.WeatherUseCase {
	return &usecase{
		cepHTTPClient:     cepHTTPClient,
		weatherHTTPClient: weatherHTTPClient,
	}
}

func (usecase usecase) Get(ctx context.Context, cep string) (*entity.Weather, error) {
	cepResponse, err := usecase.cepHTTPClient.Get(ctx, cep)

	if err != nil {
		return nil, err
	}

	weatherResponse, err := usecase.weatherHTTPClient.Get(ctx, cepResponse.CityName)

	if err != nil {
		return nil, err
	}

	weatherResponse.CalculateFarenheit()
	weatherResponse.CalculateKelvin()

	return weatherResponse, nil
}
