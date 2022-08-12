/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package cmd ...
package cmd

import (
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/internal/composer"
	"github.com/sveltinio/sveltin/internal/tui/choose"
	"github.com/sveltinio/sveltin/internal/tui/input"
	"github.com/sveltinio/sveltin/pkg/sveltinerr"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var (
	templateFlag string
)

const (
	// Blank represents the fontmatter-only template id used when generating the content file.
	Blank string = "blank"
	// Sample represents the sample-content template id used when generating the content file.
	Sample string = "sample"
)

//=============================================================================

var newContentCmd = &cobra.Command{
	Use:     "content [name]",
	Aliases: []string{"c"},
	Short:   "Command to create a new content for existing resource",
	Long: resources.GetASCIIArt() + `
Create a new markdown file for your content and a folder to store the statics used by the content itself.

New file can contain just the frontmatter or a sample content.
Use the --template flag to select the right one to you. Valid options: blank or sample

**Note**: This command must be used after you create a resource for the content.

Example:

1. You have already created a new resource called "posts"
2. call "sveltin new content posts/my-first-content --template sample"

As result:

- a new "my-first-post" folder within "content/posts" is created
- an index.svx file is placed there
- a new "posts/my-first-port" folder created within the "static" folder to store images relative to the content
`,
	Run: RunNewContentCmd,
}

// RunNewContentCmd is the actual work function.
func RunNewContentCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	contentData, err := promptContentName(cfg.fs, args, cfg.sveltin)
	utils.ExitIfError(err)

	cfg.log.Plain(utils.Underline(fmt.Sprintf("Creating '%s' as content for the %s resource", contentData.Name, contentData.Resource)))

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
	cfg.log.Success("Done")
}

func contentCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&templateFlag, "template", "t", "", "Generate a markdown file based on templates (valid options: blank or sample).")
}

func init() {
	newCmd.AddCommand(newContentCmd)
	contentCmdFlags(newContentCmd)
}

//=============================================================================

func promptContentName(fs afero.Fs, inputs []string, c *config.SveltinConfig) (*config.TemplateData, error) {
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		contentNamePromptContent := &input.Config{
			Placeholder: "What's the content title? (it will be the slug to the page)",
			ErrorMsg:    "Please, provide a title for the content.",
		}
		//contentName, err := common.PromptGetInput(contentNamePromptContent, nil, "")
		contentName, err := input.Run(contentNamePromptContent)
		if err != nil {
			return nil, err
		}
		contentType, err := promptContentTemplateSelection(templateFlag)
		if err != nil {
			return nil, err
		}

		contentResource, err := promptResourceList(fs, c)
		if err != nil {
			return nil, err
		}

		return &config.TemplateData{
			Name:     utils.ToSlug(contentName),
			Type:     contentType,
			Resource: contentResource,
		}, nil
	case numOfArgs == 1:
		name := inputs[0]
		contentResource, contentName := path.Split(name)
		contentResource = strings.ReplaceAll(contentResource, "/", "")

		err := helpers.ResourceExists(fs, contentResource, cfg.sveltin)
		if err != nil {
			return nil, err
		}

		contentType, err := promptContentTemplateSelection(templateFlag)
		if err != nil {
			return nil, err
		}

		return &config.TemplateData{
			Name:     utils.ToSlug(contentName),
			Type:     contentType,
			Resource: contentResource,
		}, nil
	default:
		err := errors.New("something went wrong: value not valid")
		return nil, sveltinerr.NewDefaultError(err)
	}
}

func promptContentTemplateSelection(templateType string) (string, error) {
	entries := []list.Item{
		choose.Item{Name: Blank, Desc: "Frontmatter only"},
		choose.Item{Name: Sample, Desc: "Full sample content"},
	}

	switch nameLenght := len(templateType); {
	case nameLenght == 0:
		templatePromptContent := &choose.Config{
			Title:    "Which template for your content?",
			ErrorMsg: "Please, provide a template name for the content file.",
		}
		result, err := choose.Run(templatePromptContent, entries)
		if err != nil {
			return "", err
		}
		return result, nil
	case nameLenght != 0:
		valid := choose.GetItemsKeys(entries)
		if !common.Contains(valid, templateType) {
			return "", sveltinerr.NewContentTemplateTypeNotValidError()
		}
		return templateType, nil
	default:
		return "", sveltinerr.NewContentTemplateTypeNotValidError()
	}
}

func promptResourceList(fs afero.Fs, c *config.SveltinConfig) (string, error) {
	availableResources := helpers.GetAllResources(fs, c.GetContentPath())

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

func makeContentFolderStructure(folderName string, contentData *config.TemplateData) (*composer.Folder, error) {
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

func createContentLocalFolder(contentData *config.TemplateData) *composer.Folder {
	// GET FOLDER: content
	contentFolder := cfg.fsManager.GetFolder(ContentFolder)

	// NEW FOLDER content/<resource_name>/<content_name>
	resourceContentFolder := cfg.fsManager.NewResourceContentFolder(contentData.Name, contentData.Resource)

	// NEW FILE: content/<resource_name>/<content_name>/index.svx
	contentFile := &composer.File{
		Name:       cfg.pathMaker.GetResourceContentFilename(),
		TemplateID: contentData.Type,
		TemplateData: &config.TemplateData{
			Name: contentData.Name,
		},
	}

	resourceContentFolder.Add(contentFile)
	contentFolder.Add(resourceContentFolder)
	return contentFolder
}

func createStaticFolderStructure(contentData *config.TemplateData) *composer.Folder {
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
