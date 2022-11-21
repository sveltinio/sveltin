/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/internal/composer"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/internal/tpltypes"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/tui/prompts"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var pageType string

//=============================================================================

var newPageCmd = &cobra.Command{
	Use:     "page [name]",
	Aliases: []string{"p"},
	Short:   "Create a new page route",
	Long: resources.GetASCIIArt() + `
Command used to create a new public page route.

Pages are Svelte components written in .svelte or .svx (for markdown) files. The filename determines the route,
so creating a page named "about" will generate the following route /about/+page.(svelte|svx)

This command allows you to select between a svelte component page and a markdown page.
`,
	Run: NewPageCmdRun,
}

// NewPageCmdRun is the actual work function.
func NewPageCmdRun(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands either from a not valid directory or not latest sveltin version.
	isValidProject(true)

	pageName, err := prompts.AskPageNameHandler(args)
	utils.ExitIfError(err)

	pageType, err := prompts.SelectPageTypeHandler(pageType)
	utils.ExitIfError(err)

	pageData := &tpltypes.PageData{
		Name: pageName,
		Type: pageType,
	}

	headingText := fmt.Sprintf("Creating the '%s' page (type: %s)", pageName, pageType)
	cfg.log.Plain(markup.H1(headingText))

	// GET FOLDER: src/routes
	routesFolder := cfg.fsManager.GetFolder(RoutesFolder)

	// NEW FOLDER: src/routes/<page_name>
	pageFolder := composer.NewFolder(pageName)
	// NEW FILE: src/routes/<page_name>/+page.svelte|svx>
	pageFile := cfg.fsManager.NewPublicPageFile(pageData, &cfg.projectSettings)
	utils.ExitIfError(err)

	// ADD TO THE ROUTES FOLDER
	pageFolder.Add(pageFile)
	routesFolder.Add(pageFolder)

	// SET FOLDER STRUCTURE
	projectFolder := cfg.fsManager.GetFolder(RootFolder)
	projectFolder.Add(routesFolder)

	// GENERATE THE FOLDER TREE
	sfs := factory.NewPageArtifact(&resources.SveltinTemplatesFS, cfg.fs)
	err = projectFolder.Create(sfs)
	utils.ExitIfError(err)
	cfg.log.Success("Done\n")
}

func pageCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&pageType, "as", "a", "", "Sveltekit page as Svelte component or markdown via mdsvex. Possible values: svelte, markdown")
}

func init() {
	newCmd.AddCommand(newPageCmd)
	pageCmdFlags(newPageCmd)
}
