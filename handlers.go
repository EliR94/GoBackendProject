package main

import (
	"maps"
	"net/http"
	"slices"

	randomnumber "project/services/randomNumber"
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

func getGreeting(c *gin.Context) {
	uuid := c.Param("uuid")
	message, exists := greetingsMap[uuid]
	if exists {
		c.JSON(http.StatusOK, gin.H{
			"id":      uuid,
			"message": message,
		})
	} else {
		c.Status(http.StatusNotFound)
	}
}

func getRandomGreeting(c *gin.Context, randomNumber randomnumber.RandomNumberService) {
	greetingsMapLength := len(greetingsMap)
	if greetingsMapLength > 0 {
		randomNumber := randomNumber.NewRandomNumber(greetingsMapLength)

		var greetingIdsSlice []string
		for greetingId := range maps.Keys(greetingsMap) {
			greetingIdsSlice = append(greetingIdsSlice, greetingId)
		}
		slices.Sort(greetingIdsSlice)
		randomGreetingId := greetingIdsSlice[randomNumber]
		randomGreetingMessage := greetingsMap[randomGreetingId]
		c.JSON(http.StatusOK, gin.H{
			"id":      randomGreetingId,
			"message": randomGreetingMessage,
		})
	} else {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "no greetings in system"})
	}
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

func putGreeting(c *gin.Context) {
	var requestBody PutRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uuid := c.Param("uuid")
	_, exists := greetingsMap[uuid]
	if exists {
		greetingsMap[uuid] = requestBody.Message
		c.JSON(http.StatusOK, gin.H{
			"id":      uuid,
			"message": requestBody.Message,
		})
	} else {
		c.Status(http.StatusNotFound)
	}
}

func deleteGreeting(c *gin.Context) {
	uuid := c.Param("uuid")
	_, exists := greetingsMap[uuid]
	if exists {
		delete(greetingsMap, uuid)
		c.Status(http.StatusNoContent)
	} else {
		c.Status(http.StatusNotFound)
	}
}
