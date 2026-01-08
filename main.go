package main

import (
	"fmt"

	randomnumber "project/services/randomNumber"
	"project/services/uuid"

	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := dynamodb.NewFromConfig(cfg)

	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := client.ListTables(context.TODO(), &dynamodb.ListTablesInput{Limit: aws.Int32(5)})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("first page results")
	for _, object := range output.TableNames {
		fmt.Println("Table Name: " + object)
	}

	greetingsMap = make(map[string]string)

	port := "3000"

	fmt.Println("Starting API on port " + port)

	// this key would be stored extenally in real life projects
	key := "mySecretCodeABC123"

	uuidService := uuid.RealUUIDService{}
	randomNumberService := randomnumber.RealRandomNumberService{}

	err = getRouter(greetingsMap, &uuidService, &randomNumberService, key).Run(":" + port)
	fmt.Println(err)
}

var greetingsMap map[string]string
