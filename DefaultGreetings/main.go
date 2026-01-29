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
	greetingsMap := make(map[string][]byte)
	greetingsMap["12345678-9012-3456-7890-123456789012"] = []byte(`{"message":"Hello World"}`)
	greetingsMap["12345678-9012-3456-7890-123456789013"] = []byte(`{"message":"Hey"}`)
	greetingsMap["12345678-9012-3456-7890-123456789014"] = []byte(`{"message":"Howdy"}`)
	greetingsMap["12345678-9012-3456-7890-123456789015"] = []byte(`{"message":"Sup!"}`)
	greetingsMap["12345678-9012-3456-7890-123456789016"] = []byte(`{"message":"Yo Yo Yo"}`)
	greetingsMap["12345678-9012-3456-7890-123456789017"] = []byte(`{"message":"Wassup"}`)
	greetingsMap["12345678-9012-3456-7890-123456789018"] = []byte(`{"message":"Bonjour"}`)
	greetingsMap["12345678-9012-3456-7890-123456789019"] = []byte(`{"message":"Γειά σου"}`)
	for _, message := range greetingsMap {
		req, err := http.NewRequest("POST", "http://greetings:3000/greeting", bytes.NewBuffer(message))
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
}
