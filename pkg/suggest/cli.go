package suggest

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/pkg/explain"
	"github.com/yusufcanb/tlm/pkg/shell"
)

func (s *Suggest) before(c *cli.Context) error {
	prompt := c.Args().First()
	if prompt == "" {
		cli.ShowSubcommandHelp(c)
		return cli.Exit("", -1)
	}

	overrideModel := c.String("model")
	if overrideModel != "" {
		s.model = overrideModel
	}

	overrideStyle := c.String("style")
	if overrideStyle != "" {
		s.style = overrideStyle
	}

	return nil
}

func (s *Suggest) action(c *cli.Context) error {
	var responseText string
	var err error

	var t1, t2 time.Time

	prompt := c.Args().Get(0)
	spinner.New().
		Type(spinner.Line).
		Title(fmt.Sprintf(" %s is thinking... ", s.model)).
		Style(lipgloss.NewStyle().Foreground(lipgloss.Color("2"))).
		Action(func() {
			t1 = time.Now()
			responseText, err = s.getCommandSuggestionFor(viper.GetString("shell"), prompt)
			t2 = time.Now()
		}).
		Run()

	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("model '%s' not found", s.model)) {
			fmt.Println(fmt.Sprintf(shell.Err()+" %s - run `ollama pull %s` to download the model.", err.Error(), s.model))
			return nil
		}

		cli.Exit(fmt.Sprintf("error getting suggestion: %s", err.Error()), -1)
	}

	fmt.Printf(shell.SuccessMessage("┃ >")+" %s is thinking... (%s) \n", s.model, t2.Sub(t1).String())
	if len(s.extractCommandsFromResponse(responseText)) == 0 {
		fmt.Println(shell.WarnMessage("┃ >") + " No command found for given prompt..." + "\n")
		return nil
	}

	form := NewCommandForm(s.extractCommandsFromResponse(responseText)[0])
	err = form.Run()

	fmt.Println(shell.SuccessMessage("┃ > ") + form.command)
	if err != nil {
		fmt.Println(shell.WarnMessage("┃ > ") + "Aborted..." + "\n")
		return nil
	}

	if form.action == Execute {
		fmt.Println(shell.SuccessMessage("┃ > ") + "Executing..." + "\n")
		cmd, stdout, stderr := shell.Exec2(form.command)
		err = cmd.Run()
		if err != nil {
			fmt.Println(stderr.String())
			return nil
		}

		if stderr.String() != "" {
			fmt.Println(stderr.String())
			return nil
		}

		fmt.Println(stdout.String())
		return nil
	}

	if form.action == Explain {
		fmt.Println(shell.SuccessMessage("┃ > ") + fmt.Sprintf("%s is explaining...", s.model) + "\n")

		exp := explain.New(s.api, s.version)
		err = exp.StreamExplanationFor(s.style, form.command)
		if err != nil {
			return cli.Exit(err, -1)
		}

	} else {
		fmt.Println(shell.WarnMessage("┃ > ") + "Aborted..." + "\n")
	}

	return nil
}

func (s *Suggest) Command() *cli.Command {

	model := viper.GetString("llm.model")
	style := viper.GetString("llm.suggest")

	var overrideModel *string
	var overrideStyle *string

	return &cli.Command{
		Name:        "suggest",
		Aliases:     []string{"s"},
		Usage:       "Suggests a command.",
		UsageText:   "tlm suggest <prompt> \ntlm suggest --model=llama3.2:1b <prompt>\ntlm suggest --model=llama3.2:1b --style=<stable|balanced|creative> <prompt>",
		Description: "suggests a command for given prompt.",
		Before:      s.before,
		Action:      s.action,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "model",
				Aliases:     []string{"m"},
				Usage:       "override the model for command suggestion.",
				DefaultText: model,
				Destination: overrideModel,
			},
			&cli.StringFlag{
				Name:        "style",
				Aliases:     []string{"s"},
				Usage:       "override the style for command suggestion.",
				DefaultText: style,
				Destination: overrideStyle,
			},
		},
	}
}
