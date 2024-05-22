package viacep

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-otel-go-expert-pos/microservices/service-b/internal/dto"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-otel-go-expert-pos/microservices/service-b/internal/entity"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-otel-go-expert-pos/microservices/service-b/pkg/adapter/errorhandle"
	"go.opentelemetry.io/otel"
)

var (
	BASE_URL = "https://viacep.com.br"
)

type httpclient struct {
	client *http.Client
}

func NewCepHTTPClient(client *http.Client) entity.CepHTTPClient {
	return &httpclient{
		client: client,
	}
}

func (httpclient httpclient) Get(ctx context.Context, cep string) (*entity.Cep, error) {
	tr := otel.Tracer("microservice-trace")
	ctx, span := tr.Start(ctx, "get cep")
	defer span.End()

	var cepOutput dto.CepOutput
	url := fmt.Sprintf("%s/ws/%s/json/", BASE_URL, cep)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewReader(nil))

	if err != nil {

		return nil, err
	}

	defer request.Body.Close()

	response, err := httpclient.client.Do(request)

	if err != nil {

		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		err = json.NewDecoder(response.Body).Decode(&cepOutput)

		if err != nil {
			return nil, err
		}

		if cepOutput.Erro == "true" {
			return nil, errorhandle.ErrNotFound
		}

		return &entity.Cep{
			Cep:      cepOutput.Cep,
			CityName: cepOutput.Localidade,
		}, nil
	}

	return nil, errorhandle.ErrUnprocessableEntity
}
