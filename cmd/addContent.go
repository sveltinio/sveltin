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
	"path"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/sveltinio/prompti/choose"
	"github.com/sveltinio/prompti/input"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/internal/composer"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/internal/tpltypes"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var (
	withSampleContent bool
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
	Short:   "Add new content to an existing resource.",
	Long: resources.GetASCIIArt() + `
Command used to create a new markdown file as content and a folder to store the statics used by the content itself.

New file can contain just the frontmatter or a sample content.
Use the --template flag to select the right one to you. Valid options: blank or sample

**Note**: This command needs an existing resource created by running: sveltin new resource <resource_name>.

Example:

1. You have already created a "posts" resource
2. run: sveltin new content posts/my-first-content --sample

As result:

- a new "my-first-post" folder within "content/posts" is created
- an index.svx file is placed there
- a new "posts/my-first-port" folder created within the "static" folder to store images relative to the content
`,
	Run: RunAddContentCmd,
}

// RunAddContentCmd is the actual work function.
func RunAddContentCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	contentData, err := promptContentName(cfg.fs, args, withSampleContent, cfg.settings)
	utils.ExitIfError(err)

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
	sfs := factory.NewContentArtifact(&resources.SveltinFS, cfg.fs)
	err = projectFolder.Create(sfs)
	utils.ExitIfError(err)
	cfg.log.Success("Done\n")
}

func contentCmdFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&withSampleContent, "sample", "s", false, "Add sample content to the markdown file.")
}

func init() {
	contentCmdFlags(addContentCmd)
	addCmd.AddCommand(addContentCmd)
}

//=============================================================================

func promptContentName(fs afero.Fs, inputs []string, isSample bool, s *config.SveltinSettings) (*tpltypes.ContentData, error) {
	contentType := Blank
	if isSample {
		contentType = Sample
	}

	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		contentNamePromptContent := &input.Config{
			Placeholder: "What's the content title? (it will be the slug to the page)",
			ErrorMsg:    "Please, provide a title for the content.",
		}
		contentName, err := input.Run(contentNamePromptContent)
		if err != nil {
			return nil, err
		}

		contentResource, err := promptResourceList(fs, s)
		if err != nil {
			return nil, err
		}

		return &tpltypes.ContentData{
			Name:     utils.ToSlug(contentName),
			Type:     contentType,
			Resource: contentResource,
		}, nil
	case numOfArgs == 1:
		name := inputs[0]
		contentResource, contentName := path.Split(name)
		contentResource = strings.ReplaceAll(contentResource, "/", "")

		err := helpers.ResourceExists(fs, contentResource, cfg.settings)
		if err != nil {
			return nil, err
		}

		return &tpltypes.ContentData{
			Name:     utils.ToSlug(contentName),
			Type:     contentType,
			Resource: contentResource,
		}, nil
	default:
		err := errors.New("something went wrong: value not valid")
		return nil, sveltinerr.NewDefaultError(err)
	}
}

func promptResourceList(fs afero.Fs, s *config.SveltinSettings) (string, error) {
	availableResources := helpers.GetAllResources(fs, s.GetContentPath())

	entries := choose.ToListItem(availableResources)
	resourcePromptContent := &choose.Config{
		Title:    "Which existing resource?",
		ErrorMsg: "Please, provide an existing resource name.",
	}
	result, err := choose.Run(resourcePromptContent, entries)
	if err != nil {
		return "", err
	}
	return result, nil
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
