package prompts

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/sveltinio/prompti/confirm"
	"github.com/sveltinio/sveltin/internal/markup"
)

// ContinueOnUpdateAvailable display a dialog box alerting about newer release available.
func ContinueOnUpdateAvailable(current, latest string) (bool, error) {
	infoText := `Update available! %s ➜ %s
Release Notes: https://docs.sveltin.io/release-notes

Check your installation method to update, then run
"sveltin migrate".`
	message := fmt.Sprintf(infoText, markup.Red(current), markup.Green(latest))

	// Light: orange-500, Dark: orange-400
	orange := lipgloss.AdaptiveColor{Light: "#f97316", Dark: "#fb923c"}
	// Light: blank, Dark: white
	blankAndWhite := lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}

	confirmStyles := confirm.Styles{
		Width:       70,
		BorderColor: orange,
		ActiveButtonStyle: lipgloss.NewStyle().Padding(0, 3).
			Margin(1, 1).Background(orange).Foreground(blankAndWhite),
	}

	return confirm.Run(&confirm.Config{
		Message:  message,
		Question: "Continue without updating?",
		Styles:   confirmStyles,
	})
}
