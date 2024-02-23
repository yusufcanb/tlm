package views

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type requestViewModel struct {
	spinner  spinner.Model
	quitting bool
	err      error
}

func (m requestViewModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m requestViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEscape, tea.KeyCtrlC:
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m requestViewModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n%s Thinking...\n", m.spinner.View())
	if m.quitting {
		return str + "\n"
	}
	return str
}

func NewRequestView() *requestViewModel {
	s := spinner.New()
	s.Spinner = spinner.Line
	return &requestViewModel{spinner: s}
}
