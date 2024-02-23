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

func (e *Explain) streamExplanationFor(mode, prompt string) error {
	onResponseFunc := func(res ollama.GenerateResponse) error {
		fmt.Print(res.Response)
		return nil
	}

	err := e.api.Generate(context.Background(), &ollama.GenerateRequest{
		Model:   "codellama:7b",
		Prompt:  "Explain the command briefly: " + prompt,
		Options: e.getParametersFor(mode),
	}, onResponseFunc)

	if err != nil {
		fmt.Println("Error during generation:", err)
	}
	return nil
}
