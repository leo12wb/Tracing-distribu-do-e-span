package weatherapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"errors"

	"github.com/leo12wb/Tracing-distribu-do-e-span/microservices/service-a/internal/dto"
	"github.com/leo12wb/Tracing-distribu-do-e-span/microservices/service-a/internal/entity"
	"github.com/leo12wb/Tracing-distribu-do-e-span/microservices/service-a/pkg/adapter/errorhandle"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

var (
	BASE_URL = "http://service-b:8081/cep"
)

type httpclient struct {
	client *http.Client
}

func NewWeatherHTTPClient(client *http.Client) entity.WeatherHTTPClient {
	return &httpclient{
		client: client,
	}
}

func (httpclient httpclient) Get(ctx context.Context, cep string) (*dto.WeatherOutput, error) {
	tr := otel.Tracer("microservice-trace")
	ctx, span := tr.Start(ctx, "get weather from service b")
	defer span.End()

	var weatherOutput dto.WeatherOutput

	url := fmt.Sprintf("%s/%s", BASE_URL, cep)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewReader(nil))

	if err != nil {
		return nil, err
	}

	defer request.Body.Close()

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(request.Header))
	response, err := httpclient.client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode == 404 {
		//return nil, errorhandle.ErrNotFound
		return nil, errors.New("404 can not find zipcode.")
	}

	if response.StatusCode == 422 {
		return nil, errorhandle.ErrUnprocessableEntity
	}

	err = json.NewDecoder(response.Body).Decode(&weatherOutput)

	if err != nil {
		return nil, err
	}

	return &weatherOutput, nil
}
