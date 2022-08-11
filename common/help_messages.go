/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package common ...
package common

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/sveltinio/sveltin/config"
	styles "github.com/sveltinio/sveltin/internal/tui/styles"
	"github.com/sveltinio/sveltin/pkg/logger"
)

type UserProjectConfig struct {
	ProjectName   string
	CSSLibName    string
	ThemeName     string
	NPMClientName string
}

// PrintSummary shows a summary text with the project settings selected by the user.
func (uc *UserProjectConfig) PrintSummary() {
	fmt.Println(
		styles.List.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				styles.ListHeader("Here are your choices"),
				styles.ListDone("Project Name: ", uc.ProjectName),
				styles.ListDone("CSS Lib: ", uc.CSSLibName),
				styles.ListDone("Theme: ", uc.ThemeName),
				styles.ListDone("NPM Client: ", uc.NPMClientName),
			),
		),
	)
}

// PrintNextStepsHelperForNewProject prints an help message as next steps for a project creation.
func (uc *UserProjectConfig) PrintNextStepsHelperForNewProject() {
	desc := styles.List.Render(lipgloss.JoinVertical(lipgloss.Left,
		styles.ListHeader("Next Steps"),
		styles.ListNumbered("1. cd "+uc.ProjectName),
		styles.ListNumbered("2. sveltin install "+styles.Gray("(or npm run install, pnpm install, ...)")),
		styles.ListNumbered("3. sveltin server"+styles.Gray(" (or npm run dev, pnpm dev, ...)")),
	))
	additionalInfo := styles.List.Render(lipgloss.JoinVertical(lipgloss.Left,
		styles.DescStyle.Render("To stop the dev server, hit Ctrl-C"),
		styles.InfoStyle.Render("Visit the Quick Start guide"+styles.Divider+styles.URL("https://docs.sveltin.io/quick-start")),
	))
	fmt.Println(desc)
	fmt.Println(additionalInfo)
}

// PrintNextStepsHelperForNewProjectWithExistingTheme pprints an help message as next steps for a project creation with an existing theme'.
func (uc *UserProjectConfig) PrintNextStepsHelperForNewProjectWithExistingTheme() {
	desc := styles.List.Render(lipgloss.JoinVertical(lipgloss.Left,
		styles.ListHeader("Next Steps"),
		styles.ListNumbered("1. cd "+uc.ProjectName),
		styles.ListNumbered("2. sveltin install "+styles.Gray("(or npm run install, pnpm install, ...)")),
		styles.ListNumbered("3. git init"),
		styles.ListNumbered("4. git submodule add <github_repu_url_for_the_theme> themes/<theme_name>"),
		styles.ListNumbered("5. Follow the instructions on the README file from the theme creator"),
		styles.ListNumbered("6. sveltin server"+styles.Gray(" (or npm run dev, pnpm dev, ...)")),
	))
	additionalInfo := styles.List.Render(lipgloss.JoinVertical(lipgloss.Left,
		styles.DescStyle.Render("To stop the dev server, hit Ctrl-C"),
		styles.InfoStyle.Render("Visit the Theme guide"+styles.Divider+styles.URL("https://docs.sveltin.io/theming")),
	))
	fmt.Println(desc)
	fmt.Println(additionalInfo)
}

// PrintHelperTextNewMetadata prints an help message string for 'metadata creation'.
func PrintHelperTextNewMetadata(metadataInfo *config.TemplateData) {

	var exampleString string
	if metadataInfo.Type == "single" {
		exampleString = fmt.Sprintf("%s%s", toSnakeCase(metadataInfo.Name), ": your_value")
	} else {
		exampleString = toSnakeCase(metadataInfo.Name) + `:
  - value 1
  - value 2`
	}

	infoMsg := styles.List.Render(lipgloss.JoinVertical(lipgloss.Left,
		styles.TitleStyle.Render("Metadata ready to be used!"),
		styles.InfoStyle.Render("Ensure your markdown frontmatter includes it."),
		styles.NewLine,
		styles.Plain("Example:"),
		styles.NewLine,
		styles.Italic(exampleString),
	))

	fmt.Println(infoMsg)

}

// PrintHelperTextNewTheme returns an help message string for 'theme creation'.
func (uc *UserProjectConfig) PrintHelperTextNewTheme() {
	desc := styles.List.Render(lipgloss.JoinVertical(lipgloss.Left,
		styles.ListHeader("Next Steps"),
		styles.ListNumbered("1. cd "+uc.ProjectName),
		styles.ListNumbered("2. sveltin install "+styles.Gray("(or npm run install, pnpm install, ...)")),
		styles.ListNumbered("3. Create your theme components and partials"),
		styles.ListNumbered("4. sveltin server"+styles.Gray(" (or npm run dev, pnpm dev, ...)")),
	))
	additionalInfo := styles.List.Render(lipgloss.JoinVertical(lipgloss.Left,
		styles.DescStyle.Render("To stop the dev server, hit Ctrl-C"),
		styles.InfoStyle.Render("Visit the Theme guide"+styles.Divider+styles.URL("https://docs.sveltin.io/theming")),
	))
	fmt.Println(desc)
	fmt.Println(additionalInfo)
}

// PrintHelperTextDryRunFlag prints a message box for commands supporting the 'dry-run mode'.
func PrintHelperTextDryRunFlag() {
	fmt.Println(styles.Bordered("RUNNING IN DRY-RUN MODE"))
}

// PrintHelperTextDeploySummary prints a summary message string for commands like deploy.
func PrintHelperTextDeploySummary(numOfFolders, numOfFiles int) {
	folderCounter := strconv.Itoa(numOfFolders)
	filesCounter := strconv.Itoa(numOfFiles)
	fmt.Println(
		styles.List.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				styles.ListHeader("SUMMARY"),
				styles.ListDone("Total number of created folders: ", folderCounter),
				styles.ListDone("Total number of copied files: ", filesCounter),
			),
		),
	)
}

// ShowDeployCommandWarningMessages display a set of useful information for the deploy over FTP process.
func ShowDeployCommandWarningMessages(log *logger.Logger) {
	log.Printer.SetPrinterOptions(&logger.PrinterOptions{
		Timestamp: false,
		Colors:    true,
		Labels:    true,
		Icons:     true,
	})
	listLogger := log.WithList()
	listLogger.Append(logger.LevelWarning, "Create a backup of the existing content on the remote folder")
	listLogger.Append(logger.LevelWarning, "Delete existing content except what specified with --exclude flag")
	listLogger.Append(logger.LevelWarning, "Upload content to the remote folder")
	listLogger.Info("Be aware! The deploy command will perform the following actions")

}

func toSnakeCase(txt string) string {
	cleanString := strings.ToLower(txt)
	cleanString = strings.ReplaceAll(cleanString, " ", "_")
	cleanString = strings.ReplaceAll(cleanString, "-", "_")
	return cleanString
}
