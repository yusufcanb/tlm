package suggest

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlm/shell"
	"time"
)

func (s *Suggest) Action(c *cli.Context) error {
	var responseText string
	var err error

	var t1, t2 time.Time

	prompt := c.Args().Get(0)
	spinner.New().
		Type(spinner.Line).
		Title(" Thinking...").
		Style(lipgloss.NewStyle().Foreground(lipgloss.Color("2"))).
		Action(func() {
			t1 = time.Now()
			responseText, err = s.getCommandSuggestionFor(Stable, viper.GetString("shell"), prompt)
			t2 = time.Now()
		}).
		Run()

	if err != nil {
		fmt.Println(shell.Err()+" error getting suggestion:", err)
	}

	fmt.Printf("┃ >"+" Thinking... %s\n", shell.SuccessMessage("("+t2.Sub(t1).String()+")"))
	form := NewCommandForm(s.extractCommandsFromResponse(responseText)[0])
	err = form.Run()

	if err != nil {
		fmt.Println(shell.Err() + " " + err.Error())
	}

	fmt.Println("┃ > " + form.command + "\n")
	if form.confirm {
		cmd, stdout, stderr := shell.Exec2(form.command)
		cmd.Run()

		if stderr.String() != "" {
			fmt.Println(stderr.String())
			return errors.New("command failed")
		}

		fmt.Println(stdout.String())
		return nil
	}

	fmt.Println("suggestion elapsed time:", t2.Sub(t1))
	return nil
}

func (s *Suggest) Command() *cli.Command {
	return &cli.Command{
		Name:    "suggest",
		Aliases: []string{"s"},
		Usage:   "suggest a command.",
		Action:  s.Action,
	}
}
