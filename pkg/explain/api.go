package explain

import (
	"context"
	"fmt"
	"strings"

	ollama "github.com/jmorganca/ollama/api"
	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/pkg/shell"
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
		Model:   e.model,
		Prompt:  "Explain command: " + prompt,
		System:  e.system,
		Options: e.getParametersFor(e.mode),
	}, onResponseFunc)

	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("model '%s' not found", e.model)) {
			fmt.Println(fmt.Sprintf(shell.Err()+" %s - run `ollama pull %s` to download the model.", err.Error(), e.model))
			return nil
		}

		return cli.Exit(fmt.Sprintf(shell.Err()+"error getting explain: %s", err.Error()), -1)

	}
	fmt.Printf("\n")
	return nil
}
