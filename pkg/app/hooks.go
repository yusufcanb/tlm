package app

import (
	"fmt"
	"os"

	ollama "github.com/jmorganca/ollama/api"

	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/pkg/shell"
)

func shouldCheckOllamaIsSetandUp(commandName string) bool {
	switch commandName {
	case "ask", "a", "explain", "e", "suggest", "s":
		return true
	default:
		return false
	}
}

func notFound(_ *cli.Context, _ string) {
	fmt.Println(shell.Err() + " command not found.")
	os.Exit(-1)
}

func beforeRun(o *ollama.Client) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		arg := c.Args().Get(0)

		if shouldCheckOllamaIsSetandUp(arg) {
			var err error

			err = shell.CheckOllamaIsSet()
			if err != nil {
				fmt.Println(shell.Err() + " " + err.Error())
				os.Exit(-1)
			}

			err = shell.CheckOllamaIsUp(o)
			if err != nil {
				fmt.Println(shell.Err() + " " + err.Error())
				os.Exit(-1)
			}
		}

		return nil
	}
}

func afterRun(version string) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		switch c.Args().Get(0) {
		// ignore update checks for suggest, explain or ask commands
		case "suggest", "s", "explain", "e", "ask", "a":
			return nil

		default:
			rm := c.App.Metadata["releaseManager"].(*ReleaseManager) // Get the ReleaseManager from the app's metadata
			rm.CheckForUpdates(version)
			return nil
		}
	}
}
