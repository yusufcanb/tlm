package ask

import (
	"fmt"
	"os/user"

	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/pkg/chroma"
	"github.com/yusufcanb/tlm/pkg/rag"
)

var usageText string = `tlm ask "<prompt>" # ask a question
tlm ask --context . --include *.md "<prompt>" # ask a question with a context`

func (a *Ask) beforeAction(c *cli.Context) error {
	prompt := c.Args().First()
	if prompt == "" {
		cli.ShowSubcommandHelp(c)
		return cli.Exit("", -1)
	}

	overrideModel := c.String("model")
	if overrideModel != "" {
		a.model = overrideModel
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
	var numCtx int = 1024 * 8 // num_ctx in Ollama API

	fmt.Printf("\n🤖 %s\n───────────────────\n", a.model)

	chromaClient := chroma.NewChromaClient("http://localhost:8000")

	prompt := c.Args().First()
	rag := rag.NewRAGChat(a.api, chromaClient, "", a.model)
	_, err := rag.Send(prompt, numCtx)
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

func (a *Ask) afterAction(c *cli.Context) error {
	return nil
}

func (a *Ask) Command() *cli.Command {
	model := viper.GetString("llm.model")

	return &cli.Command{
		Name:      "ask",
		Usage:     "Asks a question (beta)",
		UsageText: usageText,
		Aliases:   []string{"a"},
		Action:    a.action,
		Before:    a.beforeAction,
		After:     a.afterAction,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "interactive",
				Aliases: []string{"it"},
				Usage:   "enable interactive chat mode",
			},
			&cli.StringFlag{
				Name:        "model",
				Aliases:     []string{"m"},
				Usage:       "override the model for command suggestion.",
				DefaultText: model,
			},
		},
	}
}
