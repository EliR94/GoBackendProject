package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Greeting struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

func main() {
	for {
		resp, err := http.Get("http://greetings:3000/greeting/random")
		if err != nil {
			panic(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var greeting Greeting
		json.Unmarshal(body, &greeting)

		fmt.Println(greeting.Message)
		resp.Body.Close()

		time.Sleep(time.Second * 5)
	}
}
