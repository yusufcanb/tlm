package ask

import (
	"bufio"
	"context"
	_ "embed"
	"fmt"
	"os"
	"strings"

	ollama "github.com/jmorganca/ollama/api"
	"github.com/spf13/viper"
	"github.com/yusufcanb/tlm/pkg/config"
)

//go:embed SYSTEM
var system string

type Ask struct {
	api     *ollama.Client
	version string
	system  string
	model   string
	style   string
}

func (a *Ask) Chat() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interactive Chat Session. Type 'exit' to quit.")

	for {
		fmt.Print("You: ")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		userInput = strings.TrimSpace(userInput)
		if strings.EqualFold(userInput, "exit") {
			fmt.Println("Exiting chat. Goodbye!")
			break
		}

		// For this simple chat, we send the user input as the prompt.
		prompt := userInput

		// Stream the response from the language model.
		fmt.Print("Assistant: ")
		err = a.api.Generate(context.Background(), &ollama.GenerateRequest{
			Model:  a.model,
			Prompt: prompt,
			System: a.system,
		}, func(res ollama.GenerateResponse) error {
			// Stream the response to stdout.
			fmt.Print(res.Response)
			return nil
		})
		if err != nil {
			fmt.Printf("\nError generating response: %s\n", err.Error())
		}
		fmt.Println()
	}
	return nil
}

func New(o *ollama.Client, version string) *Ask {
	model := viper.GetString("llm.model")
	return &Ask{
		model:   model,
		system:  system,
		style:   config.Balanced,
		api:     o,
		version: version,
	}
}
