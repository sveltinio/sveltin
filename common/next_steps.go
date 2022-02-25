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
)

// HelperTextNewProject returns a 'next step' text used after the project creation.
func HelperTextNewProject(projectName string) string {
	placeHolderText := `
   1. cd %s
   2. sveltin prepare (or sveltin init, etc)
   3. sveltin server

To close the dev server, hit Ctrl-C

Visit the Quick Start guide at https://docs.sveltin.io/quick-start
`
	return fmt.Sprintf(placeHolderText, projectName)
}

// HelperTextNewResource returns a 'next step' text used after the resource creation.
func HelperTextNewResource(resourceName string) string {
	placeHolderText := `Your resource is ready to be used. Start by adding content to it.

Eg: sveltin new content %s/getting-started
`
	return fmt.Sprintf(placeHolderText, resourceName)
}

// HelperTextNewMetadata returns a 'next step' text used after the resource creation.
func HelperTextNewMetadata(metadataName string) string {
	placeHolderText := `Your metadata %s is ready to be used.

Ensure your markdown frontmatter consider it.
`
	return fmt.Sprintf(placeHolderText, metadataName)
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
