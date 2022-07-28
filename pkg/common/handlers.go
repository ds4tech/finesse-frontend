package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var dat map[string]interface{}

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to commonulator!")
}

func SumHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter integer numbers")
	}
	json.Unmarshal(reqBody, &dat)
	n1 := dat["num1"].(string)
	n2 := dat["num2"].(string)

	i1, err := strconv.ParseInt(n1, 10, 64)
	i2, err := strconv.ParseInt(n2, 10, 64)
	result := Sum(i1, i2)
	if err == nil {
		jsonMap := map[string]int64{"result": result}
		jsonResult, _ := json.Marshal(jsonMap)
		fmt.Fprintf(w, string(jsonResult))
		fmt.Printf("Sum of %v and %v is: %v\n", n1, n2, result)
	}
}

func SqrtHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter integer number")
	}
	json.Unmarshal(reqBody, &dat)
	num := dat["number"].(string)

	i, err := strconv.ParseFloat(num, 64)
	result := Sqrt(i)
	if err == nil {
		jsonMap := map[string]float64{"result": result}
		jsonResult, _ := json.Marshal(jsonMap)
		fmt.Fprintf(w, string(jsonResult))
		fmt.Printf("Sqrt of %v is: %v\n", num, result)
	}
}

func LogHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter integer number")
	}
	json.Unmarshal(reqBody, &dat)
	num := dat["number"].(string)

	i, err := strconv.ParseFloat(num, 64)
	result := Log(i)
	if err == nil {
		jsonMap := map[string]float64{"result": result}
		jsonResult, _ := json.Marshal(jsonMap)
		fmt.Fprintf(w, string(jsonResult))
		fmt.Printf("Log of %v is: %v\n", num, result)
	}
}

func FactorialHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter integer number")
	}
	json.Unmarshal(reqBody, &dat)
	num := dat["number"].(string)

	i, err := strconv.ParseUint(num, 10, 64)
	result := Factorial(i)
	if err == nil {
		jsonMap := map[string]uint64{"result": result}
		jsonResult, _ := json.Marshal(jsonMap)
		fmt.Fprintf(w, string(jsonResult))
		fmt.Printf("Factorial of %v is: %v\n", num, result)
	}
}

func IsPrimeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter integer number")
	}

	if err := json.Unmarshal(reqBody, &dat); err != nil {
		w.Header().Set("content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		panic(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	num := dat["number"].(string)

	i, err := strconv.Atoi(num)
	result := IsPrime(i)
	if err == nil {
		jsonMap := map[string]bool{"isPrime": result}
		jsonResult, _ := json.Marshal(jsonMap)
		fmt.Fprintf(w, string(jsonResult))
		fmt.Printf("%v is prime: %v\n", num, result)
	}
}
