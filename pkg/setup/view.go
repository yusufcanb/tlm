package setup

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	textStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Render
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("98"))
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render
)

type viewModel struct {
	index   int
	spinner spinner.Model
}

func (m viewModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m viewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		default:
			return m, nil
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	default:
		return m, nil
	}
}

func (m *viewModel) resetSpinner() {
	m.spinner = spinner.New()
	m.spinner.Style = spinnerStyle
	m.spinner.Spinner = spinner.Monkey
}

func (m viewModel) View() (s string) {
	var gap string
	switch m.index {
	case 1:
		gap = ""
	default:
		gap = " "
	}

	s += fmt.Sprintf("\n %s%s%s\n\n", m.spinner.View(), gap, textStyle("Interacting with CodeLLama2..."))
	s += helpStyle("❌  (esc|q to cancel)\n")
	return
}
