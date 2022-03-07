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

	"github.com/sveltinio/sveltin/sveltinlib/logger"
)

// HelperTextNewProject returns a 'next step' text used after the project creation.
func HelperTextNewProject(projectName string) string {
	placeHolderText := `1. cd %s
  2. sveltin prepare (or sveltin init, etc)
  3. sveltin server

To stop the dev server, hit Ctrl-C

Visit the Quick Start guide at https://docs.sveltin.io/quick-start`
	return fmt.Sprintf(placeHolderText, projectName)
}

// HelperTextDryRunFlag returns a text used when running an action in dry-run mode.
func HelperTextDryRunFlag() string {
	return `  * *********************** *
  * RUNNING IN DRY-RUN MODE *
  * *********************** *`
}

// HelperTextDeploySummary returns a 'summary' text for commands like deploy.
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
