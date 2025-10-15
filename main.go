package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "3000"
	r := gin.Default()

	r.GET("/healthcheck", getHealthcheck)

	fmt.Println("Starting API on port " + port)
	err := r.Run(":" + port)
	fmt.Println(err)
}

func getHealthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, "All good!")
}
