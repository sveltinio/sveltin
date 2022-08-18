package logger

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	timestampStyle = lipgloss.NewStyle().Foreground(muted).Render
	iconStyle      = func(color lipgloss.AdaptiveColor, s string) string {
		return lipgloss.NewStyle().Foreground(color).Render(s)
	}

	levelStyle = func(color lipgloss.AdaptiveColor, s string) string {
		return lipgloss.NewStyle().Padding(0).Foreground(color).Render(s)
	}

	textStyle = func(color lipgloss.AdaptiveColor, s string) string {
		return lipgloss.NewStyle().Foreground(color).Render(s)
	}

	// List Styles
	listTitleStyle = lipgloss.NewStyle().Margin(0, 0, 1, 0).
			Padding(0, 1).
			Italic(true).Underline(true).Render

	listItemStyle = func(s, indentChar string, indentSize int) string {
		return lipgloss.NewStyle().PaddingLeft(indentSize).Render(indentChar + s)
	}
)
