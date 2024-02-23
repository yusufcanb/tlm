package suggest

import (
	"context"
	"fmt"
	ollama "github.com/jmorganca/ollama/api"
	"regexp"
	"runtime"
	"strings"
)

const (
	Stable   string = "stable"
	Balanced        = "balanced"
	Creative        = "creative"
)

func (s *Suggest) getParametersFor(preference string) map[string]interface{} {
	switch preference {
	case Stable:
		return map[string]interface{}{
			"seed":        42,
			"temperature": 0.1,
			"top_p":       0.25,
		}

	case Balanced:
		return map[string]interface{}{
			"seed":        42,
			"temperature": 0.5,
			"top_p":       0.4,
		}

	case Creative:
		return map[string]interface{}{
			"seed":        0,
			"temperature": 0.9,
			"top_p":       0.7,
		}

	default:
		return map[string]interface{}{}
	}
}

func (s *Suggest) extractCommandsFromResponse(response string) []string {
	re := regexp.MustCompile("```([^\n]*)\n([^\n]*)\n```")

	matches := re.FindAllStringSubmatch(response, -1)

	if len(matches) == 0 {
		return nil
	}

	var codeSnippets []string
	for _, match := range matches {
		if len(match) == 3 {
			codeSnippets = append(codeSnippets, match[2])
		}
	}

	return codeSnippets
}

func (s *Suggest) getCommandSuggestionFor(mode, shell string, prompt string) (string, error) {
	var responseText string

	builder := strings.Builder{}
	builder.WriteString(prompt)
	builder.WriteString(fmt.Sprintf(". I'm using %s terminal", shell))
	builder.WriteString(fmt.Sprintf("on operating system: %s", runtime.GOOS))

	stream := false
	req := &ollama.GenerateRequest{
		Model:   "suggest:7b",
		Prompt:  builder.String(),
		Stream:  &stream,
		Options: s.getParametersFor(mode),
	}

	onResponse := func(res ollama.GenerateResponse) error {
		responseText = res.Response
		return nil
	}

	err := s.api.Generate(context.Background(), req, onResponse)
	if err != nil {
		return "", err
	}

	return responseText, nil
}
