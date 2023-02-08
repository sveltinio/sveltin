/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package cmd

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/internal/composer"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/internal/tpltypes"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/tui/activehelps"
	"github.com/sveltinio/sveltin/tui/prompts"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var (
	resourceNameForContent string
	withSampleContent      bool
)

const (
	// Blank represents the fontmatter-only template id used when generating the content file.
	Blank string = "blank"
	// Sample represents the sample-content template id used when generating the content file.
	Sample string = "sample"
)

//=============================================================================

var addContentCmd = &cobra.Command{
	Use:     "content [name]",
	Aliases: []string{"c"},
	Short:   "Add new content to an existing resource",
	Long: resources.GetASCIIArt() + `
Command used to create a new markdown file as content and a folder to store the statics used by the content itself.

New file can contain just the frontmatter or a sample content.
Use the --template flag to select the right one to you. Valid options: blank or sample

**Note**: This command needs an existing resource created by running: sveltin new resource <resource_name>.

Example:

1. You have already created some resource by running "sveltin new resource"
2. run: sveltin add content welcome --to posts

As result:

- a new "welcome" folder within "content/posts" is created
- an index.svx file is placed there
- a new "posts/welcome" folder created within the "static" folder to store images relative to the content
`,
	Run: RunAddContentCmd,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var comps []string
		if len(args) == 0 {
			comps = cobra.AppendActiveHelp(comps, activehelps.Hint("You must choose a name for the content"))
		} else {
			comps = cobra.AppendActiveHelp(comps, activehelps.Hint("[WARN] This command does not take any more arguments but accepts flags"))
		}
		return comps, cobra.ShellCompDirectiveDefault
	},
}

// RunAddContentCmd is the actual work function.
func RunAddContentCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands either from a not valid directory or not latest sveltin version.
	isValidProject(true)

	contentName, err := prompts.AskContentNameHandler(args)
	utils.ExitIfError(err)

	contentResource, err := prompts.SelectResourceHandler(cfg.fs, resourceNameForContent, cfg.settings)
	utils.ExitIfError(err)

	contentData := tpltypes.NewContentData(contentName, contentResource, withSampleContent)

	headingText := fmt.Sprintf("Adding '%s' as content to the '%s' resource", contentData.Name, contentData.Resource)
	cfg.log.Plain(markup.H1(headingText))

	// MAKE FOLDER STRUCTURE: static/resources/<resource_name>/<content_name>
	contentFolder, err := makeContentFolderStructure(ContentFolder, contentData)
	utils.ExitIfError(err)

	// MAKE FOLDER STRUCTURE: static/resources/<resource_name>/<content_name>
	staticFolder, err := makeContentFolderStructure(StaticFolder, contentData)
	utils.ExitIfError(err)

	// SET FOLDER STRUCTURE
	projectFolder := cfg.fsManager.GetFolder(RootFolder)
	projectFolder.Add(contentFolder)
	projectFolder.Add(staticFolder)

	// GENERATE THE FOLDER TREE
	sfs := factory.NewContentArtifact(&resources.SveltinTemplatesFS, cfg.fs)
	err = projectFolder.Create(sfs)
	utils.ExitIfError(err)

	if withSampleContent {
		err := addSampleCoverImage(contentData)
		utils.ExitIfError(err)
	}
	cfg.log.Success("Done\n")
}

func contentCmdFlags(cmd *cobra.Command) {
	// to flag
	cmd.Flags().StringVarP(&resourceNameForContent, "to", "t", "", "Name of the resource the new content is belongs to")
	err := cmd.RegisterFlagCompletionFunc("to", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		availableResources := helpers.GetAllResources(cfg.fs, cfg.settings.GetContentPath())
		return availableResources, cobra.ShellCompDirectiveDefault
	})
	utils.ExitIfError(err)
	// sample flag
	cmd.Flags().BoolVarP(&withSampleContent, "sample", "s", false, "Add sample content to the markdown file")
}

func init() {
	contentCmdFlags(addContentCmd)
	addCmd.AddCommand(addContentCmd)
}

//=============================================================================

func makeContentFolderStructure(folderName string, contentData *tpltypes.ContentData) (*composer.Folder, error) {
	switch folderName {
	case ContentFolder:
		return createContentLocalFolder(contentData), nil
	case StaticFolder:
		return createStaticFolderStructure(contentData), nil
	default:
		err := errors.New("something went wrong: folder not found as mapped resource for sveltin projects")
		return nil, sveltinerr.NewDefaultError(err)
	}
}

//=============================================================================

func createContentLocalFolder(contentData *tpltypes.ContentData) *composer.Folder {
	// GET FOLDER: content
	contentFolder := cfg.fsManager.GetFolder(ContentFolder)
	// NEW FOLDER content/<resource_name>/<content_name>
	resourceContentFolder := cfg.fsManager.NewResourceContentFolder(contentData)
	// NEW FILE: content/<resource_name>/<content_name>/index.svx
	contentFile := cfg.fsManager.NewResourceContentFile(contentData)
	// SET FOLDER STRUCTURE
	resourceContentFolder.Add(contentFile)
	contentFolder.Add(resourceContentFolder)

	return contentFolder
}

func createStaticFolderStructure(contentData *tpltypes.ContentData) *composer.Folder {
	// GET FOLDER: static
	staticFolder := cfg.fsManager.GetFolder(StaticFolder)
	// NEW FOLDER static/resources
	imagesFolder := composer.NewFolder("resources")
	// NEW FOLDER static/resources/<resource_name>
	resourceImagesFolder := composer.NewFolder(contentData.Resource)
	// NEW FOLDER static/resources/<resource_name>/<content_name>
	contentImagesFolder := composer.NewFolder(contentData.Name)
	// SET FOLDER STRUCTURE
	resourceImagesFolder.Add(contentImagesFolder)
	imagesFolder.Add(resourceImagesFolder)
	staticFolder.Add(imagesFolder)

	return staticFolder
}

func addSampleCoverImage(contentData *tpltypes.ContentData) error {
	saveTo := cfg.fsManager.GetFolder(filepath.Join(StaticFolder, "resources", contentData.Resource, contentData.Name)).Name
	return cfg.fsManager.CopyFileFromEmbed(&resources.SveltinStaticFS, cfg.fs, resources.SveltinImagesFS, DummyImgFileId, saveTo)
}
