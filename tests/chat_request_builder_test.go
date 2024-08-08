package endpoints

import (
	"encoding/json"
	"github.com/Gamut-Technologies/goppy.git/endpoints"
	"reflect"
	"testing"
)

// TestChatRequestBuilder tests the JSON serialization of the ChatRequest built by the builder.
func TestChatRequestBuilder(t *testing.T) {
	// Initialize test data
	messages := []endpoints.ChatMessage{
		{Role: "user", Content: "Hello!"},
	}
	model := "gpt-4"
	temperature := 0.7
	topP := 0.9

	expected := endpoints.ChatRequest{
		Messages: []endpoints.ChatMessage{
			{Role: "user", Content: "Hello!"},
		},
		Model:             "gpt-4",
		N:                 new(int),
		Temperature:       &temperature,
		TopP:              &topP,
		Stream:            new(bool),
		ParallelToolCalls: new(bool),
	}

	*expected.N = 0
	*expected.Stream = true
	*expected.ParallelToolCalls = false

	expectedJSON, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("failed to marshal expected JSON: %v", err)
	}

	// Create a ChatRequest using the builder
	builder := endpoints.NewChatRequestBuilder(messages, model)
	request := builder.SetTemperature(temperature).SetTopP(topP).SetN(0).SetStream(true).SetParallelToolCalls(false).Build()

	// Serialize to JSON
	data, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("failed to marshal request to JSON: %v", err)
	}

	// Compare the actual and expected JSON strings
	if !reflect.DeepEqual(data, expectedJSON) {
		t.Errorf("expected %s, got %s", string(expectedJSON), string(data))
	}
}
