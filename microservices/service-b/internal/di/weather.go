package di

import (
	"net/http"

	"github.com/leo12wb/Tracing-distribu-do-e-span/microservices/service-b/internal/entity"
	"github.com/leo12wb/Tracing-distribu-do-e-span/microservices/service-b/internal/infra/viacep"
	"github.com/leo12wb/Tracing-distribu-do-e-span/microservices/service-b/internal/infra/weatherapi"
	"github.com/leo12wb/Tracing-distribu-do-e-span/microservices/service-b/internal/infra/web"
	"github.com/leo12wb/Tracing-distribu-do-e-span/microservices/service-b/internal/usecase"
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
