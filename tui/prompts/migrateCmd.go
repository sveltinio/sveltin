package prompts

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/sveltinio/prompti/confirm"
)

// ConfirmMigration renders a confirmation dialog box.
func ConfirmMigration() (bool, error) {
	message := `Sveltin will try to migrate as much as it can to make your
project ready with the latest introduced by Sveltin.

Remember: the main goal is to update sveltin internal files.

If something will not be handled by the migrations, comments
will be added above the affected code lines and/or you will
see errors from SvelteKit by running the server as usual.

Ensure to commit your changes to keep track of what the
command applies.`

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
		Question: "Continue?",
		Styles:   confirmStyles,
	})
}
