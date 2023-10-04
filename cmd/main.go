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

var Calculator_url string

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

	fmt.Println("Starting server on port ", port)
	fmt.Println("Environment: ", env)
	fmt.Println("Calculator URL: ", Calculator_url)

	router := common.NewRouter()
	log.Fatal(http.ListenAndServe(port, router))
}
