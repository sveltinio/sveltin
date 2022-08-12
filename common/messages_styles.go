package common

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Colors
	border  = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	success = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	muted   = lipgloss.AdaptiveColor{Light: "#969B86", Dark: "#696969"}

	// Text Styles
	plain  = lipgloss.NewStyle().Render
	italic = lipgloss.NewStyle().Italic(true).Render
	bold   = lipgloss.NewStyle().Bold(true).Render
	gray   = lipgloss.NewStyle().Foreground(muted).Render

	bordered = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Margin(1).
			Padding(1, 2).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true).Render

	// Misc
	newLine = lipgloss.NewStyle().Render("")
	divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(border).
		String()
	url = lipgloss.NewStyle().Foreground(success).Render

	// Title.
	titleStyle = lipgloss.NewStyle().
			MarginLeft(1).
			MarginRight(5).
			Padding(0, 1).
			Italic(true).
			Foreground(lipgloss.Color("#FFF7DB"))

	descStyle = lipgloss.NewStyle().MarginTop(1)

	infoStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderForeground(border)

	// List
	list = lipgloss.NewStyle().Margin(1, 2, 1, 1)

	listHeader = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(border).
			MarginRight(2).
			Render

	listNumbered = lipgloss.NewStyle().Render

	checkMark = lipgloss.NewStyle().SetString("✓").
			Foreground(success).
			PaddingRight(1).
			String()

	listDone = func(s, v string) string {
		return checkMark + lipgloss.NewStyle().
			Foreground(muted).
			Render(s) + bold(v)
	}
)
