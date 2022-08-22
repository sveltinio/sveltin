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
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"

	"github.com/spf13/cobra"
)

//=============================================================================

var newResourceCmd = &cobra.Command{
	Use:     "resource [name]",
	Aliases: []string{"r"},
	Short:   "Command to create a new resource",
	Long: resources.GetASCIIArt() + `
Command to create new resources.

Why "resource" instead of "route"?
Although a resource is basically a route on SvelteKit router, a resource is not an empty route. The retional behind
the name is related to the fact that to serve content a route alone is not enough. To name a few, we need libs,
pages, endpoints, a place to serve static contents like images for the content belongs the route etc.

That's why "resource", all you need to group, serve and expose the content's route.

This command:

- Create a <resource_name> folder within "content" folder, so that you can add new content for the resource
- Add the resource as route within the "src/routes" folder, creating its own folder
- Scaffold a GET endpoint for the resource within "src/routes/api/<api_version>/<resource_name>
- Scaffold index.svelte component and index.ts endpoint to list all the content belongs to a resource
- Scaffold [slug].svelte component and [slug].ts endpoint to get access to a specific content page
	`,
	Run: RunNewResourceCmd,
}

// RunNewResourceCmd is the actual work function.
func RunNewResourceCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	resourceName, err := promptResourceName(args)
	utils.ExitIfError(err)

	// MAKE FOLDER STRUCTURE: content folder
	headingText := fmt.Sprintf("Creating '%s' as resource", resourceName)
	cfg.log.Plain(markup.H1(headingText))

	contentFolder, err := makeResourceFolderStructure(ContentFolder, resourceName, cfg)
	utils.ExitIfError(err)

	// MAKE FOLDER STRUCTURE: src/params
	paramsFolder, err := makeResourceFolderStructure(ParamsFolder, resourceName, cfg)
	utils.ExitIfError(err)

	// MAKE FOLDER STRUCTURE: src/lib folder
	libFolder, err := makeResourceFolderStructure(LibFolder, resourceName, cfg)
	utils.ExitIfError(err)

	// MAKE FOLDER STRUCTURE: src/routes/<resource_name>/{index.svelte, index.ts, [slug].svelte, [slug].json.ts}
	routesFolder, err := makeResourceFolderStructure(RoutesFolder, resourceName, cfg)
	utils.ExitIfError(err)

	// MAKE FOLDER STRUCTURE: src/routes/api/<api_version> folder
	apiFolder, err := makeResourceFolderStructure(ApiFolder, resourceName, cfg)
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
	//cfg.log.Plain(utils.Underline("Next Steps"))
	//cfg.log.Success("Resource ready to be used. Start by adding content to it.")
	//cfg.log.Important(fmt.Sprintf("Eg: sveltin new content %s/getting-started", resourceName))
}

func init() {
	newCmd.AddCommand(newResourceCmd)
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

func makeResourceFolderStructure(folderName string, resourceName string, cfg appConfig) (*composer.Folder, error) {
	switch folderName {
	case ContentFolder:
		return createResourceContentLocalFolder(resourceName), nil
	case ParamsFolder:
		return createResourceParamsLocalFolder(), nil
	case LibFolder:
		return createResourceLibLocalFolder(resourceName), nil
	case RoutesFolder:
		return createResourceRoutesLocalFolder(cfg, resourceName), nil
	case ApiFolder:
		return createResourceAPIRoutesLocalFolder(resourceName), nil
	default:
		err := errors.New("something went wrong: folder not found as mapped resource for sveltin projects")
		return nil, sveltinerr.NewDefaultError(err)
	}
}

//=============================================================================

func createResourceContentLocalFolder(resourceName string) *composer.Folder {
	// GET FOLDER: content folder
	contentFolder := cfg.fsManager.GetFolder(ContentFolder)

	// NEW FOLDER: content/<resource_name>. Here is where the "new content" command saves files
	cfg.log.Info("Content folder")
	resourceContentFolder := composer.NewFolder(resourceName)
	contentFolder.Add(resourceContentFolder)

	return contentFolder
}

func createResourceLibLocalFolder(resourceName string) *composer.Folder {
	// GET FOLDER: src/lib folder
	libFolder := cfg.fsManager.GetFolder(LibFolder)

	// NEW FOLDER: /src/lib/<resource_name>
	resourceLibFolder := composer.NewFolder(resourceName)

	// NEW FILE: src/lib/<resource_name>/load<resource_name>.ts
	cfg.log.Info("Lib files")
	libFile := &composer.File{
		Name:       utils.ToLibFile(resourceName),
		TemplateID: LibFolder,
		TemplateData: &config.TemplateData{
			Name:   resourceName,
			Config: cfg.sveltin,
		},
	}
	resourceLibFolder.Add(libFile)
	libFolder.Add(resourceLibFolder)

	return libFolder
}

func createResourceParamsLocalFolder() *composer.Folder {
	cfg.log.Info("Parameters matchers")
	// GET FOLDER: src/params folder
	paramsFolder := cfg.fsManager.GetFolder(ParamsFolder)

	// NEW FILE: src/params/string.js
	stringMatcherFile := &composer.File{
		Name:       "string.js",
		TemplateID: StringMatcher,
		TemplateData: &config.TemplateData{
			Config: cfg.sveltin,
		},
	}
	// Add file to folder
	paramsFolder.Add(stringMatcherFile)

	// NEW FILE: src/params/slug.js
	slugMatcherFile := &composer.File{
		Name:       "slug.js",
		TemplateID: GenericMatcher,
		TemplateData: &config.TemplateData{
			Name:   "slug",
			Config: cfg.sveltin,
		},
	}
	// Add file to folder
	paramsFolder.Add(slugMatcherFile)

	return paramsFolder
}

func createResourceRoutesLocalFolder(cfg appConfig, resourceName string) *composer.Folder {
	// GET FOLDER: src/routes folder
	routesFolder := cfg.fsManager.GetFolder(RoutesFolder)

	// NEW FOLDER: src/routes/<resource_name>
	resourceRoutesFolder := composer.NewFolder(resourceName)
	// NEW FILE: src/routes/<resource_name>/{+page.svelte, +page.server.ts}
	cfg.log.Info("Routes")
	for _, item := range []string{IndexFile, IndexEndpointFile} {
		f := &composer.File{
			Name:       helpers.GetResourceRouteFilename(item, cfg.sveltin),
			TemplateID: item,
			TemplateData: &config.TemplateData{
				Name:   resourceName,
				Config: cfg.sveltin,
			},
		}
		resourceRoutesFolder.Add(f)
	}

	// NEW FOLDER: src/routes/<resource_name>/[slug]
	slugFolder := composer.NewFolder("[slug]")
	// NEW FILE: src/routes/<resource_name>/[slug]{+page.svelte, +page.server.ts}
	for _, item := range []string{SlugFile, SlugEndpointFile} {
		f := &composer.File{
			Name:       helpers.GetResourceRouteFilename(item, cfg.sveltin),
			TemplateID: item,
			TemplateData: &config.TemplateData{
				Name:   resourceName,
				Config: cfg.sveltin,
			},
		}
		slugFolder.Add(f)
	}
	resourceRoutesFolder.Add(slugFolder)
	routesFolder.Add(resourceRoutesFolder)

	return routesFolder
}

func createResourceAPIRoutesLocalFolder(resourceName string) *composer.Folder {
	cfg.log.Info("REST endpoints")
	// GET FOLDER: src/routes/api/<api_version> folder
	apiFolder := cfg.fsManager.GetFolder(ApiFolder)

	// NEW FOLDER: src/routes/api/<version>/<resource_name>
	resourceAPIFolder := composer.NewFolder(resourceName)

	// NEW FILE: src/routes/api/<version>/<resource_name>/+server.ts
	apiFile := &composer.File{
		Name:       cfg.sveltin.GetAPIFilename(),
		TemplateID: ApiIndexFile,
		TemplateData: &config.TemplateData{
			Name:   resourceName,
			Config: cfg.sveltin,
		},
	}
	resourceAPIFolder.Add(apiFile)

	// NEW FOLDER: src/routes/api/<version>/<resource_name>/[slug=string]
	slugStringFolder := composer.NewFolder(fmt.Sprintf("%s%s%s%s%s", "[", "slug", "=", "string", "]"))
	// NEW FILE: src/routes/api/<version>/<resource_name>/[slug=string]/+server.ts
	apiSlugFile := &composer.File{
		Name:       cfg.sveltin.GetAPIFilename(),
		TemplateID: ApiSlugFile,
		TemplateData: &config.TemplateData{
			Name:   resourceName,
			Config: cfg.sveltin,
		},
	}
	slugStringFolder.Add(apiSlugFile)
	resourceAPIFolder.Add(slugStringFolder)

	// Add folders to src/routes/api/<version>/
	apiFolder.Add(resourceAPIFolder)

	return apiFolder
}
