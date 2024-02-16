package api

import (
	"time"
)

type options struct { // This is a struct that contains the options for the request
	Temperature float64 `json:"temperature"`
	TopP        float64 `json:"top_p"`
	TopK        int     `json:"top_k"`
}

type generateRequestPayload struct {
	Model   string  `json:"model"`
	System  string  `json:"system"`
	Prompt  string  `json:"prompt"`
	Stream  bool    `json:"stream"`
	Options options `json:"options"`
}

type generateResponsePayload struct {
	Model              string    `json:"model"`
	CreatedAt          time.Time `json:"created_at"`
	Response           string    `json:"response"`
	Done               bool      `json:"done"`
	Context            []int     `json:"context"`
	TotalDuration      int       `json:"total_duration"`
	LoadDuration       int       `json:"load_duration"`
	PromptEvalCount    int       `json:"prompt_eval_count"`
	PromptEvalDuration int       `json:"prompt_eval_duration"`
	EvalCount          int       `json:"eval_count"`
	EvalDuration       int       `json:"eval_duration"`
}
