package install

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
)

func (i *Install) Action(c *cli.Context) error {
	var err error
	var version string

	version, err = i.api.Version(context.Background())
	if err != nil {
		version = ""
	}

	form := NewInstallForm2(version, i.suggestModelfile, i.explainModelfile)
	err = form.Run()
	if err != nil {
		return err
	}

	err = i.installAndConfigureOllama(form)
	if err != nil {
		return err
	}

	fmt.Println("\nInstallation has been completed..")
	fmt.Println("\nStart using it by;\ntlm suggest \"list all files in cwd\"\n")
	return nil
}

func (i *Install) Command() *cli.Command {
	return &cli.Command{
		Name:    "install",
		Aliases: []string{"i"},
		Usage:   "deploy CodeLLaMa to your system.",
		Action:  i.Action,
	}
}
