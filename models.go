package main

// This struct represents the data structure of a Greeting response body
type Greeting struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

// This struct represents the body for an incoming POST request
type PostRequest struct {
	Message string `json:"message" binding:"required"`
}

// This struct represents the body for an incoming PUT request
type PutRequest struct {
	Message string `json:"message" binding:"required"`
}

// This struct represents the GET request response on endpoint /greetings
type ResponceMap struct {
	Items []Greeting `json:"items"`
}

// This struct represents the error response for a bad POST request on endpoint /greeting
type ErrorResponse struct {
	Error string `json:"error"`
}
