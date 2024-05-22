/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"flag"
	"log"

	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-otel-go-expert-pos/microservices/service-a/pkg/adapter/http/rest"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-otel-go-expert-pos/microservices/service-a/pkg/adapter/otel"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-otel-go-expert-pos/microservices/service-a/pkg/adapter/validator"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		//
		// url := flag.String("zipkin", "http://zipkin:9411/api/v2/spans", "zipkin url")
		// otel.Initialize(*url)
		//

		// ctx := context.Background()

		// tr := otel1.GetTracerProvider().Tracer("component-main")
		// ctx, span := tr.Start(ctx, "foo", trace.WithSpanKind(trace.SpanKindServer))
		// <-time.After(6 * time.Millisecond)
		// bar(ctx)
		// <-time.After(6 * time.Millisecond)
		// span.End()

		validator := validator.Initialize()

		url := flag.String("zipkin", "http://zipkin:9411/api/v2/spans", "zipkin url")
		flag.Parse()

		_, err := otel.Initialize(*url, "service-a")
		if err != nil {
			log.Fatal(err)
		}

		rest.Initialize(validator)
	},
}

// func bar(ctx context.Context) {
// 	tr := otel1.GetTracerProvider().Tracer("component-bar")
// 	_, span := tr.Start(ctx, "bar")
// 	<-time.After(6 * time.Millisecond)
// 	span.End()
// }

func init() {
	rootCmd.AddCommand(serveCmd)
}
