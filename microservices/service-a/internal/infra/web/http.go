package web

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/leo12wb/Tracing-distribu-do-e-span/microservices/service-a/internal/dto"
	"github.com/leo12wb/Tracing-distribu-do-e-span/microservices/service-a/internal/entity"
	"github.com/leo12wb/Tracing-distribu-do-e-span/microservices/service-a/pkg/adapter/errorhandle"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type controller struct {
	usecase   entity.WeatherUseCase
	validator *validator.Validate
}

func NewWebController(usecase entity.WeatherUseCase, validator *validator.Validate) entity.WeatherController {
	return &controller{
		usecase:   usecase,
		validator: validator,
	}
}

// Get goDoc
// @Summary Get temperature
// @Description Get temperature in celcius, kelvin and fahrenheit
// @Tags cep
// @Accept  json
// @Produce  json
// @Param cep path string true "cep"
// @Success 200 {object} dto.WeatherOutput
// @Failure 404 {object} errorhandle.Response
// @Failure 422 {object} errorhandle.Response
// @Router /cep/{cep} [get]
func (controller controller) Get(response http.ResponseWriter, request *http.Request) {
	carrier := propagation.HeaderCarrier(request.Header)
	ctx := request.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	tr := otel.Tracer("microservice-trace")
	ctx, span := tr.Start(ctx, "get weather")
	defer span.End()

	time.Sleep(time.Millisecond * 1000)

	cep, err := dto.FromQueryStringRequestToCep(chi.URLParam(request, "cep"), controller.validator)

	if err != nil {
		statusCode, message := errorhandle.Handle(errorhandle.ErrUnprocessableEntity)
		response.WriteHeader(statusCode)
		json.NewEncoder(response).Encode(message)

		return
	}

	weatherOutput, err := controller.usecase.Get(ctx, cep.Cep)

	if err != nil {
		statusCode, message := errorhandle.Handle(err)
		response.WriteHeader(statusCode)
		json.NewEncoder(response).Encode(message)

		return
	}

	json.NewEncoder(response).Encode(weatherOutput)
}
