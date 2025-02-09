package ask

import (
	"errors"
	"fmt"
	"os/user"

	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/pkg/packer"
	"github.com/yusufcanb/tlm/pkg/rag"
)

func (a *Ask) before(c *cli.Context) error {
	message := c.Args().First()
	if message == "" {
		return errors.New("message is required")
	}

	user, err := user.Current()
	if err != nil {
		a.user = "User"
	}
	a.user = user.Username

	return nil
}

func (a *Ask) action(c *cli.Context) error {
	isInteractive := c.Bool("interactive")
	contextDir := c.Path("context")
	var chatContext string    // chat context
	var numCtx int = 1024 * 8 // num_ctx in Ollama API

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

	fmt.Printf("\n🤖 %s\n───────────────────\n", a.model)

	message := c.Args().First()
	rag := rag.NewRAGChat(a.api, chatContext)
	_, err := rag.Send(message, numCtx)
	if err != nil {
		return err
	}

	user.Current()

	if isInteractive {
		for {
			fmt.Printf("\n\n👤 %s\n───────────────────\n", a.user)
			var input string
			fmt.Scanln(&input)

			if input == "exit" {
				break
			}

			_, err := rag.Send(input, numCtx)
			if err != nil {
				return err
			}

			fmt.Printf("\n\n")
		}
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
				Usage:   "Include patterns. e.g. --include=*.txt",
			},
			&cli.StringSliceFlag{
				Name:    "exclude",
				Aliases: []string{"e"},
				Usage:   "Exclude patterns. e.g. --exclude=*.binary",
			},
			&cli.BoolFlag{
				Name:    "interactive",
				Aliases: []string{"it"},
				Usage:   "enable interactive chat mode",
			},
		},
	}
}
