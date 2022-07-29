package common

import (
	"fmt"
	"net/http"
)

var dat map[string]interface{}

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to simple Webserver!")
}

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here the logs come:")
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	txt := r.URL.Query().Get("text")
	if len([]rune(txt)) < 1 {
		txt = "Text not provided in a query string."
	}
	fmt.Println("text =>", txt)
	fmt.Fprintf(w, txt)
}
