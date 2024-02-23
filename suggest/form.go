package suggest

import (
	"github.com/charmbracelet/huh"
)

type CommandForm struct {
	command string
	confirm bool
}

func (s *CommandForm) Run() error {
	group := huh.NewGroup(
		huh.NewInput().
			Value(&s.command),
		huh.NewConfirm().
			Value(&s.confirm).
			Affirmative("execute").
			Negative("abort").
			WithHeight(1),
	)

	form := huh.NewForm(group).
		WithTheme(huh.ThemeBase16()).
		WithKeyMap(huh.NewDefaultKeyMap())

	return form.Run()
}

func NewCommandForm(command string) *CommandForm {
	return &CommandForm{command: command, confirm: true}
}
