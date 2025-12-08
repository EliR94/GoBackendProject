package main

import (
	"fmt"

	randomnumber "project/services/randomNumber"
	"project/services/uuid"
)

func main() {
	greetingsMap = make(map[string]string)

	port := "3000"

	fmt.Println("Starting API on port " + port)

	uuidService := uuid.RealUUIDService{}
	randomNumberService := randomnumber.RealRandomNumberService{}

	err := getRouter(greetingsMap, &uuidService, &randomNumberService).Run(":" + port)
	fmt.Println(err)
}

var greetingsMap map[string]string
