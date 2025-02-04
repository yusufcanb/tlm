package ask

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/pkg/rag"
)

func (a *Ask) before(c *cli.Context) error {
	return nil
}

func (a *Ask) action(c *cli.Context) error {
	contextDir := c.Path("context")
	var ctx string // chat context
	var dir string

	if contextDir != "" {
		includePatterns := c.StringSlice("include")
		excludePatterns := c.StringSlice("exclude")
		dir, ctx = rag.GetContext(contextDir, includePatterns, excludePatterns)
		fmt.Printf("\nContext Files\n--------------------\n%s\n", dir)
	}

	message := c.Args().First()
	if message == "" {
		return errors.New("message is required")
	}

	rag := rag.NewRAGChat(a.api, ctx)
	_, err := rag.Send(message)
	if err != nil {
		return err
	}

	return nil
}

func (a *Ask) after(c *cli.Context) error {
	return nil
}

func (a *Ask) Command() *cli.Command {
	return &cli.Command{
		Name:    "ask",
		Usage:   "Asks a question",
		Aliases: []string{"a"},
		Action:  a.action,
		Before:  a.before,
		After:   a.after,
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:    "context",
				Aliases: []string{"c"},
				Usage:   "Context directory",
			},
			&cli.StringSliceFlag{
				Name:    "include",
				Aliases: []string{"i"},
				Usage:   "Include patterns. E.g. --include=*.txt",
			},
			&cli.StringSliceFlag{
				Name:    "exclude",
				Aliases: []string{"e"},
				Usage:   "Exclude patterns. E.g. --exclude=*.binary",
			},
			&cli.BoolFlag{
				Name:    "interactive",
				Aliases: []string{"it"},
				Usage:   "enable interactive chat mode",
			},
		},
	}
}
