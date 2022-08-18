package confirm

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	//cyan    = lipgloss.AdaptiveColor{Light: "#4f46e5", Dark: "#c7d2fe"}
	//muted   = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	//gray    = lipgloss.AdaptiveColor{Light: "#6b7280", Dark: "#9ca3af"} // Light: gray-500, Dark: gray-400
	purple = lipgloss.AdaptiveColor{Light: "#7e22ce", Dark: "#a855f7"} // Light: purple-700, Dark: purple-500
	//pink    = lipgloss.AdaptiveColor{Light: "#ec4899", Dark: "#f472b6"} // Light: pink-500, Dark: pink-400
	neutral = lipgloss.AdaptiveColor{Light: "737373", Dark: "#a3a3a3"}  // Light: neutral-500, Dark: neutral-400
	amber   = lipgloss.AdaptiveColor{Light: "#fef3c7", Dark: "#fef3c7"} // Light: amber-100, Dark: amber-100

	// Styles
	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(purple).
			Margin(1, 0, 0, 0).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true).
			Width(50).Align(lipgloss.Center)

	questionStyle = lipgloss.NewStyle().Bold(true)

	defaultButtonStyle = lipgloss.NewStyle().
				Foreground(amber).
				Background(neutral).
				Padding(0, 3).
				Margin(1, 1)

	activeButtonStyle = defaultButtonStyle.Copy().
				Foreground(amber).
				Background(purple).
				Underline(true)
)
