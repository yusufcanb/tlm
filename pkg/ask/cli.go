package ask

import (
	"errors"

	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/pkg/packer"
	"github.com/yusufcanb/tlm/pkg/rag"
)

func (a *Ask) before(c *cli.Context) error {
	message := c.Args().First()
	if message == "" {
		return errors.New("message is required")
	}

	return nil
}

func (a *Ask) action(c *cli.Context) error {
	contextDir := c.Path("context")
	var chatContext string // chat context

	if contextDir != "" {
		includePatterns := c.StringSlice("include")
		excludePatterns := c.StringSlice("exclude")

		// fmt.Printf("include=%v, exclude=%v\n\n", includePatterns, excludePatterns)

		// Pack files under the context directory
		packer := packer.New()
		res, err := packer.Pack(contextDir, includePatterns, excludePatterns)
		if err != nil {
			return err
		}

		// Sort the files by the number of tokens
		packer.PrintTopFiles(res, 5)

		// Print the context summary
		packer.PrintContextSummary(res)

		// Render the packer result
		chatContext, err = packer.Render(res)
		if err != nil {
			return err
		}
	}

	message := c.Args().First()
	rag := rag.NewRAGChat(a.api, chatContext)
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
