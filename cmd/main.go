package main

import (
	"fmt"
	"log"
	"net/http"
    "os"

	"github.com/ds4tech/covantis-sre/pkg/common"
)

func main() {
	port := ":9090"
	env := os.Getenv("ENV")
	fmt.Println("Starting server on port ", port)
	fmt.Println("Environment: ", env)

	router := common.NewRouter()
	log.Fatal(http.ListenAndServe(port, router))
}
