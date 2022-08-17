package styles

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	border  = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	success = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	muted   = lipgloss.AdaptiveColor{Light: "#969B86", Dark: "#696969"}

	// Plain sets an as it is formatting rule for the text.
	Plain = lipgloss.NewStyle().Render
	// Italic sets an italic formatting rule for the text.
	Italic = lipgloss.NewStyle().Italic(true).Render
	// Bold sets a bold formatting rule for the text.
	Bold = lipgloss.NewStyle().Bold(true).Render
	// Gray sets foregrund color to gray for the text.
	Gray = lipgloss.NewStyle().Foreground(muted).Render
	// Bordered prints a bordered text.
	Bordered = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Margin(1).
			Padding(1, 2).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true).Render

	actionStyle       = lipgloss.NewStyle().Margin(2, 0, 1, 1)
	actionHeaderStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				BorderBottom(true).
				BorderForeground(border).
				MarginRight(2).
				Render

	// Title is used to style the heading text.
	Title = func(s string) string {
		return actionStyle.Render(actionHeaderStyle(s))
	}
	// ActionResult is used to style an action result text.
	ActionResult = lipgloss.NewStyle().
			MarginLeft(1).
			MarginRight(5).
			Padding(0, 1).
			Italic(true).
			Foreground(lipgloss.Color("#FFF7DB")).Render
	// Desc is used to style a description text.
	Desc = lipgloss.NewStyle().MarginTop(1).Render
	// Info is used to style an info text.
	Info = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderTop(true).
		BorderForeground(border).Render

	// List is used to style a list.
	List = lipgloss.NewStyle().Margin(2, 0, 1, 1)
	// ListHeader is used to style a list header text.
	ListHeader = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(border).
			MarginRight(2).
			Render
	// ListNumbered is used to style a numbered list item.
	ListNumbered = lipgloss.NewStyle().Render

	checkMark = lipgloss.NewStyle().SetString("✓").
			Foreground(success).
			PaddingRight(1).
			String()
	// ListDone is used to style a list item with a check mark as prefix.
	ListDone = func(s, v string) string {
		return checkMark + lipgloss.NewStyle().
			Foreground(muted).
			Render(s) + Bold(v)
	}

	// NewLine prints an empty line.
	NewLine = lipgloss.NewStyle().Render("")
	// Divider is used to style a divider char.
	Divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(border).
		String()
	// Url is used to style an URL string.
	Url = lipgloss.NewStyle().Foreground(success).Render
)
