package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	router := getRouter()
	responseRecorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Error("Failed to create request")
	}
	router.ServeHTTP(responseRecorder, request)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "\"All good!\"", responseRecorder.Body.String())
}

func TestGetGreetings(t *testing.T) {
	router := getRouter()
	responce := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/greetings", nil)
	if err != nil {
		t.Error("Failed to create request")
	}
	router.ServeHTTP(responce, request)
	assert.Equal(t, http.StatusOK, responce.Code)
	assert.Equal(t, "???", responce.Body.String())
}
