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

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/pkg/composer"
	"github.com/sveltinio/sveltin/pkg/sveltinerr"
	"github.com/sveltinio/sveltin/resources"
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
	Short:   "Command to add a new metadata to your content as a Sveltekit resource",
	Long: resources.GetASCIIArt() + `
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
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	mdName, err := promptMetadataName(args)
	utils.ExitIfError(err)

	mdResource, err := promptResource(cfg.fs, resourceName, cfg.sveltin)
	utils.ExitIfError(err)

	mdType, err := promptMetadataType(metadataType)
	utils.ExitIfError(err)

	cfg.log.Plain(utils.Underline(fmt.Sprintf("'%s' will be created as metadata for %s", mdName, mdResource)))

	// NEW FILE: <metadata_name>.json.ts
	cfg.log.Info("Creating the API endpoint for the metadata")
	resourceMetadataAPIFile := &composer.File{
		Name:       cfg.sveltin.GetMetadataAPIFilename(mdName),
		TemplateID: ApiFolder,
		TemplateData: &config.TemplateData{
			Name:     mdName,
			Resource: mdResource,
			Type:     mdType,
			Config:   cfg.sveltin,
		},
	}

	// NEW FOLDER: src/routes/api/<api_version>/<resource_name>
	resourceAPIFolder := composer.NewFolder(mdResource)
	resourceAPIFolder.Add(resourceMetadataAPIFile)

	// GET FOLDER: src/routes/api/<api_version> folder
	apiFolder := cfg.fsManager.GetFolder(ApiFolder)
	apiFolder.Add(resourceAPIFolder)

	// NEW FILE: api<metadata_name>.ts file into src/lib/<resource_name> folder
	cfg.log.Info("Creating the lib file for the metadata")
	libFile := &composer.File{
		Name:       cfg.pathMaker.GetResourceLibFilename(mdName),
		TemplateID: LibFolder,
		TemplateData: &config.TemplateData{
			Name:     mdName,
			Resource: mdResource,
			Type:     mdType,
			Config:   cfg.sveltin,
		},
	}
	// NEW FOLDER: src/lib/<resource_name>
	resourceLibFolder := composer.NewFolder(mdResource)
	resourceLibFolder.Add(libFile)

	// GET FOLDER: src/lib folder
	libFolder := cfg.fsManager.GetFolder(LibFolder)
	libFolder.Add(resourceLibFolder)

	// NEW FOLDER: <metadata_name>
	resourceMedatadaRoutesFolder := composer.NewFolder(mdName)

	// NEW FILE: src/routes/<resource_name>/<metadata_name>/{index.svelte, index.ts, [slug].svelte, [slug].ts}
	cfg.log.Info("Creating the components and endpoints for the metadata")
	for _, item := range []string{IndexFile, IndexEndpointFile, SlugFile, SlugEndpointFile} {
		f := &composer.File{
			Name:       helpers.GetResourceRouteFilename(item, cfg.sveltin),
			TemplateID: item,
			TemplateData: &config.TemplateData{
				Name:     mdName,
				Resource: mdResource,
				Type:     mdType,
				Config:   cfg.sveltin,
			},
		}
		resourceMedatadaRoutesFolder.Add(f)
	}

	// NEW FOLDER: src/routes/<resource_name>/<metadata_name>
	resourceRoutesFolder := composer.NewFolder(mdResource)
	resourceRoutesFolder.Add(resourceMedatadaRoutesFolder)

	// GET FOLDER: src/routes folder
	routesFolder := cfg.fsManager.GetFolder(RoutesFolder)
	routesFolder.Add(resourceRoutesFolder)

	// SET FOLDER STRUCTURE
	projectFolder := cfg.fsManager.GetFolder(RootFolder)
	projectFolder.Add(apiFolder)
	projectFolder.Add(libFolder)
	projectFolder.Add(routesFolder)

	// GENERATE FOLDER STRUCTURE
	sfs := factory.NewMetadataArtifact(&resources.SveltinFS, cfg.fs)
	err = projectFolder.Create(sfs)
	utils.ExitIfError(err)

	cfg.log.Success("Done")

	// NEXT STEPS
	cfg.log.Plain(utils.Underline("Next Steps"))
	cfg.log.Success("Metadata ready to be used.")
	cfg.log.Important("Ensure your markdown frontmatter consider it.")

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
	availableResources := helpers.GetAllResources(fs, c.GetContentPath())

	switch nameLenght := len(mdResourceFlag); {
	case nameLenght == 0:
		resourcePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide an existing resource name.",
			Label:    "Which existing resource?",
		}
		result, err := common.PromptGetSelect(resourcePromptContent, availableResources, false)
		if err != nil {
			return "", err
		}
		return utils.ToSlug(result), nil
	case nameLenght != 0:
		if !common.Contains(availableResources, mdResourceFlag) {
			return "", sveltinerr.NewResourceNotFoundError()
		}
		return utils.ToSlug(mdResourceFlag), nil
	default:
		return "", sveltinerr.NewResourceNotFoundError()
	}
}

func promptMetadataName(inputs []string) (string, error) {
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		metadataNamePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide a name for the metadata.",
			Label:    "What's the metadata name?",
		}

		result, err := common.PromptGetInput(metadataNamePromptContent, nil, "")
		if err != nil {
			return "", err
		}

		return utils.ToSlug(result), nil
	case numOfArgs == 1:
		return utils.ToSlug(inputs[0]), nil
	default:
		err := errors.New("something went wrong: name not valid")
		return "", sveltinerr.NewDefaultError(err)
	}

}

func promptMetadataType(mdTypeFlag string) (string, error) {
	promptObjects := []config.PromptObject{
		{ID: "single", Name: "(1:1) One-to-One"},
		{ID: "list", Name: "(1:m) One-to-Many"},
	}

	switch nameLenght := len(mdTypeFlag); {
	case nameLenght == 0:
		metadataTypePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide a metadata type.",
			Label:    "Which relationship between your content and the metadata?",
		}
		result, err := common.PromptGetSelect(metadataTypePromptContent, promptObjects, true)
		if err != nil {
			return "", err
		}
		return result, nil
	case nameLenght != 0:
		valid := common.GetPromptObjectKeys(promptObjects)
		if !common.Contains(valid, mdTypeFlag) {
			return "", sveltinerr.NewMetadataTypeNotValidError()
		}
		return mdTypeFlag, nil
	default:
		return "", sveltinerr.NewMetadataTypeNotValidError()
	}
}
