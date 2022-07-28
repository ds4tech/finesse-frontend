package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ds4tech/covantis-sre/pkg/common"
)

func main() {
	fmt.Println("Starting server...\n")

	router := common.NewRouter()
	log.Fatal(http.ListenAndServe(":9090", router))
}
