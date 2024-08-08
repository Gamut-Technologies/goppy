package endpoints

// ChatRequest represents the data structure for a chat request.
// Documentation: https://platform.openai.com/docs/api-reference/chat/create
type ChatRequest struct {
	Messages          []ChatMessage    `json:"messages"`
	Model             string           `json:"model"`
	FrequencyPenalty  *float64         `json:"frequency_penalty,omitempty"`   // Optional, -2.0 to 2.0 or nil. Defaults to 0
	LogitBias         *map[int]float64 `json:"logit_bias,omitempty"`          // Optional, defaults to null
	Logprobs          *bool            `json:"logprobs,omitempty"`            // Optional, boolean or null. Defaults to false
	TopLogprobs       *int             `json:"top_logprobs,omitempty"`        // Optional, integer or null. 0 to 20. Requires logprobs to be true
	MaxTokens         *int             `json:"max_tokens,omitempty"`          // Optional, integer or null
	N                 *int             `json:"n,omitempty"`                   // Optional, integer or null. Defaults to 1
	PresencePenalty   *float64         `json:"presence_penalty,omitempty"`    // Optional, number or null. -2.0 to 2.0. Defaults to 0
	ResponseFormat    *ResponseFormat  `json:"response_format,omitempty"`     // Optional, object or null
	Seed              *int             `json:"seed,omitempty"`                // Optional, integer or null. Beta feature
	ServiceTier       *string          `json:"service_tier,omitempty"`        // Optional, string or null. Defaults to null
	Stop              *interface{}     `json:"stop,omitempty"`                // Optional, string, array or null. Defaults to null
	Stream            *bool            `json:"stream,omitempty"`              // Optional, boolean or null. Defaults to false
	StreamOptions     *StreamOptions   `json:"stream_options,omitempty"`      // Optional, object or null. Defaults to null
	Temperature       *float64         `json:"temperature,omitempty"`         // Optional, number or null. 0 to 2. Defaults to 1
	TopP              *float64         `json:"top_p,omitempty"`               // Optional, number or null. Defaults to 1
	Tools             *[]string        `json:"tools,omitempty"`               // Optional, array. Defaults to null
	ToolChoice        *interface{}     `json:"tool_choice,omitempty"`         // Optional, string or object. Defaults to null
	ParallelToolCalls *bool            `json:"parallel_tool_calls,omitempty"` // Optional, boolean. Defaults to true
	User              *string          `json:"user,omitempty"`                // Optional, string. Unique identifier for the end-user
}

// ChatMessage represents a message in the chat
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// StreamOptions represents options for streaming responses
type StreamOptions struct {
	IncludeUsage *bool `json:"include_usage"` // Optional, boolean. Defaults to null
}

// ResponseFormat represents the format that the model must output
type ResponseFormat struct {
	Type       string      `json:"type"`
	JSONSchema *JSONSchema `json:"json_schema,omitempty"` // Included only if type is "json_schema"
}

// JSONSchema represents the schema for structured outputs
type JSONSchema struct {
	Description string                 `json:"description,omitempty"` // Optional
	Name        string                 `json:"name"`                  // Required
	Schema      map[string]interface{} `json:"schema,omitempty"`      // Optional
	Strict      *bool                  `json:"strict,omitempty"`      // Optional, defaults to false
}

// ChatResponse represents the response from the chat API
type ChatResponse struct {
	ID                string       `json:"id"`
	Object            string       `json:"object"`
	Created           int          `json:"created"`
	Model             string       `json:"model"`
	SystemFingerPrint string       `json:"system_finger_print"`
	Choices           []Choice     `json:"choices"`
	Usage             UsageDetails `json:"usage,omitempty"`
}

// Choice represents a single response choice from the chat API
type Choice struct {
	Message      ChatMessage `json:"message"`
	Index        int         `json:"index"`
	Logprobs     *Logprobs   `json:"logprobs,omitempty"`
	FinishReason string      `json:"finish_reason"`
}

// Logprobs represents the log probabilities of tokens
type Logprobs struct {
	Tokens        []string             `json:"tokens"`
	TokenLogprobs []float64            `json:"token_logprobs"`
	TopLogprobs   []map[string]float64 `json:"top_logprobs"`
	TextOffset    []int                `json:"text_offset"`
}

// UsageDetails represents the usage details of the chat request
type UsageDetails struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ChatCompletionChunk represents a chunk of the chat completion response
type ChatCompletionChunk struct {
	ID                string        `json:"id"`
	Object            string        `json:"object"`
	Created           int           `json:"created"`
	Model             string        `json:"model"`
	SystemFingerPrint string        `json:"system_fingerprint"`
	Choices           []ChunkChoice `json:"choices"`
}

// ChunkChoice represents a single chunk choice in the chat completion chunk
type ChunkChoice struct {
	Index        int         `json:"index"`
	Delta        ChatMessage `json:"delta"`
	Logprobs     *Logprobs   `json:"logprobs,omitempty"`
	FinishReason string      `json:"finish_reason"`
}

// ChatRequestBuilder helps to build a ChatRequest struct
type ChatRequestBuilder struct {
	request *ChatRequest
}

// NewChatRequestBuilder initializes and returns a new ChatRequestBuilder
func NewChatRequestBuilder(messages []ChatMessage, model string) *ChatRequestBuilder {
	return &ChatRequestBuilder{
		request: &ChatRequest{
			Messages: messages,
			Model:    model,
		},
	}
}

func (b *ChatRequestBuilder) SetFrequencyPenalty(value float64) *ChatRequestBuilder {
	b.request.FrequencyPenalty = &value
	return b
}

func (b *ChatRequestBuilder) SetLogitBias(value map[int]float64) *ChatRequestBuilder {
	b.request.LogitBias = &value
	return b
}

func (b *ChatRequestBuilder) SetLogprobs(value bool) *ChatRequestBuilder {
	b.request.Logprobs = &value
	return b
}

func (b *ChatRequestBuilder) SetTopLogprobs(value int) *ChatRequestBuilder {
	b.request.TopLogprobs = &value
	return b
}

func (b *ChatRequestBuilder) SetMaxTokens(value int) *ChatRequestBuilder {
	b.request.MaxTokens = &value
	return b
}

func (b *ChatRequestBuilder) SetN(value int) *ChatRequestBuilder {
	b.request.N = &value
	return b
}

func (b *ChatRequestBuilder) SetPresencePenalty(value float64) *ChatRequestBuilder {
	b.request.PresencePenalty = &value
	return b
}

func (b *ChatRequestBuilder) SetResponseFormat(value ResponseFormat) *ChatRequestBuilder {
	b.request.ResponseFormat = &value
	return b
}

func (b *ChatRequestBuilder) SetSeed(value int) *ChatRequestBuilder {
	b.request.Seed = &value
	return b
}

func (b *ChatRequestBuilder) SetServiceTier(value string) *ChatRequestBuilder {
	b.request.ServiceTier = &value
	return b
}

func (b *ChatRequestBuilder) SetStop(value interface{}) *ChatRequestBuilder {
	b.request.Stop = &value
	return b
}

func (b *ChatRequestBuilder) SetStream(value bool) *ChatRequestBuilder {
	b.request.Stream = &value
	return b
}

func (b *ChatRequestBuilder) SetStreamOptions(value StreamOptions) *ChatRequestBuilder {
	b.request.StreamOptions = &value
	return b
}

func (b *ChatRequestBuilder) SetTemperature(value float64) *ChatRequestBuilder {
	b.request.Temperature = &value
	return b
}

func (b *ChatRequestBuilder) SetTopP(value float64) *ChatRequestBuilder {
	b.request.TopP = &value
	return b
}

func (b *ChatRequestBuilder) SetTools(value []string) *ChatRequestBuilder {
	b.request.Tools = &value
	return b
}

func (b *ChatRequestBuilder) SetToolChoice(value interface{}) *ChatRequestBuilder {
	b.request.ToolChoice = &value
	return b
}

func (b *ChatRequestBuilder) SetParallelToolCalls(value bool) *ChatRequestBuilder {
	b.request.ParallelToolCalls = &value
	return b
}

func (b *ChatRequestBuilder) SetUser(value string) *ChatRequestBuilder {
	b.request.User = &value
	return b
}

// Build finalizes the ChatRequest and returns it
func (b *ChatRequestBuilder) Build() *ChatRequest {
	return b.request
}
