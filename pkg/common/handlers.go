package common

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

var dat map[string]interface{}

type CalcDetails struct {
	Num1 string
	Num2 string
	// Message string
}

func HomeLink(w http.ResponseWriter, r *http.Request) {
	// env := os.Getenv("ENV")
	// welcome_string := "Welcome to simple Webserver!\n" + env
	// fmt.Fprintf(w, welcome_string)

	tmpl := template.Must(template.ParseFiles("html/forms.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := CalcDetails{
		Num1: r.FormValue("num1"),
		Num2: r.FormValue("num2"),
		// Message: r.FormValue("message"),
	}

	// do something with details
	_ = details

	tmpl.Execute(w, struct{ Success bool }{true})
}

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server version 0.0.4\nmore information will be added soon.")
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
