package explain

import (
	"github.com/urfave/cli/v2"
)

func (e *Explain) before(_ *cli.Context) error {
	return nil
}

func (e *Explain) action(c *cli.Context) error {
	return e.StreamExplanationFor(Balanced, c.Args().Get(0))
}

func (e *Explain) Command() *cli.Command {
	return &cli.Command{
		Name:        "explain",
		Aliases:     []string{"e"},
		Usage:       "Explains a command.",
		UsageText:   "tlm explain <command>",
		Description: "explains given shell command.",
		Before:      e.before,
		Action:      e.action,
	}
}
