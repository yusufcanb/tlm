package setup

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
)

func GetCommand() *cli.Command {
	return &cli.Command{
		Name:  "setup",
		Usage: "Helper for installing tlama.",

		Action: func(c *cli.Context) error {
			m := viewModel{}
			m.resetSpinner()

			if _, err := tea.NewProgram(m).Run(); err != nil {
				fmt.Println("could not run program:", err)
				os.Exit(1)
			}
			return nil
		},
	}
}
