package doctor

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
)

func GetCommand() *cli.Command {
	return &cli.Command{
		Name:  "doctor",
		Usage: "Diagnose and fix common problems.",
		Action: func(c *cli.Context) error {
			p := tea.NewProgram(initialModel())
			if _, err := p.Run(); err != nil {
				log.Fatal(err)
			}
			return nil
		},
	}
}
