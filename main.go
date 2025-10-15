package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "3000"
	r := gin.Default()
	fmt.Println("Starting API on port " + port)
	err := r.Run(":" + port)
	fmt.Println(err)
}
