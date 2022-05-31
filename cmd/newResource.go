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

	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/pkg/composer"
	"github.com/sveltinio/sveltin/pkg/sveltinerr"
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

What is a "resource" for Sveltin?
A resource is a way to group, serve and expose your content.

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

	// GET FOLDER: content folder
	contentFolder := cfg.fsManager.GetFolder(ContentFolder)

	cfg.log.Plain(utils.Underline(fmt.Sprintf("'%s' will be created as resource", resourceName)))

	// NEW FOLDER: content/<resource_name>. Here is where the "new content" command saves files
	cfg.log.Info("Creating the content folder for your resource")
	resourceContentFolder := composer.NewFolder(resourceName)
	contentFolder.Add(resourceContentFolder)

	// GET FOLDER: src/routes/api/<api_version> folder
	apiFolder := cfg.fsManager.GetFolder(ApiFolder)

	// NEW FOLDER: src/routes/api/<version>/<resource_name>
	resourceAPIFolder := composer.NewFolder(resourceName)

	// NEW FILE: src/routes/api/<resource_name>/published.json.ts
	cfg.log.Info("Creating the API endpoint for the resource")
	apiFile := &composer.File{
		Name:       cfg.sveltin.GetAPIFilename(),
		TemplateID: ApiFolder,
		TemplateData: &config.TemplateData{
			Name:   resourceName,
			Config: cfg.sveltin,
		},
	}
	resourceAPIFolder.Add(apiFile)
	apiFolder.Add(resourceAPIFolder)

	// GET FOLDER: src/lib folder
	libFolder := cfg.fsManager.GetFolder(LibFolder)

	// NEW FOLDER: /src/lib/<resource_name>
	resourceLibFolder := composer.NewFolder(resourceName)

	// NEW FILE: src/lib/<resource_name>/api<resource_name>.ts
	cfg.log.Info("Creating the lib file for the resource")
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

	// GET FOLDER: src/routes folder
	routesFolder := cfg.fsManager.GetFolder(RoutesFolder)

	// NEW FOLDER: src/routes/<resource_name>
	resourceRoutesFolder := composer.NewFolder(resourceName)

	// NEW FILE: src/routes/<resource_name>/{index.svelte, index.ts, [slug].svelte, [slug].ts}
	cfg.log.Info("Creating the components and endpoints for the resource")
	for _, item := range []string{IndexFile, IndexEndpointFile, SlugFile, SlugEndpointFile} {
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
	routesFolder.Add(resourceRoutesFolder)

	// SET FOLDER STRUCTURE
	projectFolder := cfg.fsManager.GetFolder(RootFolder)
	projectFolder.Add(contentFolder)
	projectFolder.Add(apiFolder)
	projectFolder.Add(libFolder)
	projectFolder.Add(routesFolder)

	// CREATE FOLDER STRUCTURE
	sfs := factory.NewResourceArtifact(&resources.SveltinFS, cfg.fs)
	err = projectFolder.Create(sfs)
	utils.ExitIfError(err)

	cfg.log.Success("Done")

	// NEXT STEPS
	cfg.log.Plain(utils.Underline("Next Steps"))
	cfg.log.Success("Resource ready to be used. Start by adding content to it.")
	cfg.log.Important(fmt.Sprintf("Eg: sveltin new content %s/getting-started", resourceName))
}

func init() {
	newCmd.AddCommand(newResourceCmd)
}

//=============================================================================

func promptResourceName(inputs []string) (string, error) {
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		resourceNamePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide a name for the resource.",
			Label:    "What's the resource name? (e.g. posts, portfolio ...)",
		}
		result, err := common.PromptGetInput(resourceNamePromptContent, nil, "")
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
