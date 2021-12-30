/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package cmd

import (
	"errors"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/sveltinlib/composer"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var (
	resourceName string
	metadataType string
)

//=============================================================================

var newMetadataCmd = &cobra.Command{
	Use:   "metadata [name] --resource [resource] --type [single|list]",
	Short: "Command to add new metadata to an existing resource",
	Long: resources.GetAsciiArt() + `
Command to add new metadata from your content to an existing resource.

What is a "metadata" for Sveltin?
Whatever you enter in the front-matter of your markdown content for which you want content grouped by it.

- 1:1 relationship is the use for metadata type "single" (E.g. posts by category)
- 1:many relationship is the use for metadata type "list" (E.g. posts by tags)
	`,
	Run: RunNewMetadataCmd,
}

func RunNewMetadataCmd(cmd *cobra.Command, args []string) {
	logger.Reset()

	printer := utils.PrinterContent{
		Title: "New metadata folder structure is going to be created:",
	}

	mdName, err := promptMetadataName(args)
	common.CheckIfError(err)

	mdResource, err := promptResource(AppFs, resourceName, &conf)
	common.CheckIfError(err)

	mdType, err := promptMetadataType(metadataType)
	common.CheckIfError(err)

	// NEW FOLDER: <metadata_name>
	resourceMetadataAPIFolder := composer.NewFolder(mdName)

	// NEW FILE: groupedBy.json.ts
	logger.AppendItem("Creating an API endpoint for the metadata")
	resourceMetadataAPIFile := &composer.File{
		Name:       conf.GetMetadataAPIFilename(),
		TemplateId: API,
		TemplateData: &config.TemplateData{
			Name:     mdName,
			Resource: mdResource,
			Type:     mdType,
			Config:   &conf,
		},
	}
	resourceMetadataAPIFolder.Add(resourceMetadataAPIFile)

	// NEW FOLDER: <metadata_name>
	resourceAPIFolder := composer.NewFolder(mdResource)
	resourceAPIFolder.Add(resourceMetadataAPIFolder)

	// GET FOLDER: src/routes/api/<api_version> folder
	apiFolder := fsManager.GetFolder(API)
	apiFolder.Add(resourceAPIFolder)

	// NEW FILE: <metadata_name>.js file into src/lib folder
	logger.AppendItem("Creating a lib file for the metadata")
	libFile := &composer.File{
		Name:       pathMaker.GetResourceLibFilename(mdName),
		TemplateId: LIB,
		TemplateData: &config.TemplateData{
			Name:     mdName,
			Resource: mdResource,
			Type:     mdType,
			Config:   &conf,
		},
	}
	// NEW FOLDER: src/lib/<resource_name>
	resourceLibFolder := composer.NewFolder(mdResource)
	resourceLibFolder.Add(libFile)

	// GET FOLDER: src/lib folder
	libFolder := fsManager.GetFolder(LIB)
	libFolder.Add(resourceLibFolder)

	// NEW FOLDER: <metadata_name>
	resourceMedatadaRoutesFolder := composer.NewFolder(mdName)

	// NEW FILE: <resource_name>/<metadata_name>/index.svelte
	logger.AppendItem("Creating an index.svelte component for the metadata")
	for _, item := range []string{INDEX, SLUG} {
		f := &composer.File{
			Name:       helpers.GetResourceRouteFilename(item, &conf),
			TemplateId: item,
			TemplateData: &config.TemplateData{
				Name:     mdName,
				Resource: mdResource,
				Type:     mdType,
				Config:   &conf,
			},
		}
		resourceMedatadaRoutesFolder.Add(f)
	}

	// NEW FOLDER: src/routes/<resource_name>/<metadata_name>
	resourceRoutesFolder := composer.NewFolder(mdResource)
	resourceRoutesFolder.Add(resourceMedatadaRoutesFolder)

	// GET FOLDER: src/routes folder
	routesFolder := fsManager.GetFolder(ROUTES)
	routesFolder.Add(resourceRoutesFolder)

	// SET FOLDER STRUCTURE
	projectFolder := fsManager.GetFolder(ROOT)
	projectFolder.Add(apiFolder)
	projectFolder.Add(libFolder)
	projectFolder.Add(routesFolder)

	// GENERATE FOLDER STRUCTURE
	sfs := factory.NewMetadataArtifact(&resources.SveltinFS, AppFs)
	err = projectFolder.Create(sfs)
	common.CheckIfError(err)

	// LOG TO STDOUT
	// LOG TO STDOUT
	printer.SetContent(logger.Render())
	utils.PrettyPrinter(&printer).Print()
	jww.FEEDBACK.Println("Your metadata is ready to be used. Ensure your markdown frontmatter consider it.")

}

func metadataCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&resourceName, "resource", "r", "", "Name of the resource the new metadata is belongs to.")
	cmd.Flags().StringVarP(&metadataType, "type", "t", "", "Type of the new metadata. (possible values: single or list)")
}

func init() {
	newCmd.AddCommand(newMetadataCmd)
	metadataCmdFlags(newMetadataCmd)
}

//=============================================================================

func promptResource(fs afero.Fs, mdResourceFlag string, c *config.SveltinConfig) (string, error) {
	var resource string

	availableResources := helpers.GetAllResources(fs, c.GetContentPath())

	switch nameLenght := len(mdResourceFlag); {
	case nameLenght == 0:
		resourcePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide select an existing resource.",
			Label:    "What's the existing resource to be used?",
		}
		resource = common.PromptGetSelect(availableResources, resourcePromptContent)
		return resource, nil
	case nameLenght != 0:
		resource = mdResourceFlag
		if !common.Contains(availableResources, resource) {
			return "", common.NewResourceNotFoundError()
		} else {
			return resource, nil
		}
	default:
		return "", common.NewResourceNotFoundError()
	}
}

func promptMetadataName(inputs []string) (string, error) {
	var name string
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		metadataNamePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide a name for the metadata.",
			Label:    "What's the name of the metadata to be added?",
		}
		name = common.PromptGetInput(metadataNamePromptContent)
		return name, nil
	case numOfArgs == 1:
		name = inputs[0]
		return name, nil
	default:
		err := errors.New("something went wrong: name not valid")
		return "", common.NewDefaultError(err)
	}

}

func promptMetadataType(mdTypeFlag string) (string, error) {
	var metadataType string
	valid := []string{"single", "list"}

	switch nameLenght := len(mdTypeFlag); {
	case nameLenght == 0:
		metadataTypePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide select a metadata type.",
			Label:    "What's the metadata type?",
		}
		metadataType = common.PromptGetSelect(valid, metadataTypePromptContent)
		return metadataType, nil
	case nameLenght != 0:
		metadataType = mdTypeFlag
		if !common.Contains(valid, metadataType) {
			return "", common.NewMetadataTypeNotValidError()
		} else {
			return metadataType, nil
		}
	default:
		return "", common.NewMetadataTypeNotValidError()
	}
}
