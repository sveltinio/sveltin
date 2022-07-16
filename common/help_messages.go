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

	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/pkg/logger"
)

// HelperTextNewProject returns an help message text for 'project creation'.
func HelperTextNewProject(projectName string) string {
	placeHolderText := `1. cd %s
  2. sveltin install (or npm run install, pnpm install, ...)
  3. sveltin server (or npm run dev, pnpm dev, ...)

To stop the dev server, hit Ctrl-C

Visit the Quick Start guide at https://docs.sveltin.io/quick-start
`
	return fmt.Sprintf(placeHolderText, projectName)
}

// HelperTextNewProjectWithExistingTheme returns an help message text for 'project creation when using an existing theme'.
func HelperTextNewProjectWithExistingTheme(projectName string) string {
	placeHolderText := `1. cd %s
  2. sveltin install (or npm run install, pnpm install, ...)
  3. git init
  4. git submodule add <github_repu_url_for_the_theme> themes/<theme_name>
  5. Follow the instructions on the README from the theme creator
  6. sveltin server

To stop the dev server, hit Ctrl-C

Visit the Quick Start guide at https://docs.sveltin.io/theming
`
	return fmt.Sprintf(placeHolderText, projectName)
}

// HelperTextNewMetadata returns an help message string for 'metadata creation'.
func HelperTextNewMetadata(metadataInfo *config.TemplateData) string {
	placeHolderText := `Metadata ready to be used!
   Ensure your markdown frontmatter includes it.

   Example:

   %s
`
	var exampleString string
	if metadataInfo.Type == "single" {
		exampleString = fmt.Sprintf("%s%s", toSnakeCase(metadataInfo.Name), ": your_value")
	} else {
		exampleString = toSnakeCase(metadataInfo.Name) + `:
    - value 1
    - value 2`
	}
	return fmt.Sprintf(placeHolderText, exampleString)
}

// HelperTextNewTheme returns an help message string for 'theme creation'.
func HelperTextNewTheme(projectName string) string {
	placeHolderText := `1. cd %s
  2. sveltin install (or npm run install, pnpm install, ...)
  3. Create your theme components and partials
  4. sveltin server (or npm run dev, pnpm dev, ...)

To stop the dev server, hit Ctrl-C

Visit the Quick Start guide at https://docs.sveltin.io/theming
`
	return fmt.Sprintf(placeHolderText, projectName)
}

// HelperTextDryRunFlag returns an help message string for commands supporting the 'dry-run mode'.
func HelperTextDryRunFlag() string {
	return `  * *********************** *
  * RUNNING IN DRY-RUN MODE *
  * *********************** *`
}

// HelperTextDeploySummary returns a summary message string for commands like deploy.
func HelperTextDeploySummary(numOfFolders, numOfFiles int) string {
	placeHolderText := `
SUMMARY:

- Total number of created folders: %s
- Total number of copied files: %s`

	folderCounter := strconv.Itoa(numOfFolders)
	filesCounter := strconv.Itoa(numOfFiles)

	return fmt.Sprintf(placeHolderText, folderCounter, filesCounter)
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
