package main

import (
	"fmt"

	randomnumber "project/services/randomNumber"
	"project/services/uuid"
)

func main() {
	greetingsMap = make(map[string]string)
	greetingsMap["123"] = "Hello World"
	greetingsMap["456"] = "Yo!"
	greetingsMap["789"] = "Sup"
	greetingsMap["987"] = "Bonjour"
	greetingsMap["654"] = "Hi"
	greetingsMap["321"] = "Morning"

	port := "3000"

	fmt.Println("Starting API on port " + port)

	uuidService := uuid.RealUUIDService{}
	randomNumberService := randomnumber.RealRandomNumberService{}

	err := getRouter(greetingsMap, &uuidService, &randomNumberService).Run(":" + port)
	fmt.Println(err)
}

var greetingsMap map[string]string
