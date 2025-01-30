package suggest

import (
	"context"
	"fmt"
	"regexp"
	"runtime"
	"strings"

	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlm/pkg/config"
	"github.com/yusufcanb/tlm/pkg/shell"
)

var ShellPrefix = []string{"$", "❯", "#"}

func (s *Suggest) getParametersFor(preference string) map[string]interface{} {
	switch preference {
	case config.Stable:
		return map[string]interface{}{
			"seed":        42,
			"temperature": 0.1,
			"top_p":       0.25,
		}

	case config.Balanced:
		return map[string]interface{}{
			"seed":        21,
			"temperature": 0.5,
			"top_p":       0.4,
		}

	case config.Creative:
		return map[string]interface{}{
			"seed":        0,
			"temperature": 0.9,
			"top_p":       0.7,
		}

	default:
		return map[string]interface{}{
			"seed":        21,
			"temperature": 0.5,
			"top_p":       0.4,
		}
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
			codeSnippets = append(codeSnippets, s.refineCommand(match[2]))
		}
	}

	return codeSnippets
}

func (s *Suggest) refineCommand(command string) string {
	result := strings.Clone(command)

	// Trim shell prefixes
	for _, prefix := range ShellPrefix {
		if strings.HasPrefix(result, prefix) {
			result = strings.TrimPrefix(result, prefix)
			break
		}
	}

	// Trim leading and trailing whitespaces
	result = strings.TrimSpace(result)

	return result
}

func (s *Suggest) getCommandSuggestionFor(term string, prompt string) (string, error) {
	var responseText string

	builder := strings.Builder{}
	builder.WriteString(prompt)

	usingTerminalStr := ". I'm using %s terminal"
	onOperatingSystemStr := "on operating system: %s"

	switch term {
	case "zsh":
		builder.WriteString(fmt.Sprintf(usingTerminalStr, term))
		builder.WriteString(fmt.Sprintf(onOperatingSystemStr, "macOS"))
	case "bash":
		builder.WriteString(fmt.Sprintf(usingTerminalStr, term))
		builder.WriteString(fmt.Sprintf(onOperatingSystemStr, "Linux"))
	case "powershell":
		builder.WriteString(fmt.Sprintf(usingTerminalStr, term))
		builder.WriteString(fmt.Sprintf(onOperatingSystemStr, "Windows"))

	default:
		builder.WriteString(fmt.Sprintf(usingTerminalStr, shell.GetShell()))
		builder.WriteString(fmt.Sprintf(onOperatingSystemStr, runtime.GOOS))
	}

	stream := false
	req := &ollama.GenerateRequest{
		Model:   s.model,
		System:  s.system,
		Prompt:  builder.String(),
		Stream:  &stream,
		Options: s.getParametersFor(s.style),
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
