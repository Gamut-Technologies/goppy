package main

import (
	"encoding/json"
	"fmt"
	"github.com/Gamut-Technologies/goppy.git/endpoints"
	"log"
)

func main() {
	client, err := NewClient("", "", "", "")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(client)

	messages := []endpoints.ChatMessage{
		{Role: "user", Content: "Hello!"},
	}
	model := "gpt-4"

	builder := endpoints.NewChatRequestBuilder(messages, model)
	request := builder.
		SetTemperature(0.7).
		SetTopP(0.9).
		SetStream(true).
		Build()

	data, _ := json.Marshal(request)
	fmt.Println(string(data))
}
