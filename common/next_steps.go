package common

func HelperTextNewProject(projectName string) string {
	return `
   1. cd ` + projectName + `
   2. sveltin prepare (or sveltin init, etc)")
   3. sveltin server

To close the dev server, hit Ctrl-C

Visit the Quick Start guide at https://docs.sveltin.io/quick-start
`
}

func HelperTextNewResource(resourceName string) string {
	return `Your resource is ready to be used. Start by adding content to it.

Eg: sveltin new content ` + resourceName + `/getting-started
`
}

func HelperTextNewMetadata(metadataName string) string {
	return `Your metadata ` + metadataName + ` is ready to be used.

Ensure your markdown frontmatter consider it.
`
}
