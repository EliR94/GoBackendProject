package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"project/services/uuid"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// ARRANGE
	testGreetings := make(map[string]string)
	testGreetings["abc"] = "123"

	// ACT
	fakeUUIDService := uuid.FakeUUIDService{}
	router := getRouter(testGreetings, &fakeUUIDService)
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

func TestGetGreetings(t *testing.T) {
	// ARRANGE
	testGreetings := make(map[string]string)
	testGreetings["abc"] = "123"
	testGreetings["def"] = "456"
	testGreetings["ghi"] = "789"

	// ACT
	fakeUUIDService := uuid.FakeUUIDService{}
	router := getRouter(testGreetings, &fakeUUIDService)
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
	fakeUUIDService := uuid.FakeUUIDService{}
	router := getRouter(testGreetings, &fakeUUIDService)

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
	fakeUUIDService := uuid.FakeUUIDService{}
	router := getRouter(testGreetings, &fakeUUIDService)

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

func TestPostGreetings(t *testing.T) {
	// ARRANGE
	var fakeUUID = "12345678-9012-3456-7890-123456789012"
	testGreetings := make(map[string]string)
	fakeUUIDService := uuid.FakeUUIDService{}
	fakeUUIDService.StoreFakeUUID(fakeUUID)
	router := getRouter(testGreetings, &fakeUUIDService)

	// assert the greeting is posted
	// ACT
	responce := httptest.NewRecorder()
	var postBody PostRequest
	postBody.Message = "Hello World"
	jsonBody, err := json.Marshal(postBody)
	if err != nil {
		t.Error("Failed to marshal")
	}
	request, err := http.NewRequest("POST", "/greeting", bytes.NewReader(jsonBody))
	if err != nil {
		t.Error("Failed to create request")
	}
	router.ServeHTTP(responce, request)

	// ASSERT
	var postResponse Greeting
	err = json.Unmarshal(responce.Body.Bytes(), &postResponse)
	if err != nil {
		t.Error("Failed to unmarshal")
	}

	assert.Equal(t, http.StatusCreated, responce.Code)
	assert.Equal(t, fakeUUID, postResponse.Id)
	assert.Equal(t, postBody.Message, postResponse.Message)

	// now assert the greeting persists in the system
	// ACT
	getResponce := httptest.NewRecorder()
	getRequest, err := http.NewRequest("GET", "/greetings", nil)
	if err != nil {
		t.Error("Failed to create request")
	}
	router.ServeHTTP(getResponce, getRequest)

	// ASSERT
	var responceMap ResponceMap

	err = json.Unmarshal(getResponce.Body.Bytes(), &responceMap)
	if err != nil {
		t.Error("Failed to unmarshal")
	}

	correctGreetingId := false
	correctGreetingMessage := false

	for _, items := range responceMap.Items {
		if items.Id == postResponse.Id {
			correctGreetingId = true
			if items.Message == postResponse.Message {
				correctGreetingMessage = true
			}
		}
	}

	assert.Equal(t, http.StatusOK, getResponce.Code)
	assert.Equal(t, true, correctGreetingId)
	assert.Equal(t, true, correctGreetingMessage)
	assert.Equal(t, 1, len(responceMap.Items))
}

func TestPostGreetingsBadRequest(t *testing.T) {
	// ARRANGE
	testGreetings := make(map[string]string)

	fakeUUIDService := uuid.FakeUUIDService{}
	router := getRouter(testGreetings, &fakeUUIDService)

	// assert the greeting is posted
	// ACT
	responce := httptest.NewRecorder()
	request, err := http.NewRequest("POST", "/greeting", strings.NewReader(`{"thisPayload": "hasTheWrongData"}`))
	if err != nil {
		t.Error("Failed to create request")
	}
	router.ServeHTTP(responce, request)

	// ASSERT
	var errorResponse ErrorResponse
	err = json.Unmarshal(responce.Body.Bytes(), &errorResponse)
	if err != nil {
		t.Error("Failed to unmarshal")
	}

	assert.Equal(t, http.StatusBadRequest, responce.Code)
	assert.Equal(t, `Key: 'PostRequest.Message' Error:Field validation for 'Message' failed on the 'required' tag`, errorResponse.Error)
}

func TestPostGreetingsEmptyGreeting(t *testing.T) {
	// ARRANGE
	testGreetings := make(map[string]string)
	fakeUUIDService := uuid.FakeUUIDService{}
	router := getRouter(testGreetings, &fakeUUIDService)

	// ACT
	responce := httptest.NewRecorder()
	var postBody PostRequest
	postBody.Message = ""
	jsonBody, err := json.Marshal(postBody)
	if err != nil {
		t.Error("Failed to marshal")
	}
	request, err := http.NewRequest("POST", "/greeting", bytes.NewReader(jsonBody))
	if err != nil {
		t.Error("Failed to create request")
	}
	router.ServeHTTP(responce, request)

	// ASSERT
	var errorResponse ErrorResponse
	err = json.Unmarshal(responce.Body.Bytes(), &errorResponse)
	if err != nil {
		t.Error("Failed to unmarshal")
	}

	assert.Equal(t, http.StatusBadRequest, responce.Code)
	assert.Equal(t, `Key: 'PostRequest.Message' Error:Field validation for 'Message' failed on the 'required' tag`, errorResponse.Error)
}

func TestGetGreeting(t *testing.T) {
	// ARRANGE
	testGreetings := make(map[string]string)
	expectedId := "12345678-9012-3456-7890-123456789012"
	expectedMessage := "Hello World"
	testGreetings[expectedId] = expectedMessage

	// ACT
	fakeUUIDService := uuid.FakeUUIDService{}
	router := getRouter(testGreetings, &fakeUUIDService)
	responce := httptest.NewRecorder()
	request, err := http.NewRequest("GET", fmt.Sprintf("/greeting/%s", expectedId), nil)
	if err != nil {
		t.Error("Failed to create request")
	}
	router.ServeHTTP(responce, request)

	// ASSERT
	var getResponse Greeting
	err = json.Unmarshal(responce.Body.Bytes(), &getResponse)
	if err != nil {
		t.Error("Failed to unmarshal")
	}

	assert.Equal(t, http.StatusOK, responce.Code)
	assert.Equal(t, expectedId, getResponse.Id)
	assert.Equal(t, expectedMessage, getResponse.Message)
}

func TestGetGreetingNotFound(t *testing.T) {
	// ARRANGE
	testGreetings := make(map[string]string)
	incorrectId := "87654321-8765-4321-8765-432187654321"
	testGreetings["12345678-9012-3456-7890-123456789012"] = "Hello World"

	// ACT
	fakeUUIDService := uuid.FakeUUIDService{}
	router := getRouter(testGreetings, &fakeUUIDService)
	responce := httptest.NewRecorder()
	request, err := http.NewRequest("GET", fmt.Sprintf("/greeting/%s", incorrectId), nil)
	if err != nil {
		t.Error("Failed to create request")
	}
	router.ServeHTTP(responce, request)

	// ASSERT
	assert.Equal(t, http.StatusNotFound, responce.Code)
	assert.Empty(t, responce.Body.String())
}

func TestGetGreetingNotFoundNotUUID(t *testing.T) {
	// ARRANGE
	testGreetings := make(map[string]string)
	incorrectId := "somethingThatDoesntExist"
	testGreetings["12345678-9012-3456-7890-123456789012"] = "Hello World"

	// ACT
	fakeUUIDService := uuid.FakeUUIDService{}
	router := getRouter(testGreetings, &fakeUUIDService)
	responce := httptest.NewRecorder()
	request, err := http.NewRequest("GET", fmt.Sprintf("/greeting/%s", incorrectId), nil)
	if err != nil {
		t.Error("Failed to create request")
	}
	router.ServeHTTP(responce, request)

	// ASSERT
	assert.Equal(t, http.StatusNotFound, responce.Code)
	assert.Empty(t, responce.Body.String())
}

func TestPutGreeting(t *testing.T) {
	// ARRANGE
	testGreetings := make(map[string]string)
	id := "12345678-9012-3456-7890-123456789012"
	originalMessage := "Original Message"
	testGreetings[id] = originalMessage
	var putBody PutRequest
	putBody.Message = "New Message Here"
	jsonBody, err := json.Marshal(putBody)
	if err != nil {
		t.Error("Failed to marshal")
	}

	// ACT
	// make an original get requst before the put request
	fakeUUIDService := uuid.FakeUUIDService{}
	router := getRouter(testGreetings, &fakeUUIDService)
	responce := httptest.NewRecorder()
	getRequestBefore, err := http.NewRequest("GET", fmt.Sprintf("/greeting/%s", id), nil)
	if err != nil {
		t.Error("Failed to create request")
	}
	router.ServeHTTP(responce, getRequestBefore)

	// ASSERT
	// ensure the original message was added during the ARRANGE steps
	var getResponseBefore Greeting
	err = json.Unmarshal(responce.Body.Bytes(), &getResponseBefore)
	if err != nil {
		t.Error("Failed to unmarshal")
	}

	assert.Equal(t, http.StatusOK, responce.Code)
	assert.Equal(t, id, getResponseBefore.Id)
	assert.Equal(t, originalMessage, getResponseBefore.Message)

	// ACT
	// make the PUT request to modify the message stored for the specific ID
	request, err := http.NewRequest("PUT", fmt.Sprintf("/greeting/%s", id), bytes.NewReader(jsonBody))
	if err != nil {
		t.Error("Failed to create request")
	}
	responce = httptest.NewRecorder()
	router.ServeHTTP(responce, request)

	// ASSERT
	// ensure the PUT request response code is 200 and the response body Id and Message are correct
	var putResponse Greeting
	err = json.Unmarshal(responce.Body.Bytes(), &putResponse)
	if err != nil {
		t.Error("Failed to unmarshal")
	}

	assert.Equal(t, http.StatusOK, responce.Code)
	assert.Equal(t, id, putResponse.Id)
	assert.Equal(t, putBody.Message, putResponse.Message)

	// ACT
	// make a new get requst after the put request
	getRequestAfter, err := http.NewRequest("GET", fmt.Sprintf("/greeting/%s", id), nil)
	if err != nil {
		t.Error("Failed to create request")
	}
	responce = httptest.NewRecorder()
	router.ServeHTTP(responce, getRequestAfter)

	// ASSERT
	// ensure the message has actually been modified in testGreetings not only the response body being correct
	var getResponseAfter Greeting
	err = json.Unmarshal(responce.Body.Bytes(), &getResponseAfter)
	if err != nil {
		t.Error("Failed to unmarshal")
	}

	assert.Equal(t, http.StatusOK, responce.Code)
	assert.Equal(t, id, getResponseAfter.Id)
	assert.Equal(t, putBody.Message, getResponseAfter.Message)
}

func TestPutGreetingNotFound(t *testing.T) {
	// ARRANGE
	testGreetings := make(map[string]string)
	id := "12345678-9012-3456-7890-123456789012"
	var putBody PutRequest
	putBody.Message = "New Message Here"
	jsonBody, err := json.Marshal(putBody)
	if err != nil {
		t.Error("Failed to marshal")
	}

	// ACT
	fakeUUIDService := uuid.FakeUUIDService{}
	router := getRouter(testGreetings, &fakeUUIDService)
	responce := httptest.NewRecorder()
	request, err := http.NewRequest("PUT", fmt.Sprintf("/greeting/%s", id), bytes.NewReader(jsonBody))
	if err != nil {
		t.Error("Failed to create request")
	}
	router.ServeHTTP(responce, request)

	// ASSERT
	assert.Equal(t, http.StatusNotFound, responce.Code)
	assert.Empty(t, responce.Body)
}

func TestPutGreetingNotFoundNotUUID(t *testing.T) {
	// ARRANGE
	testGreetings := make(map[string]string)
	id := "somethingThatDoesntExist"
	var putBody PutRequest
	putBody.Message = "New Message Here"
	jsonBody, err := json.Marshal(putBody)
	if err != nil {
		t.Error("Failed to marshal")
	}

	// ACT
	fakeUUIDService := uuid.FakeUUIDService{}
	router := getRouter(testGreetings, &fakeUUIDService)
	responce := httptest.NewRecorder()
	request, err := http.NewRequest("PUT", fmt.Sprintf("/greeting/%s", id), bytes.NewReader(jsonBody))
	if err != nil {
		t.Error("Failed to create request")
	}
	router.ServeHTTP(responce, request)

	// ASSERT
	assert.Equal(t, http.StatusNotFound, responce.Code)
	assert.Empty(t, responce.Body)
}

func TestPutGreetingMalformedRequest(t *testing.T) {
	// ARRANGE
	testGreetings := make(map[string]string)
	id := "12345678-9012-3456-7890-123456789012"
	testGreetings[id] = "Original Message"

	// ACT
	fakeUUIDService := uuid.FakeUUIDService{}
	router := getRouter(testGreetings, &fakeUUIDService)
	responce := httptest.NewRecorder()
	request, err := http.NewRequest("PUT", fmt.Sprintf("/greeting/%s", id), strings.NewReader(`{"someRubbish": "dataHere"}`))
	if err != nil {
		t.Error("Failed to create request")
	}
	router.ServeHTTP(responce, request)

	// ASSERT
	var errorResponse ErrorResponse
	err = json.Unmarshal(responce.Body.Bytes(), &errorResponse)
	if err != nil {
		t.Error("Failed to unmarshal")
	}

	assert.Equal(t, http.StatusBadRequest, responce.Code)
	assert.Equal(t, `Key: 'PutRequest.Message' Error:Field validation for 'Message' failed on the 'required' tag`, errorResponse.Error)
}

func TestPutGreetingMalformedRequestAndNotFound(t *testing.T) {
	// ARRANGE
	testGreetings := make(map[string]string)
	testGreetings["12345678-9012-3456-7890-123456789012"] = "Original Message"
	incorrectId := "87654321-8765-4321-8765-432187654321"

	// ACT
	fakeUUIDService := uuid.FakeUUIDService{}
	router := getRouter(testGreetings, &fakeUUIDService)
	responce := httptest.NewRecorder()
	request, err := http.NewRequest("PUT", fmt.Sprintf("/greeting/%s", incorrectId), strings.NewReader(`{"someRubbish": "dataHere"}`))
	if err != nil {
		t.Error("Failed to create request")
	}
	router.ServeHTTP(responce, request)

	// ASSERT
	var errorResponse ErrorResponse
	err = json.Unmarshal(responce.Body.Bytes(), &errorResponse)
	if err != nil {
		t.Error("Failed to unmarshal")
	}

	assert.Equal(t, http.StatusBadRequest, responce.Code)
	assert.Equal(t, `Key: 'PutRequest.Message' Error:Field validation for 'Message' failed on the 'required' tag`, errorResponse.Error)
}
