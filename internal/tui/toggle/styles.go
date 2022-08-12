package toggle

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	cyan  = lipgloss.AdaptiveColor{Light: "#4f46e5", Dark: "#c7d2fe"}
	muted = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

	// Styles
	dialogBoxStyle = lipgloss.NewStyle().
			Margin(1, 0, 0, 0).
			Padding(1, 0)

	questionStyle = lipgloss.NewStyle().Bold(true).Padding(0, 1)

	dividerStyle = lipgloss.NewStyle().
			SetString("/").
			Padding(0, 1).
			Foreground(muted).
			String()

	defaultButtonStyle = lipgloss.NewStyle().
				Foreground(muted)

	activeButtonStyle = defaultButtonStyle.Copy().
				Foreground(cyan).
				Underline(true)
)
