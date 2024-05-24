// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package dependency_injection

import (
	"context"
	"github.com/leo12wb/Tracing-distribu-do-e-span/internal/infra/http_clients"
	"github.com/leo12wb/Tracing-distribu-do-e-span/internal/infra/web/webhandlers/get_temperature_handler"
	"github.com/leo12wb/Tracing-distribu-do-e-span/internal/usecase/get_temperature"
)

// Injectors from wire.go:

func NewTemperatureHandler(ctx *context.Context, client http_clients.ZipkinClientInterface) *get_temperature_handler.WebGetTemperatureHandler {
	useCase := NewTemperatureUseCase(client)
	webGetTemperatureHandler := get_temperature_handler.NewGetTemperatureHandler(useCase, client)
	return webGetTemperatureHandler
}

// wire.go:

func NewTemperatureUseCase(client http_clients.ZipkinClientInterface) get_temperature.UseCase {
	return get_temperature.NewUseCase(client)
}
