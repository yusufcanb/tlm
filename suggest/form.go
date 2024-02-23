package suggest

import (
	"github.com/charmbracelet/huh"
)

type Action int

const (
	Cancel Action = iota
	Execute
	Explain
)

type CommandForm struct {
	command string
	action  Action
}

func (s *CommandForm) Run() error {
	group := huh.NewGroup(
		huh.NewInput().
			Value(&s.command),

		huh.NewSelect[Action]().
			Options(
				huh.NewOption("Execute", Execute),
				huh.NewOption("Explain", Explain),
				huh.NewOption("Cancel", Cancel),
			).
			Value(&s.action),
	)

	form := huh.NewForm(group).
		WithTheme(huh.ThemeBase16()).
		WithKeyMap(huh.NewDefaultKeyMap())

	return form.Run()
}

func NewCommandForm(command string) *CommandForm {
	return &CommandForm{command: command}
}
