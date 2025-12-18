package main

import (
	"fmt"

	randomnumber "project/services/randomNumber"
	"project/services/uuid"
)

func main() {
	greetingsMap = make(map[string]string)

	// default greetings for demo
	greetingsMap["12345678-9012-3456-7890-123456789012"] = "Hello World"
	greetingsMap["12345678-9012-3456-7890-123456789013"] = "Hey"
	greetingsMap["12345678-9012-3456-7890-123456789014"] = "Howdy"
	greetingsMap["12345678-9012-3456-7890-123456789015"] = "Sup!"
	greetingsMap["12345678-9012-3456-7890-123456789016"] = "Yo Yo Yo"
	greetingsMap["12345678-9012-3456-7890-123456789017"] = "Wassup"
	greetingsMap["12345678-9012-3456-7890-123456789018"] = "Bonjour"
	greetingsMap["12345678-9012-3456-7890-123456789019"] = "Γειά σου"

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
