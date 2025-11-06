package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// ARRANGE
	testGreetings := make(map[string]string)
	testGreetings["abc"] = "123"

	// ACT
	router := getRouter(testGreetings)
	responseRecorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Error("Failed to create request")
	}

	// ASSERT
	router.ServeHTTP(responseRecorder, request)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "\"All good!\"", responseRecorder.Body.String())
}

type ResponceMap struct {
	Items []Greeting `json:"items"`
}

func TestGetGreetings(t *testing.T) {
	// ARRANGE
	testGreetings := make(map[string]string)
	testGreetings["abc"] = "123"
	testGreetings["def"] = "456"
	testGreetings["ghi"] = "789"

	// ACT
	router := getRouter(testGreetings)
	responce := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/greetings", nil)
	if err != nil {
		t.Error("Failed to create request")
	}
	router.ServeHTTP(responce, request)

	// ASSERT
	var responceMap ResponceMap

	err = json.Unmarshal(responce.Body.Bytes(), &responceMap)
	if err != nil {
		t.Error("Failed to unmarshal")
	}

	keyABCExists := false
	keyDEFExists := false
	keyGHIExists := false
	messageABCCorrect := false
	messageDEFCorrect := false
	messageGHICorrect := false

	for _, item := range responceMap.Items {
		if item.Id == "abc" {
			keyABCExists = true
			if item.Message == "123" {
				messageABCCorrect = true
			}
		}
		if item.Id == "def" {
			keyDEFExists = true
			if item.Message == "456" {
				messageDEFCorrect = true
			}
		}
		if item.Id == "ghi" {
			keyGHIExists = true
			if item.Message == "789" {
				messageGHICorrect = true
			}
		}
	}

	assert.Equal(t, http.StatusOK, responce.Code)
	assert.Equal(t, true, keyABCExists)
	assert.Equal(t, true, keyDEFExists)
	assert.Equal(t, true, keyGHIExists)
	assert.Equal(t, true, messageABCCorrect)
	assert.Equal(t, true, messageDEFCorrect)
	assert.Equal(t, true, messageGHICorrect)
	assert.Equal(t, 3, len(responceMap.Items))
}

func TestGetGreetingsEmptyGreeting(t *testing.T) {
	// ARRANGE
	testGreetings := make(map[string]string)
	testGreetings["emptyGreeting"] = ""
	router := getRouter(testGreetings)

	// ACT
	responce := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/greetings", nil)
	if err != nil {
		t.Error("Failed to create request")
	}
	router.ServeHTTP(responce, request)

	// ASSERT
	var responceMap ResponceMap

	err = json.Unmarshal(responce.Body.Bytes(), &responceMap)
	if err != nil {
		t.Error("Failed to unmarshal")
	}

	emptyGreetingIdExists := false
	emptyGreetingMessage := false

	for _, items := range responceMap.Items {
		if items.Id == "emptyGreeting" {
			emptyGreetingIdExists = true
			if items.Message == "" {
				emptyGreetingMessage = true
			}
		}
	}

	assert.Equal(t, http.StatusOK, responce.Code)
	assert.Equal(t, true, emptyGreetingIdExists)
	assert.Equal(t, true, emptyGreetingMessage)
	assert.Equal(t, 1, len(responceMap.Items))
}

func TestGetGreetingsNoGreetings(t *testing.T) {
	// ARRANGE
	testGreetings := make(map[string]string)
	router := getRouter(testGreetings)

	// ACT
	responce := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/greetings", nil)
	if err != nil {
		t.Error("Failed to create request")
	}
	router.ServeHTTP(responce, request)

	// ASSERT
	var responceMap ResponceMap

	err = json.Unmarshal(responce.Body.Bytes(), &responceMap)
	if err != nil {
		t.Error("Failed to unmarshal")
	}

	assert.Equal(t, http.StatusOK, responce.Code)
	assert.Equal(t, 0, len(responceMap.Items))
}
