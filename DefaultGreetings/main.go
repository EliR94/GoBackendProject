package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Greeting struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

func main() {
	var jsonStr = []byte(`{"message":"Good Morning"}`)
	req, err := http.NewRequest("POST", "http://greetings:3000/greeting", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Auth", "mySecretCodeABC123")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("response Status:", resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("response Body:", string(body))

	resp.Body.Close()
}
