package rest

import (
	"log"
	"net/http"

	"github.com/leo12wb/Tracing-distribu-do-e-span/microservices/service-b/internal/di"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"

	httpSwagger "github.com/swaggo/http-swagger/v2"

	docs "github.com/leo12wb/Tracing-distribu-do-e-span/microservices/service-b/pkg/adapter/http/rest/docs"
)

// @title Desafio Sistema de Temperatura por Cep Go Expert API Docs
// @version 1.0.0
// @description
// @termsOfService
// @contact.name Vinícius Boscardin
// @contact.email boscardinvinicius@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /

func Initialize(validator *validator.Validate) {
	docs.SwaggerInfo.BasePath = "/"
	webController := di.ConfigWebController(validator)

	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	r.Get("/cep/{cep}", webController.Get)

	log.Println("Rodando na porta: 8081")
	http.ListenAndServe(":8081", r)
}
