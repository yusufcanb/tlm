package install

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/pkg/shell"
)

func (i *Install) deployBefore(_ *cli.Context) error {
	return shell.CheckOllamaIsUp(i.api)
}

func (i *Install) deployAction(_ *cli.Context) error {
	var err error
	var version string

	version, err = i.api.Version(context.Background())
	if err != nil {
		version = ""
		os.Exit(-1)
	}

	fmt.Println(fmt.Sprintf("Ollama version: %s\n", version))
	i.deployTlm()

	fmt.Println("\nDone..")
	return nil
}

func (i *Install) DeployCommand() *cli.Command {
	return &cli.Command{
		Name:    "deploy",
		Aliases: []string{"d"},
		Usage:   "Deploys tlm modelfiles.",
		Before:  i.deployBefore,
		Action:  i.deployAction,
	}
}
