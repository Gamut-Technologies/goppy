package main

import (
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
