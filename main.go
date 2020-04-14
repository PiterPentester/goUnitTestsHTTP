package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Add - just add two int
func Add(val1, val2 int) int {
	return val1 + val2
}

// Substract - just substract two int
func Substract(val1, val2 int) int {
	return val1 - val2
}

// RootEndpoint - write header=200 & "Hello World" to page
func RootEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}

func main() {
	num1 := 23
	num2 := 27
	fmt.Printf("%d + %d = %d \n", num1, num2, Add(num1, num2))
	fmt.Printf("%d - %d = %d \n", num1, num2, Substract(num1, num2))

	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":1234", router))
}
