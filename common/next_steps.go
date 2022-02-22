package common

import (
	"fmt"
	"strconv"
)

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

func HelperTextNewResource(resourceName string) string {
	placeHolderText := `Your resource is ready to be used. Start by adding content to it.

Eg: sveltin new content %s /getting-started
`
	return fmt.Sprintf(placeHolderText, resourceName)
}

func HelperTextNewMetadata(metadataName string) string {
	placeHolderText := `Your metadata %s is ready to be used.

Ensure your markdown frontmatter consider it.
`
	return fmt.Sprintf(placeHolderText, metadataName)
}

func HelperTextDryRunFlag() string {
	return `  * *********************** *
  * RUNNING IN DRY-RUN MODE *
  * *********************** *`
}

func HelperTextDeploySummary(numOfFolders, numOfFiles int) string {
	placeHolderText := `
SUMMARY:

- Total number of created folders: %s
- Total number of copied files: %s`

	folderCounter := strconv.Itoa(numOfFolders)
	filesCounter := strconv.Itoa(numOfFiles)

	return fmt.Sprintf(placeHolderText, folderCounter, filesCounter)
}
