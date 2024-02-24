package install

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func (i *Install) Action(c *cli.Context) error {
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
		Usage:   "deploy tlm model files",
		Action:  i.Action,
	}
}
