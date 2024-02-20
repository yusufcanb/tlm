package suggest

import (
	"bytes"
	"fmt"
	"github.com/yusufcanb/tlama/pkg/shell"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type suggestViewModel struct {
	placeholder string
	textInput   textinput.Model
	err         error
}

func (m suggestViewModel) Init() tea.Cmd {
	m.textInput.SetValue(m.placeholder)
	return textinput.Blink
}

func (m suggestViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			cmd := shell.Exec(m.textInput.Value())
			var stdout, stderr bytes.Buffer

			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println()
			fmt.Println(stdout.String())
			fmt.Println(stderr.String())
			m.textInput.Focus()
			textinput.Blink()

			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	m.textInput.Focus()
	return m, cmd
}

func (m suggestViewModel) View() string {
	return fmt.Sprintf(
		"\n> %s\n%s",
		m.textInput.Value(),
		"\n[enter] to execute command\n[ctrl-c] to cancel",
	) + "\n"
}

func NewPromptView(prompt string) suggestViewModel {
	ti := textinput.New()
	ti.SetValue(prompt)
	ti.Focus()
	ti.CharLimit = 128
	ti.Width = 128

	return suggestViewModel{
		textInput: ti,
		err:       nil,
	}
}
