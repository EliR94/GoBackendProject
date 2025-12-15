package main

import (
	randomnumber "project/services/randomNumber"
	"project/services/uuid"

	"github.com/gin-gonic/gin"
)

func getRouter(initialGreetings map[string]string, uuidService uuid.UUIDService, randomNumberService randomnumber.RandomNumberService, secretKey string) *gin.Engine {
	greetingsMap = initialGreetings

	r := gin.New()
	r.Use(loggerMiddleware(uuidService))
	r.Use(gin.Recovery())

	r.GET("/healthcheck", getHealthcheck)
	r.GET("/greetings", getGreetings)
	r.GET("/greeting/random", func(c *gin.Context) {
		getRandomGreeting(c, randomNumberService)
	})
	r.GET("/greeting/:uuid", getGreeting)
	restrictedEndpoints := r.Group("/greeting")
	restrictedEndpoints.Use(authenticationMiddleware(secretKey))
	{
		restrictedEndpoints.POST("", func(c *gin.Context) {
			postGreeting(c, uuidService)
		})
		restrictedEndpoints.PUT("/:uuid", putGreeting)
		restrictedEndpoints.DELETE("/:uuid", deleteGreeting)
	}

	return r
}
