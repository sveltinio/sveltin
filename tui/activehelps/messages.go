package activehelps

import "github.com/charmbracelet/lipgloss"

// Hint display a message with fain style used to provide an active help on commands.
func Hint(text string) string {
	return lipgloss.NewStyle().Faint(true).Render(text)
}
