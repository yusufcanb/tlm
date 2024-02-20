package install

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"

	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	err       error
}

func initialModel(i int) model {

	switch i {
	case 0:
		ti := textinput.New()
		ti.Placeholder = "localhost:11111"
		ti.Prompt = "Ollama Host: "
		ti.Focus()
		ti.CharLimit = 120
		ti.Width = 120

		return model{
			textInput: ti,
			err:       nil,
		}

	case 1:
		ti := textinput.New()
		ti.Placeholder = "codellama:7b"
		ti.Prompt = "Ollama Model: "
		ti.Focus()
		ti.CharLimit = 120
		ti.Width = 120

		return model{
			textInput: ti,
			err:       nil,
		}

	default:
		return model{}
	}

}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {

		case tea.KeyTab:
			m.textInput.SetValue(m.textInput.Placeholder)
			return m, nil

		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
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

func (m model) View() string {
	return fmt.Sprintf(
		"%s",
		m.textInput.View(),
	) + "\n"
}
