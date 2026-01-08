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

	// this key would be stored extenally in real life projects
	key := "mySecretCodeABC123"

	uuidService := uuid.RealUUIDService{}
	randomNumberService := randomnumber.RealRandomNumberService{}

	err := getRouter(greetingsMap, &uuidService, &randomNumberService, key).Run(":" + port)
	fmt.Println(err)
}

var greetingsMap map[string]string
