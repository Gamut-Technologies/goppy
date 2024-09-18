package main

import (
	"fmt"
	client2 "github.com/Gamut-Technologies/goppy.git/client"
	"github.com/Gamut-Technologies/goppy.git/endpoints"
	"log"
)

func main() {
	client, err := client2.NewClient("", "", "", "")
	if err != nil {
		log.Fatalf(err.Error())
	}

	// TODO: Make request with client using build ChatRequestBuilder result
	messages := []endpoints.Message{
		{Role: "user", Content: "Hello"},
	}

	builder := endpoints.Chat(messages, "")
	requestData := builder.
		SetModel("gpt-4").
		SetTemperature(0.7).
		SetTopP(0.9).
		//SetStream(true).
		Build()

	req, err := client.Request(requestData)
	if err != nil {
		fmt.Println(err)
	}

	var chatRes endpoints.ChatResponse
	err = client.Response(req, &chatRes, false)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(chatRes)
}
