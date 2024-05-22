package entity

import (
	"context"
)

type Cep struct {
	Cep      string `json:"cep"`
	CityName string `json:"cityName"`
}

type CepHTTPClient interface {
	Get(context.Context, string) (*Cep, error)
}
