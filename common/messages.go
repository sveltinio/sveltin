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
	"github.com/sveltinio/sveltin/pkg/logger"
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
		list.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listHeader("Here are your choices"),
				listDone("Project Name: ", uc.ProjectName),
				listDone("CSS Lib: ", uc.CSSLibName),
				listDone("Theme: ", uc.ThemeName),
				listDone("NPM Client: ", uc.NPMClientName),
			),
		),
	)
}

// PrintNextStepsHelperForNewProject prints an help message as next steps for a project creation.
func (uc *UserProjectConfig) PrintNextStepsHelperForNewProject() {
	desc := list.Render(lipgloss.JoinVertical(lipgloss.Left,
		listHeader("Next Steps"),
		listNumbered("1. cd "+uc.ProjectName),
		listNumbered("2. sveltin install "+gray("(or npm run install, pnpm install, ...)")),
		listNumbered("3. sveltin server"+gray(" (or npm run dev, pnpm dev, ...)")),
	))
	additionalInfo := list.Render(lipgloss.JoinVertical(lipgloss.Left,
		descStyle.Render("To stop the dev server, hit Ctrl-C"),
		infoStyle.Render("Visit the Quick Start guide"+divider+url("https://docs.sveltin.io/quick-start")),
	))
	fmt.Println(desc)
	fmt.Println(additionalInfo)
}

// PrintNextStepsHelperForNewProjectWithExistingTheme pprints an help message as next steps for a project creation with an existing theme'.
func (uc *UserProjectConfig) PrintNextStepsHelperForNewProjectWithExistingTheme() {
	desc := list.Render(lipgloss.JoinVertical(lipgloss.Left,
		listHeader("Next Steps"),
		listNumbered("1. cd "+uc.ProjectName),
		listNumbered("2. sveltin install "+gray("(or npm run install, pnpm install, ...)")),
		listNumbered("3. git init"),
		listNumbered("4. git submodule add <github_repu_url_for_the_theme> themes/<theme_name>"),
		listNumbered("5. Follow the instructions on the README file from the theme creator"),
		listNumbered("6. sveltin server"+gray(" (or npm run dev, pnpm dev, ...)")),
	))
	additionalInfo := list.Render(lipgloss.JoinVertical(lipgloss.Left,
		descStyle.Render("To stop the dev server, hit Ctrl-C"),
		infoStyle.Render("Visit the Theme guide"+divider+url("https://docs.sveltin.io/theming")),
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

	infoMsg := list.Render(lipgloss.JoinVertical(lipgloss.Left,
		titleStyle.Render("Metadata ready to be used!"),
		infoStyle.Render("Ensure your markdown frontmatter includes it."),
		newLine,
		plain("Example:"),
		newLine,
		italic(exampleString),
	))

	fmt.Println(infoMsg)

}

// PrintHelperTextNewTheme returns an help message string for 'theme creation'.
func (uc *UserProjectConfig) PrintHelperTextNewTheme() {
	desc := list.Render(lipgloss.JoinVertical(lipgloss.Left,
		listHeader("Next Steps"),
		listNumbered("1. cd "+uc.ProjectName),
		listNumbered("2. sveltin install "+gray("(or npm run install, pnpm install, ...)")),
		listNumbered("3. Create your theme components and partials"),
		listNumbered("4. sveltin server"+gray(" (or npm run dev, pnpm dev, ...)")),
	))
	additionalInfo := list.Render(lipgloss.JoinVertical(lipgloss.Left,
		descStyle.Render("To stop the dev server, hit Ctrl-C"),
		infoStyle.Render("Visit the Theme guide"+divider+url("https://docs.sveltin.io/theming")),
	))
	fmt.Println(desc)
	fmt.Println(additionalInfo)
}

// PrintHelperTextDryRunFlag prints a message box for commands supporting the 'dry-run mode'.
func PrintHelperTextDryRunFlag() {
	fmt.Println(bordered("RUNNING IN DRY-RUN MODE"))
}

// PrintHelperTextDeploySummary prints a summary message string for commands like deploy.
func PrintHelperTextDeploySummary(numOfFolders, numOfFiles int) {
	folderCounter := strconv.Itoa(numOfFolders)
	filesCounter := strconv.Itoa(numOfFiles)
	fmt.Println(
		list.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listHeader("SUMMARY"),
				listDone("Total number of created folders: ", folderCounter),
				listDone("Total number of copied files: ", filesCounter),
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

func toSnakeCase(txt string) string {
	cleanString := strings.ToLower(txt)
	cleanString = strings.ReplaceAll(cleanString, " ", "_")
	cleanString = strings.ReplaceAll(cleanString, "-", "_")
	return cleanString
}
