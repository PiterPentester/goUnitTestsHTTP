package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	res := Add(3, 7)
	assert.NotNil(t, res, "The `res` should not be nil")
	assert.Equal(t, 10, res, "Expecting 10")
}

func TestSubstract(t *testing.T) {
	res := Substract(3, 7)
	assert.NotNil(t, res, "The `res` should not be nil")
	assert.Equal(t, -4, res, "Expecting -4")
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", RootEndpoint).Methods("GET")
	return r
}

func TestRootEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Expecting 200")
	assert.Equal(t, "Hello World", response.Body.String(), "Incorrect body found!")
}
