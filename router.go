package main

import (
	randomnumber "project/services/randomNumber"
	"project/services/uuid"

	"github.com/gin-gonic/gin"
)

func getRouter(initialGreetings map[string]string, uuidService uuid.UUIDService, randomNumberService randomnumber.RandomNumberService) *gin.Engine {
	greetingsMap = initialGreetings

	r := gin.Default()

	r.GET("/healthcheck", getHealthcheck)
	r.GET("/greetings", getGreetings)
	r.GET("/greeting/random", func(c *gin.Context) {
		getRandomGreeting(c, randomNumberService)
	})
	r.GET("/greeting/:uuid", getGreeting)
	r.POST("/greeting", func(c *gin.Context) {
		postGreeting(c, uuidService)
	})
	r.PUT("/greeting/:uuid", putGreeting)
	r.DELETE("/greeting/:uuid", deleteGreeting)

	return r
}
