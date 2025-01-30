package app

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/pkg/install"
	"github.com/yusufcanb/tlm/pkg/shell"
)

func notFound(_ *cli.Context, _ string) {
	fmt.Println(shell.Err() + " command not found.")
	os.Exit(-1)
}

func beforeRun() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		return nil
	}
}

func afterRun(ins *install.Install, version string) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		switch c.Args().Get(0) {
		case "suggest", "s", "explain", "e":
			return nil

		default:
			return ins.ReleaseManager.CheckForUpdates(version)
		}
	}

}
