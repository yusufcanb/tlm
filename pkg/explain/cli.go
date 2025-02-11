package explain

import (
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

func (e *Explain) beforeAction(c *cli.Context) error {
	prompt := c.Args().First()
	if prompt == "" {
		cli.ShowSubcommandHelp(c)
		return cli.Exit("", -1)
	}

	overrideModel := c.String("model")
	if overrideModel != "" {
		e.model = overrideModel
	}

	overrideStyle := c.String("style")
	if overrideStyle != "" {
		e.style = overrideStyle
	}

	return nil
}

func (e *Explain) action(c *cli.Context) error {
	return e.StreamExplanationFor(e.style, c.Args().Get(0))
}

func (e *Explain) afterAction(c *cli.Context) error {
	return nil
}

func (e *Explain) Command() *cli.Command {

	model := viper.GetString("llm.model")
	style := viper.GetString("llm.explain")

	return &cli.Command{
		Name:        "explain",
		Aliases:     []string{"e"},
		Usage:       "Explains a command.",
		UsageText:   "tlm explain <command> \ntlm explain --model=llama3.2:1b <command>\ntlm explain --model=llama3.2:1b --style=<stable|balanced|creative> <command>",
		Description: "explains given shell command.",
		Before:      e.beforeAction,
		Action:      e.action,
		After:       e.afterAction,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "model",
				Aliases:     []string{"m"},
				Usage:       "override the model for command suggestion.",
				DefaultText: model,
			},
			&cli.StringFlag{
				Name:        "style",
				Aliases:     []string{"s"},
				Usage:       "override the style for command suggestion.",
				DefaultText: style,
			},
		},
	}
}
