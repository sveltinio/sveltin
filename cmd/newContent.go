/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package cmd ...
package cmd

import (
	"embed"
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/sveltinlib/composer"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var (
	templateFlag string
)

const (
	// BLANK represents the fontmatter-only template id used when generating the content file.
	BLANK string = "blank"
	// SAMPLE represents the sample-content template id used when generating the content file.
	SAMPLE string = "sample"
)

//=============================================================================

var newContentCmd = &cobra.Command{
	Use:     "content [name]",
	Aliases: []string{"c"},
	Short:   "Create a new content for existing resource",
	Long: resources.GetAsciiArt() + `
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
	contentData, err := getContentName(AppFs, args, &conf)
	utils.ExitIfError(err)

	log.Info(fmt.Sprintf("'%s' content will be added", contentData.Name))

	// GET FOLDER: content
	contentFolder := fsManager.GetFolder(CONTENT)

	// NEW FOLDER content/<resource_name>/<content_name>
	resourceContentFolder := fsManager.NewResourceContentFolder(contentData.Name, contentData.Resource)

	// NEW FILE: content/<resource_name>/<content_name>/index.svx
	contentFile := &composer.File{
		Name:       pathMaker.GetResourceContentFilename(),
		TemplateId: contentData.Type,
		TemplateData: &config.TemplateData{
			Name: contentData.Name,
		},
	}

	resourceContentFolder.Add(contentFile)
	contentFolder.Add(resourceContentFolder)

	// SET STATIC FOLDER STRUCTURE for resouce and content
	staticFolder := makeStaticFolderStructure(&resources.SveltinFS, AppFs, contentData)

	// SET FOLDER STRUCTURE
	projectFolder := fsManager.GetFolder(ROOT)
	projectFolder.Add(staticFolder)
	projectFolder.Add(contentFolder)

	// GENERATE FOLDER STRUCTURE
	sfs := factory.NewContentArtifact(&resources.SveltinFS, AppFs)
	err = projectFolder.Create(sfs)
	utils.ExitIfError(err)
	log.Success("Done")
}

func contentCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&templateFlag, "template", "t", "", "Generate a markdown file based on templates (valid options: blank or sample).")
}

func init() {
	newCmd.AddCommand(newContentCmd)
	contentCmdFlags(newContentCmd)
}

//=============================================================================

func getContentName(fs afero.Fs, inputs []string, c *config.SveltinConfig) (config.TemplateData, error) {
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		contentNamePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide a title for the content.",
			Label:    "What's the content title? (it will be the slug to the page)",
		}
		contentName := common.PromptGetInput(contentNamePromptContent)
		contentType, err := promptContentTemplateSelection(templateFlag)
		utils.ExitIfError(err)

		contentResource, err := promptResourceList(fs, c)
		utils.ExitIfError(err)

		return config.TemplateData{
			Name:     utils.ToSlug(contentName),
			Type:     contentType,
			Resource: contentResource,
		}, nil
	case numOfArgs == 1:
		name := inputs[0]

		contentResource, contentName := path.Split(name)
		contentResource = strings.ReplaceAll(contentResource, "/", "")
		err := helpers.ResourceExists(fs, contentResource, &conf)
		utils.ExitIfError(err)

		contentType, err := promptContentTemplateSelection(templateFlag)
		utils.ExitIfError(err)

		return config.TemplateData{
			Name:     utils.ToSlug(contentName),
			Type:     contentType,
			Resource: contentResource,
		}, nil
	default:
		err := errors.New("something went wrong: value not valid")
		return config.TemplateData{}, sveltinerr.NewDefaultError(err)
	}
}

func promptContentTemplateSelection(templateType string) (string, error) {
	validTemplates := []string{BLANK, SAMPLE}
	var contentTemplate string

	switch nameLenght := len(templateType); {
	case nameLenght == 0:
		templatePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide select a template for your content file.",
			Label:    "Do you prefer a blank file or with some sample content in place?",
		}
		contentTemplate = common.PromptGetSelect(validTemplates, templatePromptContent)
		return contentTemplate, nil
	case nameLenght != 0:
		contentTemplate = templateType
		if !common.Contains(validTemplates, contentTemplate) {
			return "", sveltinerr.NewContentTemplateTypeNotValidError()
		}
		return contentTemplate, nil
	default:
		return "", sveltinerr.NewContentTemplateTypeNotValidError()
	}
}

func promptResourceList(fs afero.Fs, c *config.SveltinConfig) (string, error) {
	availableResources := helpers.GetAllResources(fs, c.GetContentPath())

	resourcePromptContent := config.PromptContent{
		ErrorMsg: "Please, provide select an existing resource.",
		Label:    "What's the existing resource to be used?",
	}
	resource := common.PromptGetSelect(availableResources, resourcePromptContent)
	return resource, nil
}

//=============================================================================

func makeStaticFolderStructure(efs *embed.FS, fs afero.Fs, contentData config.TemplateData) *composer.Folder {
	// GET FOLDER: static
	staticFolder := fsManager.GetFolder(STATIC)
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
