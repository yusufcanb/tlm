package app

import (
	"fmt"
	"os"

	ollama "github.com/jmorganca/ollama/api"

	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/pkg/shell"
)

func notFound(_ *cli.Context, _ string) {
	fmt.Println(shell.Err() + " command not found.")
	os.Exit(-1)
}

func beforeRun(o *ollama.Client) func(c *cli.Context) error {

	// fmt.Println("----DEBUG----")
	// fmt.Println(strings.Join(os.Args, " "))
	// fmt.Println("----DEBUG----")

	return func(c *cli.Context) error {
		arg := c.Args().Get(0)

		// If the command is not help or version, check if Ollama is set and up
		if arg != "v" || arg != "version" || arg != "help" || arg != "h" {
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
