package app

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/install"
	"github.com/yusufcanb/tlm/shell"
	"os"
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
		case "suggest", "explain", "upgrade":
			return nil

		default:
			return ins.ReleaseManager.CheckForUpdates(version)
		}
	}

}
