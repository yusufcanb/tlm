package setup

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
	"os"
)

func GetCommand() *cli.Command {
	return &cli.Command{
		Name:    "setup",
		Aliases: []string{"s"},
		Usage:   "setup app",

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
