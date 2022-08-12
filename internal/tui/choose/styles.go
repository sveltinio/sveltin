package choose

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	cyan = lipgloss.AdaptiveColor{Light: "#4f46e5", Dark: "#c7d2fe"}

	// Styles
	titleStyle        = lipgloss.NewStyle().MarginTop(1).Underline(true)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(1)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(cyan)
)
