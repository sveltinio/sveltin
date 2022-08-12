// Package toggle ...
package toggle

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	question          string
	okButtonLabel     string
	cancelButtonLabel string
	quitting          bool

	confirmation bool

	// styles
	cursor        string
	questionStyle lipgloss.Style
	buttonStyle   lipgloss.Style
	selectedStyle lipgloss.Style
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q", "n", "N":
			m.confirmation = false
			m.quitting = true
			return m, tea.Quit
		case "left", "h", "ctrl+p", "tab",
			"right", "l", "ctrl+n", "shift+tab":
			m.confirmation = !m.confirmation
		case "enter":
			m.quitting = true
			return m, tea.Quit
		case "y", "Y":
			m.quitting = true
			m.confirmation = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return ""
	}

	var ok, cancel string

	if m.confirmation {
		ok = m.selectedStyle.Render(m.okButtonLabel)
		cancel = m.buttonStyle.Render(m.cancelButtonLabel)
	} else {
		ok = m.buttonStyle.Render(m.okButtonLabel)
		cancel = m.selectedStyle.Render(m.cancelButtonLabel)
	}

	question := m.questionStyle.Render(m.question)
	prompt := lipgloss.NewStyle().SetString(m.cursor).
		Foreground(cyan).
		PaddingRight(1).
		String()
	buttons := lipgloss.JoinHorizontal(lipgloss.Left, ok, dividerStyle, cancel)
	ui := lipgloss.JoinHorizontal(lipgloss.Center, question, prompt, buttons)

	return dialogBoxStyle.Render(ui)
}
