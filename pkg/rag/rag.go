package rag

import (
	"context"
	_ "embed"
	"fmt"
	"strings"

	ollama "github.com/jmorganca/ollama/api"
)

type RAGChat struct {
	api *ollama.Client

	model   string
	context string
	history []ollama.Message
}

func (r *RAGChat) Send(message string, numCtx int) (string, error) {
	// string builder to collect streaming response
	sb := strings.Builder{}

	// if no history, add context as the first message
	if len(r.history) == 0 {
		// Add context as the first message
		r.history = append(r.history, ollama.Message{Role: "system",
			Content: "You are a software engineer and a helpful assistant.",
		})

		// Add context as the first message
		r.history = append(r.history, ollama.Message{
			Role:    "user",
			Content: r.context + "\n" + message,
		})
	} else { // if history exists, add the new message to the history
		r.history = append(r.history, ollama.Message{
			Role:    "user",
			Content: message,
		})
	}

	err := r.api.Chat(context.Background(), &ollama.ChatRequest{
		Model:    r.model,
		Messages: r.history,
		Options: map[string]interface{}{
			"temperature": 0.5,
			// "num_ctx":     numCtx,
		},
	}, func(res ollama.ChatResponse) error {
		fmt.Print(res.Message.Content)
		sb.WriteString(res.Message.Content)
		return nil
	})

	if err != nil {
		return "", fmt.Errorf("error sending message: %s", err.Error())
	}

	// append assistant response to history
	r.history = append(r.history, ollama.Message{
		Role:    "assistant",
		Content: sb.String(),
	})

	return "", nil
}

func NewRAGChat(api *ollama.Client, context string, model string) *RAGChat {
	return &RAGChat{
		api:     api,
		model:   model,
		context: context,
		history: make([]ollama.Message, 0),
	}
}
