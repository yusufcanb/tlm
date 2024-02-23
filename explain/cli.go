package explain

import (
	"github.com/urfave/cli/v2"
)

func (e *Explain) Action(c *cli.Context) error {
	return e.streamExplanationFor(Balanced, c.Args().Get(0))
}

func (e *Explain) Command() *cli.Command {
	return &cli.Command{
		Name:    "explain",
		Aliases: []string{"e"},
		Usage:   "explain a command.",
		Action:  e.Action,
	}
}
