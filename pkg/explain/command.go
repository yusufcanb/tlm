package explain

import (
	"github.com/urfave/cli/v2"
)

func GetCommand() *cli.Command {
	return &cli.Command{
		Name:    "explain",
		Aliases: []string{"e"},
		Usage:   "Explain a command.",
		Action:  explainAction,
	}
}
