package install

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/shell"
	"os"
)

func (i *Install) before(_ *cli.Context) error {
	_, err := i.api.Version(context.Background())
	if err != nil {
		fmt.Println(shell.Err() + " " + err.Error())
		fmt.Println(shell.Err() + " Ollama connection failed. Please check your Ollama if it's running or configured correctly.")
		os.Exit(-1)
	}

	return nil
}

func (i *Install) action(_ *cli.Context) error {
	var err error
	var version string

	version, err = i.api.Version(context.Background())
	if err != nil {
		version = ""
		os.Exit(-1)
	}

	fmt.Println(fmt.Sprintf("Ollama version: %s\n", version))

	i.deployTlm(i.suggestModelfile, i.explainModelfile)

	fmt.Println("\nDone..")
	return nil
}

func (i *Install) Command() *cli.Command {
	return &cli.Command{
		Name:    "deploy",
		Aliases: []string{"d"},
		Usage:   "Deploys tlm modelfiles.",
		Before:  i.before,
		Action:  i.action,
	}
}
