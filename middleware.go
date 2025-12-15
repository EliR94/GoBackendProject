package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"project/services/uuid"
)

func loggerMiddleware(uuidService uuid.UUIDService) gin.HandlerFunc {
	return func(c *gin.Context) {
		corellationId := uuidService.NewUUID()
		c.Set("corellationId", corellationId)
		requestMethod := c.Request.Method
		requestEndpoint := c.Request.URL.Path
		requestTime := time.Now()
		beforeLog := fmt.Sprintf("Corellation ID: %s / Request Method: %s / Request Endpoint: %s / Request Time: %s", corellationId, requestMethod, requestEndpoint, requestTime.Format(time.RFC3339))
		fmt.Println(beforeLog)

		c.Next()

		responseCode := c.Writer.Status()
		responseTime := time.Now()
		duration := responseTime.Sub(requestTime)
		afterLog := fmt.Sprintf("Corellation ID: %s / Response Code: %d / Response Time: %s / Duration of Request: %v\n", corellationId, responseCode, responseTime.Format(time.RFC3339), duration)
		fmt.Println(afterLog)
	}
}

func authenticationMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		xAuth := c.Request.Header.Get("X-Auth")
		if xAuth == secretKey && xAuth != "" {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
