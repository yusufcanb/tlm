package doctor

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"os/exec"
)

type (
	errMsg error
)

type viewModel struct {
	prompt    string
	textInput textinput.Model
	err       error
}

func initialModel() viewModel {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return viewModel{
		textInput: ti,
		err:       nil,
	}
}

func (m viewModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m viewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			cmd := exec.Command("powershell.exe", "-c", m.textInput.Value())
			stdout, err := cmd.Output()

			if err != nil {
				fmt.Println(err.Error())
				return m, tea.Quit
			}
			tea.ClearScreen()
			fmt.Println(string(stdout))
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m viewModel) View() string {
	return fmt.Sprintf(
		"Would you like to execute the command?\n\n%s\n\n%s",
		m.textInput.View(),
		"✔️ (enter to execute)\n❌  (esc to cancel)",
	) + "\n"
}
