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

	"github.com/charmbracelet/lipgloss"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/internal/styles"
	"github.com/sveltinio/sveltin/pkg/logger"
	"github.com/sveltinio/sveltin/utils"
)

// UserProjectConfig represents the user selections when creating a new sveltin project.
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
	nextSteps := styles.List.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			styles.ListHeader("Next Steps"),
			styles.ListNumbered("1. cd "+uc.ProjectName),
			styles.ListNumbered("2. sveltin install "+styles.Gray("(or npm run install, pnpm install, ...)")),
			styles.ListNumbered("3. sveltin server"+styles.Gray(" (or npm run dev, pnpm dev, ...)")),
		))
	infoMsg := styles.List.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			styles.Desc("To stop the dev server, hit Ctrl-C"),
			styles.Info("Visit the Quick Start guide"+styles.Divider+styles.Url("https://docs.sveltin.io/quick-start")),
		))
	fmt.Println(nextSteps)
	fmt.Println(infoMsg)
}

// PrintNextStepsHelperForNewProjectWithExistingTheme pprints an help message as next steps for a project creation with an existing theme'.
func (uc *UserProjectConfig) PrintNextStepsHelperForNewProjectWithExistingTheme() {
	nextSteps := styles.List.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			styles.ListHeader("Next Steps"),
			styles.ListNumbered("1. cd "+uc.ProjectName),
			styles.ListNumbered("2. sveltin install "+styles.Gray("(or npm run install, pnpm install, ...)")),
			styles.ListNumbered("3. git init"),
			styles.ListNumbered("4. git submodule add <github_repu_url_for_the_theme> themes/<theme_name>"),
			styles.ListNumbered("5. Follow the instructions on the README file from the theme creator"),
			styles.ListNumbered("6. sveltin server"+styles.Gray(" (or npm run dev, pnpm dev, ...)")),
		))
	infoMsg := styles.List.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			styles.Desc("To stop the dev server, hit Ctrl-C"),
			styles.Info("Visit the Theme guide"+styles.Divider+styles.Url("https://docs.sveltin.io/theming")),
		))
	fmt.Println(nextSteps)
	fmt.Println(infoMsg)
}

// PrintHelperTextNewResource prints an help message string for 'resource creation'.
func PrintHelperTextNewResource(name string) {
	exampleString := fmt.Sprintf("Eg: sveltin new content %s/getting-started", name)
	infoMsg := styles.List.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			styles.ActionResult("Resource ready to be used. Start by adding content to it."),
			styles.Info("Ensure your markdown frontmatter includes it."),
			styles.NewLine,
			styles.Plain("Example:"),
			styles.NewLine,
			styles.Italic(exampleString),
		))
	fmt.Println(infoMsg)
}

// PrintHelperTextNewMetadata prints an help message string for 'metadata creation'.
func PrintHelperTextNewMetadata(metadataInfo *config.TemplateData) {
	var exampleString string
	if metadataInfo.Type == "single" {
		exampleString = fmt.Sprintf("%s%s", utils.ToSnakeCase(metadataInfo.Name), ": your_value")
	} else {
		exampleString = utils.ToSnakeCase(metadataInfo.Name) + `:
  - value 1
  - value 2`
	}

	infoMsg := styles.List.Render(lipgloss.JoinVertical(lipgloss.Left,
		styles.ActionResult("Metadata ready to be used!"),
		styles.Info("Ensure your markdown frontmatter includes it."),
		styles.NewLine,
		styles.Plain("Example:"),
		styles.NewLine,
		styles.Italic(exampleString),
	))
	fmt.Println(infoMsg)
}

// PrintHelperTextNewTheme returns an help message string for 'theme creation'.
func (uc *UserProjectConfig) PrintHelperTextNewTheme() {
	nextSteps := styles.List.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			styles.ListHeader("Next Steps"),
			styles.ListNumbered("1. cd "+uc.ProjectName),
			styles.ListNumbered("2. sveltin install "+styles.Gray("(or npm run install, pnpm install, ...)")),
			styles.ListNumbered("3. Create your theme components and partials"),
			styles.ListNumbered("4. sveltin server"+styles.Gray(" (or npm run dev, pnpm dev, ...)")),
		))
	infoMsg := styles.List.Render(lipgloss.JoinVertical(lipgloss.Left,
		styles.Desc("To stop the dev server, hit Ctrl-C"),
		styles.Info("Visit the Theme guide"+styles.Divider+styles.Url("https://docs.sveltin.io/theming")),
	))
	fmt.Println(nextSteps)
	fmt.Println(infoMsg)
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
func ShowDeployCommandWarningMessages() {
	listLogger := logger.NewListLogger()
	listLogger.Logger.Printer.SetPrinterOptions(&logger.PrinterOptions{
		Timestamp: false,
		Colors:    true,
		Labels:    true,
		Icons:     true,
	})

	listLogger.Title("Be aware! The deploy command will perform the following actions")
	listLogger.Append(logger.WarningLevel, "Create a backup of the existing content on the remote folder")
	listLogger.Append(logger.WarningLevel, "Delete existing content except what specified with --exclude flag")
	listLogger.Append(logger.WarningLevel, "Upload content to the remote folder")
	listLogger.Render()

}
