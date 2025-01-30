package explain

import (
	"context"
	"fmt"

	ollama "github.com/jmorganca/ollama/api"
)

const (
	Stable   string = "stable"
	Balanced        = "balanced"
	Creative        = "creative"
)

func (e *Explain) getParametersFor(preference string) map[string]interface{} {
	switch preference {
	case Stable:
		return map[string]interface{}{
			"temperature": 0.1,
			"top_p":       0.25,
		}

	case Balanced:
		return map[string]interface{}{
			"temperature": 0.5,
			"top_p":       0.4,
		}

	case Creative:
		return map[string]interface{}{
			"temperature": 0.9,
			"top_p":       0.7,
		}

	default:
		return map[string]interface{}{}
	}
}

func (e *Explain) StreamExplanationFor(mode, prompt string) error {
	onResponseFunc := func(res ollama.GenerateResponse) error {
		fmt.Print(res.Response)
		return nil
	}

	err := e.api.Generate(context.Background(), &ollama.GenerateRequest{
		Model:   "llama3.2:1b",
		Prompt:  "Explain command: " + prompt,
		System:  `You are a command line application which helps user to get brief explanations for shell commands. You will be explaining given executable shell command to user with shortest possible explanation. If given input is not a shell command, you will respond with "I can only explain shell commands. Please provide a shell command to explain". You will never respond any question out of shell command explanation context.`,
		Options: e.getParametersFor(mode),
	}, onResponseFunc)

	if err != nil {
		fmt.Println("Error during generation:", err)
	}
	fmt.Printf("\n")
	return nil
}
