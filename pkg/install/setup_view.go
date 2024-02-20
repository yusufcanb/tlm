package install

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"slices"
	"strings"
)

type (
	errMsg error
)

type question struct {
	question string
	answer   bool
	exitOnNo bool
}

type model struct {
	questions []question
	index     int // Index of the current question
	viewText  strings.Builder

	exited bool
}

type initialModelArgs struct {
	alreadyInstalled bool
}

var confirmText = `
- Image: ollama:latest
- Model: codellama:7b
- Volume: ollama
- GPU: %v

LLaMa will be deployed and model will be pulled for the first time.
This process might take a few minutes depending on your network speed.

[enter] to continue
[ctrl+c] to cancel`

func initialModel(args *initialModelArgs) *model {
	questions := []question{
		{question: "Enable GPU support (Only NVIDIA GPUs are supported)? [y/n]", answer: false, exitOnNo: false},
		{question: confirmText, answer: false, exitOnNo: false}, // Assuming confirmText is defined elsewhere
	}

	if args.alreadyInstalled {
		questions = slices.Insert(questions, 0, question{question: "Ollama is already deployed and running, proceed? [y/n]", answer: false, exitOnNo: true})
	}

	return &model{
		questions: questions,
		index:     0,
		viewText:  strings.Builder{},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.index >= len(m.questions) {
		return m, tea.Quit
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.index == len(m.questions)-1 { // last question
				m.questions[m.index].answer = true
				return m, tea.Quit
			}
			return m, nil
		case "y", "Y":
			if m.index == len(m.questions)-1 { // last question
				return m, nil
			}
			m.questions[m.index].answer = true
			m.nextQuestion()
			return m, nil
		case "n", "N":
			if m.index == len(m.questions)-1 { // last question
				return m, nil
			}
			m.questions[m.index].answer = false
			if m.questions[m.index].exitOnNo {
				return m, tea.Quit
			}
			m.nextQuestion()
		case "ctrl+c":
			m.exited = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.index >= len(m.questions) {
		for i := 0; i < m.index; i++ {
			q := m.questions[i]
			m.viewText.WriteString(fmt.Sprintf("%s: %v\n", q.question, q.answer))
		}
		return m.viewText.String()
	}

	for i := 0; i < m.index; i++ {
		q := m.questions[i]
		m.viewText.WriteString(fmt.Sprintf("%s: %v\n", q.question, q.answer))
	}

	q := m.questions[m.index]
	m.viewText.WriteString(fmt.Sprintf("%s\n", q.question))

	return m.viewText.String()
}

func (m *model) nextQuestion() {
	m.index++
	if m.index >= len(m.questions) {
		return // Reached the end
	}
}
