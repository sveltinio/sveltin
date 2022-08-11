package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	border  = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	success = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	// Text Styles
	Plain      = lipgloss.NewStyle().Render
	Italic     = lipgloss.NewStyle().Italic(true).Render
	Bold       = lipgloss.NewStyle().Bold(true).Render
	Underlined = lipgloss.NewStyle().Underline(true).Render
	Bordered   = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Margin(1).
			Padding(1, 2).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true).Render

	NewLine = lipgloss.NewStyle().Render("")
	// Colors
	Gray = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#969B86", Dark: "#696969"}).Render

	// Utils
	Divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(border).
		String()

	URL = lipgloss.NewStyle().Foreground(success).Render

	// Title.
	TitleStyle = lipgloss.NewStyle().
			MarginLeft(1).
			MarginRight(5).
			Padding(0, 1).
			Italic(true).
			Foreground(lipgloss.Color("#FFF7DB")).
			SetString("Lip Gloss")

	DescStyle = lipgloss.NewStyle().MarginTop(1)

	InfoStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderForeground(border)

	// List
	List = lipgloss.NewStyle().
		MarginTop(1).
		MarginRight(2).
		MarginLeft(1)

	ListHeader = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(border).
			MarginRight(2).
			Render

	ListItem = lipgloss.NewStyle().PaddingLeft(2).Render

	ListNumbered = lipgloss.NewStyle().Render

	checkMark = lipgloss.NewStyle().SetString("✓").
			Foreground(success).
			PaddingRight(1).
			String()

	ListDone = func(s, v string) string {
		return checkMark + lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#969B86", Dark: "#696969"}).
			Render(s) + Bold(v)
	}
)
