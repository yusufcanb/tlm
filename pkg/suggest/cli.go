package suggest

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/pkg/explain"
	"github.com/yusufcanb/tlm/pkg/shell"
)

func (s *Suggest) before(_ *cli.Context) error {
	_, err := s.api.Version(context.Background())
	if err != nil {
		fmt.Println(shell.Err() + " " + err.Error())
		fmt.Println(shell.Err() + " Ollama connection failed. Please check your Ollama if it's running or configured correctly.")
		os.Exit(-1)
	}

	// list, err := s.api.List(context.Background())
	if err != nil {
		fmt.Println(shell.Err() + " " + err.Error())
		fmt.Println(shell.Err() + " Ollama connection failed. Please check your Ollama if it's running or configured correctly.")
		os.Exit(-1)
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
		fmt.Println(shell.Err()+" error getting suggestion:", err)
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
		err = exp.StreamExplanationFor(Stable, form.command)
		if err != nil {
			return err
		}

	} else {
		fmt.Println(shell.WarnMessage("┃ > ") + "Aborted..." + "\n")
	}

	return nil
}

func (s *Suggest) Command() *cli.Command {

	model := viper.GetString("llm.model")
	style := viper.GetString("llm.suggest")

	var overridedModel *string // FIXME implement override

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
				Destination: overridedModel,
			},
			&cli.StringFlag{
				Name:        "style",
				Aliases:     []string{"s"},
				Usage:       "override the style for command suggestion.",
				DefaultText: style,
				Destination: overridedModel,
			},
		},
	}
}
