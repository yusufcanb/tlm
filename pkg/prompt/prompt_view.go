package prompt

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yusufcanb/tlama/pkg/shell"
)

type (
	errMsg error
)

type promptViewModel struct {
	placeholder string
	textInput   textinput.Model
	err         error
}

func (m promptViewModel) Init() tea.Cmd {
	m.textInput.SetValue(m.placeholder)
	m.textInput.Focus()
	return textinput.Blink
}

func (m promptViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			cmd := shell.Exec(m.textInput.Value())
			stdout, err := cmd.Output()
			tea.ClearScreen()

			if err != nil {
				fmt.Println(err.Error())
				return m, tea.Quit
			}

			fmt.Println(string(stdout))
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m promptViewModel) View() string {
	return fmt.Sprintf(
		"\n💡 %s\n\n%s",
		m.textInput.Value(),
		"✔️ [enter] to execute command\n❌  [esc] to cancel",
	) + "\n"
}

func NewPromptView(prompt string) promptViewModel {
	ti := textinput.New()
	ti.SetValue(prompt)
	ti.Focus()
	ti.CharLimit = 128
	ti.Width = 128

	return promptViewModel{
		textInput: ti,
		err:       nil,
	}
}
