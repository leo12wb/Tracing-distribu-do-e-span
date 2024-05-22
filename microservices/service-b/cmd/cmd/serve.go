/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"flag"
	"log"

	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-otel-go-expert-pos/microservices/service-b/pkg/adapter/http/rest"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-otel-go-expert-pos/microservices/service-b/pkg/adapter/otel"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-otel-go-expert-pos/microservices/service-b/pkg/adapter/validator"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		validator := validator.Initialize()

		url := flag.String("zipkin", "http://zipkin:9411/api/v2/spans", "zipkin url")
		flag.Parse()

		_, err := otel.Initialize(*url, "service-b")
		if err != nil {
			log.Fatal(err)
		}

		rest.Initialize(validator)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
