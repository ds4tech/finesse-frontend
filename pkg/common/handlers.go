package common

import (
	"encoding/json"
	"fmt"
	"net/http"
    "os"
)

var dat map[string]interface{}

func HomeLink(w http.ResponseWriter, r *http.Request) {
	env := os.Getenv("ENV")
	welcome_string := "Welcome to simple Webserver!\n" + env
	fmt.Fprintf(w, welcome_string)
}

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server version 0.0.1\nmore information will be added soon.")
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	txt := r.URL.Query().Get("text")
	if len([]rune(txt)) < 1 {
		txt = "Text not provided in a query string."
	}

	jsonMap := map[string]string{"result": txt}
	jsonResult, _ := json.Marshal(jsonMap)
	fmt.Fprintf(w, string(jsonResult))
}
