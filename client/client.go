package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Gamut-Technologies/goppy/endpoints"
	"net/http"
	"os"
)

const (
	baseUrl = "https://api.openai.com/"
)

// GoppyClient represents an OpenAI API Client
type GoppyClient struct {
	apiVersion string   // Specific OpenAI api version
	apiUrl     string   // baseUrl + version + uri?
	endpoints  []string // endpoints = api reference endpoints
	apiKey     string   // OpenAI api key
	projectId  string   // OpenAI project
	orgId      string   // OpenAI organization
	userAgent  string   // User agent header
}

// Create local types of all supported endpoints

type Chat = *endpoints.ChatRequest

// MissingApiKeyError is a custom error for missing API key
type MissingApiKeyError struct{}

func (e *MissingApiKeyError) Error() string {
	return "API not available in arguments or environment variables. Run `export OPEN_AI_API_KEY={your_api_key}`"
}

// NewClient initializes and returns a new GoppyClient
func NewClient(apiVersion string, apiKey string, projectId string, orgId string) (*GoppyClient, error) {
	if apiVersion == "" {
		apiVersion = "v1"
	}

	if apiKey == "" {
		apiKey = os.Getenv("OPEN_AI_API_KEY")
		if apiKey == "" {
			return nil, &MissingApiKeyError{}
		}
	}

	if projectId == "" {
		projectId = os.Getenv("OPEN_AI_PROJECT_ID")
	}

	if orgId == "" {
		orgId = os.Getenv("OPEN_AI_ORGANIZATION_ID")
	}

	return &GoppyClient{
		apiVersion: apiVersion,
		apiUrl:     baseUrl + apiVersion,
		userAgent:  "goppy/1.0",
		projectId:  projectId,
		orgId:      orgId,
		apiKey:     apiKey,
	}, nil
}

// MarshalRequest attempts to marshal request data as JSON for the request body
func (c *GoppyClient) MarshalRequest(reqData any) ([]byte, error) {
	data, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UnableToMarshalRequestData is thrown when the provided request data cannot be marshalled to JSON
type UnableToMarshalRequestData struct{}

func (e *UnableToMarshalRequestData) Error() string {
	return "Unable to marshal request data as JSON."
}

// Request executes the request to the OpenAI API with the provided model configuration and data
func (c *GoppyClient) Request(requestData interface{}) (*http.Response, error) {
	data, err := c.MarshalRequest(requestData)
	if err != nil {
		return nil, &UnableToMarshalRequestData{}
	}

	endpoint, err := getRequestEndpoint(requestData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.apiUrl+"/"+endpoint, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	if c.orgId != "" {
		req.Header.Set("OpenAI-Organization", c.orgId)
	}

	if c.projectId != "" {
		req.Header.Set("OpenAI-Project", c.projectId)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Response decodes the HTTP response from OpenAI into a given struct
func (c *GoppyClient) Response(resp *http.Response, result interface{}, stream bool) error {
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return c.handleAPIError(resp)
	}

	decoder := json.NewDecoder(resp.Body)
	if !stream {
		if err := decoder.Decode(result); err != nil {
			return fmt.Errorf("failed to decode JSON response: %w", err)
		}
	} else {
		// Handle stream response
	}

	return nil
}

// EndpointNotAvailableError should only be thrown if the provided request data does not match a endpoint interface
type EndpointNotAvailableError struct{}

func (e *EndpointNotAvailableError) Error() string {
	return "Request data does not match an available endpoint. Validate your request data matches an endpoint interface."
}

// getRequestEndpoint determines which endpoint interface the data matches and returns the OpenAI endpoint or returns an EndpointNotAvailableError
func getRequestEndpoint(requestData interface{}) (string, error) {
	switch requestData.(type) {
	case Chat:
		return "chat/completions", nil

	default:
		return "", &EndpointNotAvailableError{}
	}
}

// APIError represents an error from OpenAi
type APIError struct {
	StatusCode int
	Message    string
	Details    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error: %d - %s: %s", e.StatusCode, e.Message, e.Details)
}

type APIErrorResponse struct {
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Param   string `json:"param,omitempty"`
		Code    string `json:"code,omitempty"`
	} `json:"error"`
}

// handleAPIError decodes and formats and OpenAI Error
func (c *GoppyClient) handleAPIError(resp *http.Response) error {
	var apiErr APIErrorResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to parse API error response",
			Details:    err.Error(),
		}
	}

	return &APIError{
		StatusCode: resp.StatusCode,
		Message:    apiErr.Error.Message,
		Details:    fmt.Sprintf("Type: %s, Code: %s, Param: %s", apiErr.Error.Type, apiErr.Error.Code, apiErr.Error.Param),
	}
}
