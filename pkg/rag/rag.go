package rag

import (
	"context"
	_ "embed"
	"fmt"
	"strings"

	ollama "github.com/jmorganca/ollama/api"
	"github.com/spf13/viper"
)

type RAGChat struct {
	api *ollama.Client

	model   string
	context string
	history []ollama.Message
}

func (r *RAGChat) Send(message string) (string, error) {
	// string builder to collect streaming response
	sb := strings.Builder{}

	// Add context as the first message
	messages := []ollama.Message{
		{Role: "system",
			Content: "You are a software engineer and a helpful assistant.",
		},
	}

	// Add context as the first message
	r.history = append(r.history, ollama.Message{
		Role:    "user",
		Content: r.context + "\n" + message,
	})
	messages = append(messages, r.history...)

	err := r.api.Chat(context.Background(), &ollama.ChatRequest{
		Model:    r.model,
		Messages: messages,
		Options: map[string]interface{}{
			"temperature": 0.1,
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

func NewRAGChat(api *ollama.Client, context string) *RAGChat {
	model := viper.GetString("llm.model")

	return &RAGChat{
		api:     api,
		model:   model,
		context: context,
		history: make([]ollama.Message, 0),
	}
}
