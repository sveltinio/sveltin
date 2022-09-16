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

	"github.com/charmbracelet/bubbles/list"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/sveltinio/prompti/choose"
	"github.com/sveltinio/prompti/input"
	"github.com/sveltinio/sveltin/common"
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
	resourceName string
	metadataType string
)

//=============================================================================

var addMetadataCmd = &cobra.Command{
	Use:     "metadata [name] --to [resource] --as [single|list]",
	Aliases: []string{"m"},
	Short:   "Add metadata to an existing resource.",
	Long: resources.GetASCIIArt() + `
Command used to add new metadata for your content to an existing resource.

**Note**: This command needs an existing resource created by running: sveltin new resource <resource_name>.

What is a "metadata" for Sveltin?
Whatever you enter in the front-matter of your markdown content for which you want content grouped by it.

Metadata Types:

- single: 1:1 relationship (e.g. category)
- list: 1:many relationship (e.g. tags)
`,
	Run: RunAddMetadataCmd,
}

// RunAddMetadataCmd is the actual work function.
func RunAddMetadataCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	mdName, err := promptMetadataName(args)
	utils.ExitIfError(err)

	mdResource, err := promptResource(cfg.fs, resourceName, cfg.sveltin)
	utils.ExitIfError(err)

	mdType, err := promptMetadataType(metadataType)
	utils.ExitIfError(err)

	metadataTemplateData := &tpltypes.MetadataData{
		Name:     mdName,
		Resource: mdResource,
		Type:     mdType,
	}

	headingText := fmt.Sprintf("Creating '%s' as metadata for the '%s' resource", metadataTemplateData.Name, metadataTemplateData.Resource)
	cfg.log.Plain(markup.H1(headingText))

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
	common.PrintHelperTextNewMetadata(metadataTemplateData)
}

func metadataCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&resourceName, "to", "t", "", "Name of the resource the new metadata is belongs to.")
	cmd.Flags().StringVarP(&metadataType, "as", "a", "", "Type of the new metadata. (possible values: single or list)")
}

func init() {
	metadataCmdFlags(addMetadataCmd)
	addCmd.AddCommand(addMetadataCmd)
}

//=============================================================================

func promptResource(fs afero.Fs, mdResourceFlag string, c *config.SveltinConfig) (string, error) {
	availableResources := helpers.GetAllResources(fs, c.GetContentPath())

	options := choose.ToListItem(availableResources)

	switch nameLenght := len(mdResourceFlag); {
	case nameLenght == 0:
		resourcePromptContent := &choose.Config{
			Title:    "Which existing resource?",
			ErrorMsg: "Please, provide an existing resource name.",
		}

		//result, err := common.PromptGetSelect(resourcePromptContent, availableResources, false)
		result, err := choose.Run(resourcePromptContent, options)
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
		metadataNamePromptContent := &input.Config{
			Placeholder: "What's the metadata name?",
			ErrorMsg:    "Please, provide a name for the metadata.",
		}

		result, err := input.Run(metadataNamePromptContent)
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
	entries := []list.Item{
		choose.Item{Name: "single", Desc: "(1:1) One-to-One"},
		choose.Item{Name: "list", Desc: "(1:m) One-to-Many"},
	}

	switch nameLenght := len(mdTypeFlag); {
	case nameLenght == 0:
		metadataTypePromptContent := &choose.Config{
			Title:    "Which relationship between your content and the metadata?",
			ErrorMsg: "Please, provide a metadata type.",
		}
		result, err := choose.Run(metadataTypePromptContent, entries)
		if err != nil {
			return "", err
		}
		return result, nil
	case nameLenght != 0:
		valid := choose.GetItemsKeys(entries)
		if !common.Contains(valid, mdTypeFlag) {
			return "", sveltinerr.NewMetadataTypeNotValidError()
		}
		return mdTypeFlag, nil
	default:
		return "", sveltinerr.NewMetadataTypeNotValidError()
	}
}

//=============================================================================

func makeOrAddContentForMetadataToProjectStructure(folderName string, metadataData *tpltypes.MetadataData) (*composer.Folder, error) {
	switch folderName {
	case LibFolder:
		return createOrAddContentForMetadataToLibLocalFolder(metadataData), nil
	case ParamsFolder:
		return createOrAddContentForMetadataToParamsLocalFolder(metadataData), nil
	case RoutesFolder:
		return createOrAddContentForMetadataToRoutesLocalFolder(metadataData), nil
	case ApiFolder:
		return createOrAddContentForMetadataToApiLocalFolder(metadataData), nil
	default:
		err := errors.New("something went wrong: folder not found as mapped resource for sveltin projects")
		return nil, sveltinerr.NewDefaultError(err)
	}
}

//=============================================================================

func createOrAddContentForMetadataToLibLocalFolder(metadataData *tpltypes.MetadataData) *composer.Folder {
	// NEW FILE: api<metadata_name>.ts file into src/lib/<resource_name> folder
	cfg.log.Info("Lib files")
	libFile := &composer.File{
		Name:       cfg.pathMaker.GetResourceLibFilename(metadataData.Name),
		TemplateID: LibFolder,
		TemplateData: &config.TemplateData{
			Config:   cfg.sveltin,
			Metadata: metadataData,
		},
	}
	// NEW FOLDER: src/lib/<resource_name>
	resourceLibFolder := composer.NewFolder(metadataData.Resource)
	resourceLibFolder.Add(libFile)

	// GET FOLDER: src/lib folder
	libFolder := cfg.fsManager.GetFolder(LibFolder)
	libFolder.Add(resourceLibFolder)

	return libFolder
}

func createOrAddContentForMetadataToParamsLocalFolder(metadataData *tpltypes.MetadataData) *composer.Folder {
	cfg.log.Info("Parameters matchers")
	// GET FOLDER: src/params folder
	paramsFolder := cfg.fsManager.GetFolder(ParamsFolder)

	// NEW FILE: src/params/<metadata_name>.js
	metadataMatcherFile := &composer.File{
		Name:       fmt.Sprintf("%s%s", utils.ToSnakeCase(metadataData.Name), ".js"),
		TemplateID: GenericMatcher,
		TemplateData: &config.TemplateData{
			Config:   cfg.sveltin,
			Metadata: metadataData,
		},
	}
	// Add file to folder
	paramsFolder.Add(metadataMatcherFile)

	return paramsFolder
}

func createOrAddContentForMetadataToRoutesLocalFolder(metadataData *tpltypes.MetadataData) *composer.Folder {
	cfg.log.Info("Routes")
	// NEW FOLDER: <metadata_name>
	resourceMedatadaRoutesFolder := composer.NewFolder(metadataData.Name)

	// NEW FILE: src/routes/<resource_name>/<metadata_name>/{+page.svelte, +page.server.ts}
	for _, item := range []string{IndexFile, IndexEndpointFile} {
		f := &composer.File{
			Name:       helpers.GetResourceRouteFilename(item, cfg.sveltin),
			TemplateID: item,
			TemplateData: &config.TemplateData{
				Config:   cfg.sveltin,
				Metadata: metadataData,
			},
		}
		resourceMedatadaRoutesFolder.Add(f)
	}

	// NEW FOLDER: src/routes/<resource_name>/[slug]
	slugFolder := composer.NewFolder("[slug]")
	// NEW FILE: src/routes/<resource_name>/[slug]{+page.svelte, +page.ts}
	for _, item := range []string{SlugFile, SlugEndpointFile} {
		f := &composer.File{
			Name:       helpers.GetResourceRouteFilename(item, cfg.sveltin),
			TemplateID: item,
			TemplateData: &config.TemplateData{
				Config:   cfg.sveltin,
				Metadata: metadataData,
			},
		}
		slugFolder.Add(f)
	}
	resourceMedatadaRoutesFolder.Add(slugFolder)

	// NEW FOLDER: src/routes/<resource_name>/<metadata_name>
	resourceRoutesFolder := composer.NewFolder(metadataData.Resource)
	resourceRoutesFolder.Add(resourceMedatadaRoutesFolder)

	// GET FOLDER: src/routes folder
	routesFolder := cfg.fsManager.GetFolder(RoutesFolder)
	routesFolder.Add(resourceRoutesFolder)

	return routesFolder
}

func createOrAddContentForMetadataToApiLocalFolder(metadataData *tpltypes.MetadataData) *composer.Folder {
	cfg.log.Info("REST endpoint")
	// GET FOLDER: src/routes/api/<api_version> folder
	apiFolder := cfg.fsManager.GetFolder(ApiFolder)

	// NEW FOLDER: src/routes/api/<api_version>/<resource_name>
	resourceAPIFolder := composer.NewFolder(metadataData.Resource)

	// NEW FOLDER: src/routes/api/<version>/<resource_name>/<metadata_name>
	resourceAPIMetadataMatcherFolder := composer.NewFolder(utils.ToSnakeCase(metadataData.Name))

	// NEW FILE: src/routes/api/<version>/<resource_name>/[<resource_name> = <metadata_name>]/+server.ts
	resourceMetadataIndexAPIFile := &composer.File{
		Name:       cfg.sveltin.GetAPIFilename(),
		TemplateID: ApiMetadataIndex,
		TemplateData: &config.TemplateData{
			Config:   cfg.sveltin,
			Metadata: metadataData,
		},
	}
	resourceAPIMetadataMatcherFolder.Add(resourceMetadataIndexAPIFile)
	resourceAPIFolder.Add(resourceAPIMetadataMatcherFolder)

	// NEW FOLDER: src/routes/api/<version>/<resource_name>/<metadata_name>/[slug=string]
	resourceAPIMetadataNameMatcherFolder := composer.NewFolder("[slug=string]")

	// NEW FILE: src/routes/api/<version>/<resource_name>/<metadata_name>/[slug=string]/+server.ts
	resourceMetadataNameIndexAPIFile := &composer.File{
		Name:       cfg.sveltin.GetAPIFilename(),
		TemplateID: ApiFolder,
		TemplateData: &config.TemplateData{
			Config:   cfg.sveltin,
			Metadata: metadataData,
		},
	}
	resourceAPIMetadataNameMatcherFolder.Add(resourceMetadataNameIndexAPIFile)
	resourceAPIMetadataMatcherFolder.Add(resourceAPIMetadataNameMatcherFolder)

	apiFolder.Add(resourceAPIFolder)

	return apiFolder
}
