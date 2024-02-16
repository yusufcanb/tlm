package prompt

import (
	"fmt"
	"log"
	"os"
	"sync"
	"tlama/pkg/api"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
)

func PromptAction(c *cli.Context) error {
	var wg sync.WaitGroup

	var program *tea.Program
	var command string
	var err error

	prompt := c.String("prompt")
	if prompt == "" {
		fmt.Println(c.Command.Usage)
		os.Exit(-1)
	}

	wg.Add(1)

	go func() {
		program = tea.NewProgram(NewRequestView())
		_, err = program.Run()
		if err != nil {
			log.Fatalf("could not run program: %s", err)
		}
		defer program.Quit()
	}()

	go func() {
		command, err = c.App.Metadata["api"].(*api.OllamaAPI).Generate(c.String("prompt"))
		wg.Done()
		defer program.Quit()
	}()

	wg.Wait()

	if err != nil {
		log.Fatal("Couldn't get the prompt from local LLM.", err)
	}

	if command == "" {
		log.Fatal("Prompt is empty.")
	}

	if _, err := tea.NewProgram(NewPromptView(command)).Run(); err != nil {
		log.Fatalf("could not run program: %s", err)
	}

	return nil
}
