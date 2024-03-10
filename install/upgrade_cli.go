package install

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func (i *Install) upgradeBefore(_ *cli.Context) error {
	return nil
}

func (i *Install) upgradeAction(_ *cli.Context) error {
	fmt.Println("Upgrading tlm to latest version.")

	return nil
}

func (i *Install) UpgradeCommand() *cli.Command {
	return &cli.Command{
		Name:    "upgrade",
		Aliases: []string{"u"},
		Usage:   "Upgrades tlm to latest version.",
		Before:  i.upgradeBefore,
		Action:  i.upgradeAction,
	}
}
