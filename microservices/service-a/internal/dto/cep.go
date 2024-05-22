package dto

import (
	"github.com/go-playground/validator/v10"
)

type CepInput struct {
	Cep string `json:"cep" validate:"cep"`
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
