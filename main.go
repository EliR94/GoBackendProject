package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	greetingsMap = make(map[string]string)

	greetingsMap["12345678-9012-3456-7890-123456789012"] = "Hello World"
	greetingsMap["12345678-9012-3456-7890-123456789013"] = "Hey"
	greetingsMap["12345678-9012-3456-7890-123456789014"] = "Howdy"
	greetingsMap["12345678-9012-3456-7890-123456789015"] = "Sup!"
	greetingsMap["12345678-9012-3456-7890-123456789016"] = "Yo Yo Yo"
	greetingsMap["12345678-9012-3456-7890-123456789017"] = "Wassup"
	greetingsMap["12345678-9012-3456-7890-123456789018"] = "Bonjour"
	greetingsMap["12345678-9012-3456-7890-123456789019"] = "Γειά σου"

	port := "3000"

	fmt.Println("Starting API on port " + port)
	err := getRouter().Run(":" + port)
	fmt.Println(err)
}

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

type Greeting struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

func getRouter() *gin.Engine {
	// i have a list of greetings, what do?

	r := gin.Default()

	r.GET("/healthcheck", getHealthcheck)
	r.GET("/greetings", getGreetings)

	return r
}

var greetingsMap map[string]string
