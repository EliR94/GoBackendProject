package main

import (
	"fmt"

	"project/services/uuid"
)

func main() {
	greetingsMap = make(map[string]string)

	port := "3000"

	fmt.Println("Starting API on port " + port)

	uuidService := uuid.RealUUIDService{}

	err := getRouter(greetingsMap, &uuidService).Run(":" + port)
	fmt.Println(err)
}

var greetingsMap map[string]string
