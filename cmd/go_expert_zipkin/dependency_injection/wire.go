//go:build wireinject
// +build wireinject

package dependency_injection

import (
	"context"

	"github.com/google/wire"
	"github.com/leo12wb/Tracing-distribu-do-e-span/internal/infra/http_clients"
	"github.com/leo12wb/Tracing-distribu-do-e-span/internal/infra/web/webhandlers/get_temperature_handler"
	"github.com/leo12wb/Tracing-distribu-do-e-span/internal/usecase/get_temperature"
)

func NewTemperatureUseCase(client http_clients.ZipkinClientInterface) get_temperature.UseCase {
	return get_temperature.NewUseCase(client)
}
func NewTemperatureHandler(ctx *context.Context, client http_clients.ZipkinClientInterface) *get_temperature_handler.WebGetTemperatureHandler {
	wire.Build(NewTemperatureUseCase, get_temperature_handler.NewGetTemperatureHandler)
	return &get_temperature_handler.WebGetTemperatureHandler{}
}

/*
var setSampleRepositoryDependency = wire.NewSet(
	database.SampleRepository,
	wire.Bind(new(entity.SampleRepositoryInterface), new(*database.SampleRepository)),
)

func NewListAllOrdersUseCase(db *sql.DB) *usecase.MyUseCase {
	wire.Build(
		setSampleRepositoryDependency,
		usecase.NewUseCase,
	)
	return &usecase.MyUseCase{}
}
*/
