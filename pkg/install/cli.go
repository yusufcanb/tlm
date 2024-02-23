package install

import (
	"github.com/urfave/cli/v2"
)

func (i *Install) Action(c *cli.Context) error {
	return NewInstallForm2().Run()
}

func (i *Install) Command() *cli.Command {
	return &cli.Command{
		Name:    "install",
		Aliases: []string{"i"},
		Usage:   "deploy CodeLLaMa to your system.",
		Action:  i.Action,
	}
}
