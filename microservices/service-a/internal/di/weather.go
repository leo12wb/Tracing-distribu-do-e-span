package di

import (
	"net/http"

	"github.com/leo12wb/Tracing-distribu-do-e-span/microservices/service-a/internal/entity"
	"github.com/leo12wb/Tracing-distribu-do-e-span/microservices/service-a/internal/infra/weatherapi"
	"github.com/leo12wb/Tracing-distribu-do-e-span/microservices/service-a/internal/infra/web"
	"github.com/leo12wb/Tracing-distribu-do-e-span/microservices/service-a/internal/usecase"
	"github.com/go-playground/validator/v10"
)

func ConfigWebController(validator *validator.Validate) entity.WeatherController {
	httpClient := http.DefaultClient

	weatherHttpClient := weatherapi.NewWeatherHTTPClient(httpClient)
	weatherUseCase := usecase.NewWeatherUseCase(weatherHttpClient)
	weatherController := web.NewWebController(weatherUseCase, validator)

	return weatherController
}
