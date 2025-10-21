package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "3000"

	fmt.Println("Starting API on port " + port)
	err := getRouter().Run(":" + port)
	fmt.Println(err)
}

func getHealthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, "All good!")
}

func getRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/healthcheck", getHealthcheck)

	return r
}
