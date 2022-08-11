/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package confirm ...
package confirm

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

	var aff, neg string

	if m.confirmation {
		aff = m.selectedStyle.Render(m.okButtonLabel)
		neg = m.buttonStyle.Render(m.cancelButtonLabel)
	} else {
		aff = m.buttonStyle.Render(m.okButtonLabel)
		neg = m.selectedStyle.Render(m.cancelButtonLabel)
	}

	question := m.questionStyle.Render(m.question)
	buttons := lipgloss.JoinHorizontal(lipgloss.Left, aff, neg)
	ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)

	return lipgloss.JoinVertical(lipgloss.Center, dialogBoxStyle.Render(ui))
}
