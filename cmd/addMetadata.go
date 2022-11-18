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

	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/internal/composer"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/internal/tpltypes"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/tui/feedbacks"
	"github.com/sveltinio/sveltin/tui/prompts"
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
	Short:   "Add metadata to an existing resource",
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
	// Exit if running sveltin commands either from a not valid directory or not latest sveltin version.
	isValidProject(true)

	mdName, err := prompts.AskMetadataNameHandler(args)
	utils.ExitIfError(err)

	mdResource, err := prompts.SelectResourceHandler(cfg.fs, resourceName, cfg.settings)
	utils.ExitIfError(err)

	mdType, err := prompts.SelectMetadataTypeHandler(metadataType)
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
	sfs := factory.NewMetadataArtifact(&resources.SveltinTemplatesFS, cfg.fs)
	err = projectFolder.Create(sfs)
	utils.ExitIfError(err)

	cfg.log.Success("Done\n")

	// NEXT STEPS
	feedbacks.ShowNewMetadataHelpMessage(metadataTemplateData)
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
			Settings: cfg.settings,
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
			Settings: cfg.settings,
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
			Name:       helpers.GetResourceRouteFilename(item, cfg.settings),
			TemplateID: item,
			TemplateData: &config.TemplateData{
				Settings:        cfg.settings,
				Metadata:        metadataData,
				ProjectSettings: &cfg.projectSettings,
			},
		}
		resourceMedatadaRoutesFolder.Add(f)
	}

	// NEW FOLDER: src/routes/<resource_name>/[slug]
	slugFolder := composer.NewFolder("[slug]")
	// NEW FILE: src/routes/<resource_name>/[slug]{+page.svelte, +page.ts}
	for _, item := range []string{SlugFile, SlugEndpointFile} {
		f := &composer.File{
			Name:       helpers.GetResourceRouteFilename(item, cfg.settings),
			TemplateID: item,
			TemplateData: &config.TemplateData{
				Settings:        cfg.settings,
				Metadata:        metadataData,
				ProjectSettings: &cfg.projectSettings,
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
		Name:       cfg.settings.GetAPIFilename(),
		TemplateID: ApiMetadataIndex,
		TemplateData: &config.TemplateData{
			Settings: cfg.settings,
			Metadata: metadataData,
		},
	}
	resourceAPIMetadataMatcherFolder.Add(resourceMetadataIndexAPIFile)
	resourceAPIFolder.Add(resourceAPIMetadataMatcherFolder)

	// NEW FOLDER: src/routes/api/<version>/<resource_name>/<metadata_name>/[slug=string]
	resourceAPIMetadataNameMatcherFolder := composer.NewFolder("[slug=string]")

	// NEW FILE: src/routes/api/<version>/<resource_name>/<metadata_name>/[slug=string]/+server.ts
	resourceMetadataNameIndexAPIFile := &composer.File{
		Name:       cfg.settings.GetAPIFilename(),
		TemplateID: ApiFolder,
		TemplateData: &config.TemplateData{
			Settings: cfg.settings,
			Metadata: metadataData,
		},
	}
	resourceAPIMetadataNameMatcherFolder.Add(resourceMetadataNameIndexAPIFile)
	resourceAPIMetadataMatcherFolder.Add(resourceAPIMetadataNameMatcherFolder)

	apiFolder.Add(resourceAPIFolder)

	return apiFolder
}
