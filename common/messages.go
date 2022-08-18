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
	entries := map[string]string{
		"Project Name: ": uc.ProjectName,
		"CSSLib: ":       uc.CSSLibName,
		"Theme: ":        uc.ThemeName,
		"NPM Client: ":   uc.NPMClientName,
	}
	fmt.Println(styles.NewULWithIconPrefix("RECAP your choices", entries, styles.CheckMark))
}

// PrintNextStepsHelperForNewProject prints an help message as next steps for a project creation.
func (uc *UserProjectConfig) PrintNextStepsHelperForNewProject() {
	steps := []string{
		fmt.Sprintf("cd %s", uc.ProjectName),
		fmt.Sprintf("sveltin install %s", styles.Faint("(or npm run install, pnpm install, ...)")),
		fmt.Sprintf("sveltin server %s", styles.Faint(" (or npm run dev, pnpm dev, ...)")),
	}
	nextStepsMsg := styles.NewOLWithTitle("Next Steps", steps)

	fmt.Printf("%s%s", nextStepsMsg, devServerInfoMessage())
}

// PrintNextStepsHelperForNewProjectWithExistingTheme pprints an help message as next steps for a project creation with an existing theme'.
func (uc *UserProjectConfig) PrintNextStepsHelperForNewProjectWithExistingTheme() {
	steps := []string{
		fmt.Sprintf("cd %s", uc.ProjectName),
		fmt.Sprintf("sveltin install %s", styles.Faint("(or npm run install, pnpm install, ...)")),
		"git init",
		"git submodule add <github_repu_url_for_the_theme> themes/<theme_name>",
		"Follow the instructions on the README file from the theme creator",
		fmt.Sprintf("sveltin server %s", styles.Faint(" (or npm run dev, pnpm dev, ...)")),
	}
	nextStepsMsg := styles.NewOLWithTitle("Next Steps", steps)

	fmt.Printf("%s%s\n", nextStepsMsg, devServerInfoMessage())
}

// PrintHelperTextNewResource prints an help message string for 'resource creation'.
func PrintHelperTextNewResource(name string) {
	exampleString := fmt.Sprintf("sveltin new content %s/getting-started", name)
	entries := []string{
		styles.P("Start by adding content to it, e.g."),
		styles.BR,
		styles.Code(exampleString),
	}

	fmt.Println(styles.Section("Resource ready to be used.", entries))
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

	entries := []string{
		styles.P("Ensure your markdown frontmatter includes it, e.g."),
		styles.BR,
		styles.Code(exampleString),
	}

	fmt.Println(styles.Section("Metadata ready to be used.", entries))
}

// PrintHelperTextNewTheme returns an help message string for 'theme creation'.
func (uc *UserProjectConfig) PrintHelperTextNewTheme() {
	steps := []string{
		fmt.Sprintf("cd %s", uc.ProjectName),
		fmt.Sprintf("sveltin install %s", styles.Faint("(or npm run install, pnpm install, ...)")),
		"Create your theme components and partials",
		fmt.Sprintf("sveltin server %s", styles.Faint(" (or npm run dev, pnpm dev, ...)")),
	}
	nextStepsMsg := styles.NewOLWithTitle("Next Steps", steps)

	fmt.Printf("%s\n%s\n", nextStepsMsg, themingInfoMessage())
}

// PrintHelperTextDryRunFlag prints a message box for commands supporting the 'dry-run mode'.
func PrintHelperTextDryRunFlag() {
	fmt.Println(styles.Bordered("RUNNING IN DRY-RUN MODE"))
}

// PrintHelperTextDeploySummary prints a summary message string for commands like deploy.
func PrintHelperTextDeploySummary(numOfFolders, numOfFiles int) {
	entries := map[string]string{
		"Total number of created folders: ": strconv.Itoa(numOfFolders),
		"Total number of copied files: ":    strconv.Itoa(numOfFiles),
	}
	fmt.Println(styles.NewULWithIconPrefix("SUMMARY", entries, styles.CheckMark))
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

func devServerInfoMessage() string {
	return styles.Section("To stop the dev server, hit Ctrl-C",
		[]string{
			styles.Faint("Visit the Quick Start guide") +
				styles.Divider +
				styles.A("https://docs.sveltin.io/quick-start")})
}

func themingInfoMessage() string {
	return styles.Section("To stop the dev server, hit Ctrl-C",
		[]string{
			styles.Faint("Visit the Theme guide") +
				styles.Divider +
				styles.A("https://docs.sveltin.io/theming"),
		})
}
