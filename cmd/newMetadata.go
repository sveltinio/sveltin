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

	metadataTemplateData := &config.TemplateData{
		Name:     mdName,
		Resource: mdResource,
		Type:     mdType,
		Config:   cfg.sveltin,
	}

	cfg.log.Plain(utils.Underline(fmt.Sprintf("Creating '%s' as metadata for the '%s' resource", metadataTemplateData.Name, metadataTemplateData.Resource)))

	// MAKE FOLDER STRUCTURE: src/lib folder
	libFolder, err := makeOrAddContentForMetadataToProjectStructure(LibFolder, metadataTemplateData)
	utils.ExitIfError(err)

	paramsFolder, err := makeOrAddContentForMetadataToProjectStructure(ParamsFolder, metadataTemplateData)
	utils.ExitIfError(err)

	// MAKE FOLDER STRUCTURE: src/routes/<resource_name>/<metadata_name>/{index.svelte, index.ts, [slug].svelte, [slug].ts}
	routesFolder, err := makeOrAddContentForMetadataToProjectStructure(RoutesFolder, metadataTemplateData)
	utils.ExitIfError(err)

	// MAKE FOLDER STRUCTURE: src/routes/api/<api_version> folder
	apiFolder, err := makeOrAddContentForMetadataToProjectStructure(ApiFolder, metadataTemplateData)
	utils.ExitIfError(err)

	// SET FOLDER STRUCTURE
	projectFolder := cfg.fsManager.GetFolder(RootFolder)
	projectFolder.Add(libFolder)
	projectFolder.Add(paramsFolder)
	projectFolder.Add(routesFolder)
	projectFolder.Add(apiFolder)

	// GENERATE THE FOLDER TREE
	sfs := factory.NewMetadataArtifact(&resources.SveltinFS, cfg.fs)
	err = projectFolder.Create(sfs)
	utils.ExitIfError(err)

	cfg.log.Success("Done")

	// NEXT STEPS
	cfg.log.Plain(utils.Underline("Next Steps"))
	cfg.log.Important(common.HelperTextNewMetadata(metadataTemplateData))
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

//=============================================================================

func makeOrAddContentForMetadataToProjectStructure(folderName string, metadataTemaplateData *config.TemplateData) (*composer.Folder, error) {
	switch folderName {
	case LibFolder:
		return createOrAddContentForMetadataToLibLocalFolder(metadataTemaplateData), nil
	case ParamsFolder:
		return createOrAddContentForMetadataToParamsLocalFolder(metadataTemaplateData), nil
	case RoutesFolder:
		return createOrAddContentForMetadataToRoutesLocalFolder(metadataTemaplateData), nil
	case ApiFolder:
		return createOrAddContentForMetadataToApiLocalFolder(metadataTemaplateData), nil
	default:
		err := errors.New("something went wrong: folder not found as mapped resource for sveltin projects")
		return nil, sveltinerr.NewDefaultError(err)
	}
}

//=============================================================================

func createOrAddContentForMetadataToLibLocalFolder(metadataTemplateData *config.TemplateData) *composer.Folder {
	// NEW FILE: api<metadata_name>.ts file into src/lib/<resource_name> folder
	cfg.log.Info("Lib files")
	libFile := &composer.File{
		Name:         cfg.pathMaker.GetResourceLibFilename(metadataTemplateData.Name),
		TemplateID:   LibFolder,
		TemplateData: metadataTemplateData,
	}
	// NEW FOLDER: src/lib/<resource_name>
	resourceLibFolder := composer.NewFolder(metadataTemplateData.Resource)
	resourceLibFolder.Add(libFile)

	// GET FOLDER: src/lib folder
	libFolder := cfg.fsManager.GetFolder(LibFolder)
	libFolder.Add(resourceLibFolder)

	return libFolder
}

func createOrAddContentForMetadataToParamsLocalFolder(metadataTemplateData *config.TemplateData) *composer.Folder {
	cfg.log.Info("Parameters matchers")
	// GET FOLDER: src/params folder
	paramsFolder := cfg.fsManager.GetFolder(ParamsFolder)

	// NEW FILE: src/params/<metadata_name>.js
	metadataMatcherFile := &composer.File{
		Name:         fmt.Sprintf("%s%s", metadataTemplateData.Name, ".js"),
		TemplateID:   GenericMatcher,
		TemplateData: metadataTemplateData,
	}
	// Add file to folder
	paramsFolder.Add(metadataMatcherFile)

	return paramsFolder
}

func createOrAddContentForMetadataToRoutesLocalFolder(metadataTemaplateData *config.TemplateData) *composer.Folder {
	cfg.log.Info("Routes")
	// NEW FOLDER: <metadata_name>
	resourceMedatadaRoutesFolder := composer.NewFolder(metadataTemaplateData.Name)

	// NEW FILE: src/routes/<resource_name>/<metadata_name>/{index.svelte, index.ts, [slug].svelte, [slug].ts}
	for _, item := range []string{IndexFile, IndexEndpointFile, SlugFile, SlugEndpointFile} {
		f := &composer.File{
			Name:         helpers.GetResourceRouteFilename(item, cfg.sveltin),
			TemplateID:   item,
			TemplateData: metadataTemaplateData,
		}
		resourceMedatadaRoutesFolder.Add(f)
	}

	// NEW FOLDER: src/routes/<resource_name>/<metadata_name>
	resourceRoutesFolder := composer.NewFolder(metadataTemaplateData.Resource)
	resourceRoutesFolder.Add(resourceMedatadaRoutesFolder)

	// GET FOLDER: src/routes folder
	routesFolder := cfg.fsManager.GetFolder(RoutesFolder)
	routesFolder.Add(resourceRoutesFolder)

	return routesFolder
}

func createOrAddContentForMetadataToApiLocalFolder(metadataTemplateData *config.TemplateData) *composer.Folder {
	cfg.log.Info("REST endpoint")
	// GET FOLDER: src/routes/api/<api_version> folder
	apiFolder := cfg.fsManager.GetFolder(ApiFolder)

	// NEW FOLDER: src/routes/api/<api_version>/<resource_name>
	resourceAPIFolder := composer.NewFolder(metadataTemplateData.Resource)

	// NEW FOLDER: src/routes/api/<version>/<resource_name>/[<resource_name> = <metadata_name>]
	resourceAPIMetadataMatcherFolder := composer.NewFolder(fmt.Sprintf("%s%s%s%s%s", "[", metadataTemplateData.Resource, "=", metadataTemplateData.Name, "]"))

	// NEW FILE: src/routes/api/<version>/<resource_name>/[<resource_name> = <metadata_name>]/index.ts
	resourceMetadataIndexAPIFile := &composer.File{
		Name:         cfg.sveltin.GetAPIFilename(),
		TemplateID:   ApiMetadataIndex,
		TemplateData: metadataTemplateData,
	}
	resourceAPIMetadataMatcherFolder.Add(resourceMetadataIndexAPIFile)
	resourceAPIFolder.Add(resourceAPIMetadataMatcherFolder)

	// NEW FOLDER: src/routes/api/<version>/<resource_name>/[<resource_name> = <metadata_name>]/[<metadata_name> = string]
	resourceAPIMetadataNameMatcherFolder := composer.NewFolder(fmt.Sprintf("%s%s%s%s%s", "[", metadataTemplateData.Name, "=", "string", "]"))

	// NEW FILE: src/routes/api/<version>/<resource_name>/[<resource_name> = <metadata_name>]/[<metadata_name> = string]/index.ts
	resourceMetadataNameIndexAPIFile := &composer.File{
		Name:         cfg.sveltin.GetAPIFilename(),
		TemplateID:   ApiFolder,
		TemplateData: metadataTemplateData,
	}
	resourceAPIMetadataNameMatcherFolder.Add(resourceMetadataNameIndexAPIFile)
	resourceAPIMetadataMatcherFolder.Add(resourceAPIMetadataNameMatcherFolder)

	apiFolder.Add(resourceAPIFolder)

	return apiFolder
}
