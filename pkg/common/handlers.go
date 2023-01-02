package common

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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

	// send request to calc container (docker run)
	url := "http://localhost/v1/sum"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf("{\"num1\":\"%v\",\"num2\":\"%v\"}\n", details.Num1, details.Num2))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	json.Unmarshal(body, &dat)
	// end

	tmpl.Execute(w, struct{ Success bool }{true})

	result := strconv.Itoa(int(dat["result"].(float64)))
	fmt.Fprintf(w, "Result is: ")
	fmt.Fprintf(w, result)
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
