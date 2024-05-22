package entity_test

import (
	"testing"

	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-otel-go-expert-pos/microservices/service-b/internal/entity"
)

func TestCalculateFahrenheit(t *testing.T) {
	tests := []struct {
		name  string
		tempC float32
		wantF float32
	}{
		{"zero Celsius", 0, 32},
		{"boiling point of water", 100, 212},
		{"body temperature", 37, 98.6},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			weather := &entity.Weather{TempC: tc.tempC}
			weather.CalculateFarenheit()
			if weather.TempF != tc.wantF {
				t.Errorf("CalculateFarenheit() = %v, want %v", weather.TempF, tc.wantF)
			}
		})
	}
}

func TestCalculateKelvin(t *testing.T) {
	tests := []struct {
		name  string
		tempC float32
		wantK float32
	}{
		{"absolute zero", -273, 0},
		{"boiling point of water", 100, 373},
		{"body temperature", 37, 310},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			weather := &entity.Weather{TempC: tc.tempC}
			weather.CalculateKelvin()
			if weather.TempK != tc.wantK {
				t.Errorf("CalculateKelvin() = %v, want %v", weather.TempK, tc.wantK)
			}
		})
	}
}
