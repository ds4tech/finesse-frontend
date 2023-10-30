package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ds4tech/finesse-frontend/pkg/common"
)

const (
	service     = "trace-demo"
	environment = "production"
	id          = 1
	port        = ":8080"
)

var Calculator_url, Otlp_endpoint string

func main() {
	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	fmt.Println("Registering TracerProvider")

	// tr := tp.Tracer("component-main")

	webserver()
}

func webserver() {
	// Use the global TracerProvider.
	// tr := otel.Tracer("component-webserver")
	// _, span := tr.Start(ctx, "webserver")
	// span.SetAttributes(attribute.Key("testset").String("value"))
	// defer span.End()

	//starting webserver
	env := os.Getenv("ENV")
	Calculator_url = os.Getenv("CALCULATOR_URL")
	Otlp_endpoint = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")

	fmt.Println("Starting server on port ", port)
	fmt.Println("Environment: ", env)
	fmt.Println("Calculator URL: ", Calculator_url)
	if Calculator_url == "" {
		log.Println("Calculator_url Environment Variable not set! Application will throw an error when submitting request.")
		// log.Fatal("Calculator_url Environment Variable not set! Application will throw an error when submitting request.")
	}

	router := common.NewRouter()
	log.Fatal(http.ListenAndServe(port, router))
}
