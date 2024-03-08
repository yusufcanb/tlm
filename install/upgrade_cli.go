package install

import (
	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/shell"
)

func (i *Install) upgradeBefore(_ *cli.Context) error {
	return shell.CheckOllamaIsUp(i.api)
}

func (i *Install) upgradeAction(_ *cli.Context) error {
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
