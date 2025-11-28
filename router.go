package main

import (
	"project/services/uuid"

	"github.com/gin-gonic/gin"
)

func getRouter(initialGreetings map[string]string, uuidService uuid.UUIDService) *gin.Engine {
	greetingsMap = initialGreetings

	r := gin.Default()

	r.GET("/healthcheck", getHealthcheck)
	r.GET("/greetings", getGreetings)
	r.GET("/greeting/:uuid", getGreeting)
	r.POST("/greeting", func(c *gin.Context) {
		postGreeting(c, uuidService)
	})
	r.PUT("/greeting/:uuid", putGreeting)
	r.DELETE("/greeting/:uuid", deleteGreeting)

	return r
}
