package suggest

import (
	"github.com/urfave/cli/v2"
)

func GetCommand() *cli.Command {
	return &cli.Command{
		Name:    "suggest",
		Aliases: []string{"s"},
		Usage:   "Suggest a command.",
		Action:  promptAction,
	}
}
