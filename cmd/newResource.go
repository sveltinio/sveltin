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

	"github.com/spf13/cobra"
)

//=============================================================================

var (
	group          string
	withSlugLayout bool
)

//=============================================================================

var newResourceCmd = &cobra.Command{
	Use:     "resource [name]",
	Aliases: []string{"route", "r"},
	Short:   "Create a new resource (route).",
	Long: resources.GetASCIIArt() + `
Command used to create new resources.

Why "resource" instead of "route"?
Although a resource is basically a route on SvelteKit router, a resource is not an empty route. The retional behind
the name is related to the fact that to serve content a route alone is not enough. To name a few, we need libs,
pages, endpoints, a place to serve static contents like images for the content belongs the route etc.

That's why "resource", all you need to group, serve and expose the content's route.

This command:

- Create a <resource_name> folder within "content" folder, so that you can add new content for the resource
- Add the resource as route within the "src/routes" folder, creating its own folder
- Scaffold a GET endpoint for the resource within "src/routes/api/<api_version>/<resource_name>
- Scaffold +page.svelte component and +page.serve.ts endpoint to list all the content belongs to a resource
- Scaffold [slug]/+page.svelte component and [slug]/+page.ts endpoint to get access to a specific content page
	`,
	Run: RunNewResourceCmd,
}

// RunNewResourceCmd is the actual work function.
func RunNewResourceCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	resourceName, err := promptResourceName(args)
	utils.ExitIfError(err)

	resourceData := &tpltypes.ResourceData{
		Name:       resourceName,
		Group:      group,
		SlugLayout: withSlugLayout,
	}

	// MAKE FOLDER STRUCTURE: content folder
	headingText := fmt.Sprintf("Creating '%s' as resource", resourceData.Name)
	cfg.log.Plain(markup.H1(headingText))

	contentFolder, err := makeResourceFolderStructure(ContentFolder, resourceData, cfg)
	utils.ExitIfError(err)

	// MAKE FOLDER STRUCTURE: src/params
	paramsFolder, err := makeResourceFolderStructure(ParamsFolder, resourceData, cfg)
	utils.ExitIfError(err)

	// MAKE FOLDER STRUCTURE: src/lib folder
	libFolder, err := makeResourceFolderStructure(LibFolder, resourceData, cfg)
	utils.ExitIfError(err)

	// MAKE FOLDER STRUCTURE: src/routes/<resource_name>/{index.svelte, index.ts, [slug].svelte, [slug].json.ts}
	routesFolder, err := makeResourceFolderStructure(RoutesFolder, resourceData, cfg)
	utils.ExitIfError(err)

	// MAKE FOLDER STRUCTURE: src/routes/api/<api_version> folder
	apiFolder, err := makeResourceFolderStructure(ApiFolder, resourceData, cfg)
	utils.ExitIfError(err)

	// SET FOLDER STRUCTURE
	projectFolder := cfg.fsManager.GetFolder(RootFolder)
	projectFolder.Add(contentFolder)
	projectFolder.Add(libFolder)
	projectFolder.Add(paramsFolder)
	projectFolder.Add(routesFolder)
	projectFolder.Add(apiFolder)

	// GENERATE THE FOLDER TREE
	sfs := factory.NewResourceArtifact(&resources.SveltinFS, cfg.fs)
	err = projectFolder.Create(sfs)
	utils.ExitIfError(err)

	cfg.log.Success("Done")

	// NEXT STEPS
	common.PrintHelperTextNewResource(resourceName)
}

func resourceCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&group, "group", "g", "", "Group name for resource routes (https://kit.svelte.dev/docs/advanced-routing#advanced-layouts)")
	cmd.Flags().BoolVarP(&withSlugLayout, "slug", "", false, "Use a different layout for the slug pages (https://kit.svelte.dev/docs/advanced-routing#advanced-layouts-layout)")
}

func init() {
	newCmd.AddCommand(newResourceCmd)
	resourceCmdFlags(newResourceCmd)
}

//=============================================================================

func promptResourceName(inputs []string) (string, error) {
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		resourceNamePromptContent := &input.Config{
			Placeholder: "What's the resource name? (e.g. posts, portfolio ...)",
			ErrorMsg:    "Please, provide a name for the resource.",
		}
		result, err := input.Run(resourceNamePromptContent)
		if err != nil {
			return "", err
		}
		return utils.ToSlug(result), nil
	case numOfArgs == 1:
		return utils.ToSlug(inputs[0]), nil
	default:
		err := errors.New("something went wrong: value not valid")
		return "", sveltinerr.NewDefaultError(err)
	}
}

//=============================================================================

func makeResourceFolderStructure(folderName string, resourceData *tpltypes.ResourceData, cfg appConfig) (*composer.Folder, error) {
	switch folderName {
	case ContentFolder:
		return createResourceContentLocalFolder(resourceData), nil
	case ParamsFolder:
		return createResourceParamsLocalFolder(resourceData), nil
	case LibFolder:
		return createResourceLibLocalFolder(resourceData), nil
	case RoutesFolder:
		return createResourceRoutesLocalFolder(cfg, resourceData), nil
	case ApiFolder:
		return createResourceAPIRoutesLocalFolder(resourceData), nil
	default:
		err := errors.New("something went wrong: folder not found as mapped resource for sveltin projects")
		return nil, sveltinerr.NewDefaultError(err)
	}
}

//=============================================================================

func createResourceContentLocalFolder(resourceData *tpltypes.ResourceData) *composer.Folder {
	// GET FOLDER: content folder
	contentFolder := cfg.fsManager.GetFolder(ContentFolder)

	// NEW FOLDER: content/<resource_name>. Here is where the "new content" command saves files
	cfg.log.Info("Content folder")
	resourceContentFolder := composer.NewFolder(resourceData.Name)
	contentFolder.Add(resourceContentFolder)

	return contentFolder
}

func createResourceLibLocalFolder(resourceData *tpltypes.ResourceData) *composer.Folder {
	// GET FOLDER: src/lib folder
	libFolder := cfg.fsManager.GetFolder(LibFolder)

	// NEW FOLDER: /src/lib/<resource_name>
	resourceLibFolder := composer.NewFolder(resourceData.Name)

	// NEW FILE: src/lib/<resource_name>/load<resource_name>.ts
	cfg.log.Info("Lib files")
	libFile := &composer.File{
		Name:       utils.ToLibFile(resourceData.Name),
		TemplateID: LibFolder,
		TemplateData: &config.TemplateData{
			Name:     resourceData.Name,
			Resource: resourceData,
			Settings: cfg.settings,
		},
	}
	resourceLibFolder.Add(libFile)
	libFolder.Add(resourceLibFolder)

	return libFolder
}

func createResourceParamsLocalFolder(resourceData *tpltypes.ResourceData) *composer.Folder {
	cfg.log.Info("Parameters matchers")
	// GET FOLDER: src/params folder
	paramsFolder := cfg.fsManager.GetFolder(ParamsFolder)

	// NEW FILE: src/params/string.js
	stringMatcherFile := &composer.File{
		Name:       "string.js",
		TemplateID: StringMatcher,
		TemplateData: &config.TemplateData{
			Resource: resourceData,
			Settings: cfg.settings,
		},
	}
	// Add file to folder
	paramsFolder.Add(stringMatcherFile)

	// NEW FILE: src/params/slug.js
	slugMatcherFile := &composer.File{
		Name:       "slug.js",
		TemplateID: GenericMatcher,
		TemplateData: &config.TemplateData{
			Name:     "slug",
			Resource: resourceData,
			Settings: cfg.settings,
		},
	}
	// Add file to folder
	paramsFolder.Add(slugMatcherFile)

	return paramsFolder
}

func createResourceRoutesLocalFolder(cfg appConfig, resourceData *tpltypes.ResourceData) *composer.Folder {
	// GET FOLDER: src/routes folder
	routesFolder := cfg.fsManager.GetFolder(RoutesFolder)

	// NEW FOLDER: src/routes/<resource_name>
	resourceRoutesFolder := composer.NewFolder(resourceData.Name)
	// NEW FILE: src/routes/<resource_name>/{+page.svelte, +page.server.ts}
	cfg.log.Info("Routes")
	for _, item := range []string{IndexFile, IndexEndpointFile} {
		f := &composer.File{
			Name:       helpers.GetResourceRouteFilename(item, cfg.settings),
			TemplateID: item,
			TemplateData: &config.TemplateData{
				Name:     resourceData.Name,
				Resource: resourceData,
				Settings: cfg.settings,
			},
		}
		resourceRoutesFolder.Add(f)
	}

	// NEW FOLDER: src/routes/<resource_name>/[slug]
	slugFolder := composer.NewFolder("[slug]")
	// NEW FILE: src/routes/<resource_name>/[slug]{+page.svelte, +page.ts}
	slugFiles := []string{SlugFile, SlugEndpointFile}
	if resourceData.SlugLayout {
		slugFiles = append(slugFiles, SlugLayoutFile)
	}
	for _, item := range slugFiles {
		f := &composer.File{
			Name:       helpers.GetResourceRouteFilename(item, cfg.settings),
			TemplateID: item,
			TemplateData: &config.TemplateData{
				Name:     resourceData.Name,
				Resource: resourceData,
				Settings: cfg.settings,
			},
		}
		slugFolder.Add(f)
	}
	resourceRoutesFolder.Add(slugFolder)

	if utils.IsEmpty(resourceData.Group) {
		routesFolder.Add(resourceRoutesFolder)
	} else {
		// NEW FOLDER: src/routes/(group_name)/<resource_name>
		resourceGroupRoutesFolder := composer.NewFolder(fmt.Sprintf("(%s)", resourceData.Group))
		resourceGroupRoutesFolder.Add(resourceRoutesFolder)
		routesFolder.Add(resourceGroupRoutesFolder)

	}

	return routesFolder
}

func createResourceAPIRoutesLocalFolder(resourceData *tpltypes.ResourceData) *composer.Folder {
	cfg.log.Info("REST endpoints")
	// GET FOLDER: src/routes/api/<api_version> folder
	apiFolder := cfg.fsManager.GetFolder(ApiFolder)

	// NEW FOLDER: src/routes/api/<version>/<resource_name>
	resourceAPIFolder := composer.NewFolder(resourceData.Name)

	// NEW FILE: src/routes/api/<version>/<resource_name>/+server.ts
	apiFile := &composer.File{
		Name:       cfg.settings.GetAPIFilename(),
		TemplateID: ApiIndexFile,
		TemplateData: &config.TemplateData{
			Name:     resourceData.Name,
			Resource: resourceData,
			Settings: cfg.settings,
		},
	}
	resourceAPIFolder.Add(apiFile)

	// NEW FOLDER: src/routes/api/<version>/<resource_name>/[slug=string]
	slugStringFolder := composer.NewFolder(fmt.Sprintf("%s%s%s%s%s", "[", "slug", "=", "string", "]"))
	// NEW FILE: src/routes/api/<version>/<resource_name>/[slug=string]/+server.ts
	apiSlugFile := &composer.File{
		Name:       cfg.settings.GetAPIFilename(),
		TemplateID: ApiSlugFile,
		TemplateData: &config.TemplateData{
			Name:     resourceData.Name,
			Resource: resourceData,
			Settings: cfg.settings,
		},
	}
	slugStringFolder.Add(apiSlugFile)
	resourceAPIFolder.Add(slugStringFolder)

	// Add folders to src/routes/api/<version>/
	apiFolder.Add(resourceAPIFolder)

	return apiFolder
}
