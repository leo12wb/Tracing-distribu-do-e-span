package entity

import (
	"context"
	"net/http"
)

type Weather struct {
	City  string  `json:"city"`
	TempC float32 `json:"temp_C"`
	TempF float32 `json:"temp_F"`
	TempK float32 `json:"temp_K"`
}

func (weather *Weather) CalculateFarenheit() {
	weather.TempF = weather.TempC*1.8 + 32
}

func (weather *Weather) CalculateKelvin() {
	weather.TempK = weather.TempC + 273
}

type WeatherHTTPClient interface {
	Get(context.Context, string) (*Weather, error)
}

type WeatherUseCase interface {
	Get(context.Context, string) (*Weather, error)
}

type WeatherController interface {
	Get(http.ResponseWriter, *http.Request)
}
