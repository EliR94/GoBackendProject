package main

import (
	"net/http"

	"project/services/uuid"

	"github.com/gin-gonic/gin"
)

func getHealthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, "All good!")
}

func getGreetings(c *gin.Context) {
	mapOfIdtoMessage := make(map[string][]Greeting)

	var itemsSlice []Greeting
	for id, message := range greetingsMap {
		formattedGreeting := Greeting{
			Id:      id,
			Message: message,
		}
		itemsSlice = append(itemsSlice, formattedGreeting)
	}

	mapOfIdtoMessage["items"] = itemsSlice
	c.JSON(http.StatusOK, mapOfIdtoMessage)
}

func postGreeting(c *gin.Context, uuidService uuid.UUIDService) {

	var requestBody PostRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuidService.NewUUID()
	greetingsMap[id] = requestBody.Message
	c.JSON(http.StatusCreated, gin.H{
		"id":      id,
		"message": requestBody.Message,
	})
}
