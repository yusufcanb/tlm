package explain

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/shell"
	"os"
)

func (e *Explain) before(_ *cli.Context) error {
	_, err := e.api.Version(context.Background())
	if err != nil {
		fmt.Println(shell.Err() + " " + err.Error())
		fmt.Println(shell.Err() + " Ollama connection failed. Please check your Ollama if it's running or configured correctly.")
		os.Exit(-1)
	}

	list, err := e.api.List(context.Background())
	if err != nil {
		fmt.Println(shell.Err() + " " + err.Error())
		fmt.Println(shell.Err() + " Ollama connection failed. Please check your Ollama if it's running or configured correctly.")
		os.Exit(-1)
	}

	found := false
	for _, model := range list.Models {
		if model.Name == e.modelfileName {
			found = true
			break
		}
	}

	if !found {
		fmt.Println(shell.Err() + " " + "tlm's explain model not found.\n\nPlease run `tlm deploy` to deploy tlm models first.")
		os.Exit(-1)
	}

	return nil
}

func (e *Explain) action(c *cli.Context) error {
	return e.StreamExplanationFor(Balanced, c.Args().Get(0))
}

func (e *Explain) Command() *cli.Command {
	return &cli.Command{
		Name:        "explain",
		Aliases:     []string{"e"},
		Usage:       "Explains a command.",
		UsageText:   "tlm explain <command>",
		Description: "explains given shell command.",
		Before:      e.before,
		Action:      e.action,
	}
}
