package dto

import (
	"github.com/go-playground/validator/v10"
)

type CepInput struct {
	Cep string `json:"cep" validate:"cep"`
}

type CepOutput struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	DDD         string `json:"ddd"`
	Siafi       string `json:"siafi"`
	Erro        string `json:"erro"`
}

func FromQueryStringRequestToCep(cep string, validate *validator.Validate) (*CepInput, error) {
	cepInput := CepInput{
		Cep: cep,
	}

	if err := validate.Struct(cepInput); err != nil {
		return nil, err
	}

	return &cepInput, nil
}
