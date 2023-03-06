package prompts

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/sveltinio/prompti/confirm"
	"github.com/sveltinio/sveltin/internal/markup"
)

// ConfirmDeploy renders a confirmation dialog box.
func ConfirmDeploy(isDryRun, isBackup bool) (bool, error) {
	var strBuilder strings.Builder

	strBuilder.WriteString("The deploy command will perform the following actions:\n\n")
	if isBackup {
		strBuilder.WriteString("Create a backup of the existing remote content;\n")
	}

	strBuilder.WriteString("Delete existing content if not\n --exclude or --withExcludeFile flags are used;\n")
	strBuilder.WriteString("Upload content to the remote folder.")

	if isDryRun {
		strBuilder.WriteString(fmt.Sprintf("\n\n%s\nNothing will happen, you are just simulating the process!", markup.Underline("DRY-RUN Mode")))
	}

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
		Message:  strBuilder.String(),
		Question: "Continue?",
		Styles:   confirmStyles,
	})
}
