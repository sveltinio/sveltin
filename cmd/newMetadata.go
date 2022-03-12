/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package cmd ...
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/sveltinlib/composer"
	"github.com/sveltinio/sveltin/sveltinlib/logger"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var (
	resourceName string
	metadataType string
)

//=============================================================================

var newMetadataCmd = &cobra.Command{
	Use:     "metadata [name] --resource [resource] --type [single|list]",
	Aliases: []string{"m, groupedBy"},
	Short:   "Command to add new metadata to an existing resource",
	Long: resources.GetAsciiArt() + `
Command to add new metadata from your content to an existing resource.

What is a "metadata" for Sveltin?
Whatever you enter in the front-matter of your markdown content for which you want content grouped by it.

Types:

- single: 1:1 relationship (e.g. category)
- list: 1:many relationship (e.g. tags)
`,
	Run: RunNewMetadataCmd,
}

// RunNewMetadataCmd is the actual work function.
func RunNewMetadataCmd(cmd *cobra.Command, args []string) {
	listLogger := log.WithList()

	mdName, err := promptMetadataName(args)
	utils.ExitIfError(err)

	mdResource, err := promptResource(AppFs, resourceName, &conf)
	utils.ExitIfError(err)

	mdType, err := promptMetadataType(metadataType)
	utils.ExitIfError(err)

	// NEW FILE: <metadata_name>.js file into src/lib folder
	listLogger.Append(logger.LevelInfo, "Creating the lib file for the metadata")
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

	// NEW FILE: src/routes/<resource_name>/<metadata_name>/{index.svelte, index.ts, [slug].svelte, [slug].ts}
	listLogger.Append(logger.LevelInfo, "Creating the components and endpoints for the metadata")
	for _, item := range []string{INDEX, INDEX_ENDPOINT, SLUG, SLUG_ENDPOINT} {
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
	projectFolder.Add(libFolder)
	projectFolder.Add(routesFolder)

	// GENERATE FOLDER STRUCTURE
	sfs := factory.NewMetadataArtifact(&resources.SveltinFS, AppFs)
	err = projectFolder.Create(sfs)
	utils.ExitIfError(err)

	// LOG TO STDOUT
	listLogger.Info(fmt.Sprintf("New '%s' as metadata will be created", mdName))
	log.Success("Done")

	// NEXT STEPS
	listLogger = log.WithList()
	listLogger.Append(logger.LevelSuccess, "Metadata ready to be used.")
	listLogger.Append(logger.LevelImportant, "Ensure your markdown frontmatter consider it.")
	listLogger.Info("Next Steps")
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
		return utils.ToSlug(resource), nil
	case nameLenght != 0:
		resource = mdResourceFlag
		if !common.Contains(availableResources, resource) {
			return "", sveltinerr.NewResourceNotFoundError()
		}
		return utils.ToSlug(resource), nil
	default:
		return "", sveltinerr.NewResourceNotFoundError()
	}
}

func promptMetadataName(inputs []string) (string, error) {
	var name string
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		metadataNamePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide a name for the metadata.",
			Label:    "What's the metadata name?",
		}
		name = common.PromptGetInput(metadataNamePromptContent)
		return utils.ToSlug(name), nil
	case numOfArgs == 1:
		name = inputs[0]
		return utils.ToSlug(name), nil
	default:
		err := errors.New("something went wrong: name not valid")
		return "", sveltinerr.NewDefaultError(err)
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
			return "", sveltinerr.NewMetadataTypeNotValidError()
		}
		return metadataType, nil
	default:
		return "", sveltinerr.NewMetadataTypeNotValidError()
	}
}
