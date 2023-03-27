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
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/internal/composer"
	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/internal/tpltypes"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/tui/activehelps"
	"github.com/sveltinio/sveltin/tui/prompts"
	"github.com/sveltinio/sveltin/utils"
)

var (
	// How to use the command.
	newPageCmdExample = "sveltin new page about --markdown"
	// Short description shown in the 'help' output.
	newPageCmdShortMsg = "Create a new page route"
	// Long message shown in the 'help <this-command>' output.
	newPageCmdLongMsg = utils.MakeCmdLongMsg(`Command used to create a new public page route selecting between a svelte component-based page and a markdown page.

Pages are Svelte components written in .svelte or .svx (for markdown) files.
The filename determines the route so, creating a page named "about" will generate the following route /about/+page.(svelte|svx).`)
)

// Bind command flags.
var (
	withSvelte   bool
	withMarkdown bool
)

//=============================================================================

var newPageCmd = &cobra.Command{
	Use:               "page [name] --[svelte|markdown]",
	Aliases:           []string{"p"},
	GroupID:           newCmdGroupId,
	Example:           newPageCmdExample,
	Short:             newPageCmdShortMsg,
	Long:              newPageCmdLongMsg,
	ValidArgsFunction: newPageCmdValidArgs,
	Run:               NewPageCmdRun,
}

// NewPageCmdRun is the actual work function.
func NewPageCmdRun(cmd *cobra.Command, args []string) {
	pageName, err := prompts.AskPageNameHandler(args)
	utils.ExitIfError(err)

	resultPath := filepath.Join(cfg.pathMaker.GetRoutesFolder(), pageName)
	utils.ExitIfExists(cfg.fs, resultPath)

	var pageLanguage string
	if withSvelte {
		pageLanguage = tpltypes.Svelte
	} else if withMarkdown {
		pageLanguage = tpltypes.Markdown
	}

	pageLanguage, err = prompts.SelectPageLanguageHandler(pageLanguage)
	utils.ExitIfError(err)

	pageData := &tpltypes.PageData{
		Name:     pageName,
		Language: pageLanguage,
	}

	headingText := fmt.Sprintf("Creating the '%s' page (type: %s)", pageName, pageLanguage)
	cfg.log.Plain(markup.H1(headingText))

	// MAKE FOLDER STRUCTURE: src/routes/<page_name>
	routesFolder, err := makePageFolderStructure(RoutesFolder, pageData)
	utils.ExitIfError(err)

	// MAKE FOLDER STRUCTURE: static/pages/<page_name>
	staticFolder, err := makePageFolderStructure(StaticFolder, pageData)
	utils.ExitIfError(err)

	// SET FOLDER STRUCTURE
	projectFolder := cfg.fsManager.GetFolder(RootFolder)
	projectFolder.Add(routesFolder)
	projectFolder.Add(staticFolder)

	// GENERATE THE FOLDER TREE
	sfs := factory.NewPageArtifact(&resources.SveltinTemplatesFS, cfg.fs)
	err = projectFolder.Create(sfs)
	utils.ExitIfError(err)
	cfg.log.Success("Done\n")
}

// Command initialization.
func init() {
	newCmd.AddCommand(newPageCmd)
	pageCmdFlags(newPageCmd)
}

//=============================================================================

// Assign flags to the command.
func pageCmdFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&withSvelte, tpltypes.Svelte, "s", false, "Use Svelte for the page content")
	cmd.Flags().BoolVarP(&withMarkdown, tpltypes.Markdown, "m", false, "Use Markdown (mdsvex) for the page content")
	cmd.MarkFlagsMutuallyExclusive(tpltypes.Svelte, tpltypes.Markdown)
}

// Adding Active Help messages enhancing shell completions.
func newPageCmdValidArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var comps []string
	if len(args) == 0 {
		comps = cobra.AppendActiveHelp(comps, activehelps.Hint("You must choose a name for the page"))
	} else {
		comps = cobra.AppendActiveHelp(comps, activehelps.Hint("[WARN] This command does not take any more arguments but accepts flags"))
	}
	return comps, cobra.ShellCompDirectiveDefault
}

//=============================================================================

func makePageFolderStructure(folderName string, pageData *tpltypes.PageData) (*composer.Folder, error) {
	switch folderName {
	case RoutesFolder:
		return createPageRoutesFolder(folderName, pageData), nil
	case StaticFolder:
		return createPageStaticFolder(folderName, pageData), nil
	default:
		err := errors.New("something went wrong: folder not found as mapped resource for sveltin projects")
		return nil, sveltinerr.NewDefaultError(err)
	}
}

//=============================================================================

func createPageRoutesFolder(folderName string, pageData *tpltypes.PageData) *composer.Folder {
	// GET FOLDER: src/routes folder
	routesFolder := cfg.fsManager.GetFolder(folderName)

	// NEW FOLDER: src/routes/<page_name>
	pageFolder := composer.NewFolder(pageData.Name)
	// NEW FILE: src/routes/<page_name>/+page.svelte|svx>
	pageFile := cfg.fsManager.NewPublicPageFile(pageData, &cfg.projectSettings)

	// NEW FILE: src/routes/<page_name>/+page.server.ts>
	pageLoadFile := &composer.File{
		Name:       helpers.GetRouteFilename(IndexPageLoadFileId, cfg.settings),
		TemplateID: IndexPageLoadFileId,
		TemplateData: &config.TemplateData{
			Page: pageData,
		},
	}

	pageFolder.Add(pageFile)
	pageFolder.Add(pageLoadFile)
	routesFolder.Add(pageFolder)

	return routesFolder
}

func createPageStaticFolder(folderName string, PageData *tpltypes.PageData) *composer.Folder {
	// GET FOLDER: static
	staticFolder := cfg.fsManager.GetFolder(folderName)
	// NEW FOLDER static/pages
	allPagesFolder := composer.NewFolder("pages")
	// NEW FOLDER static/pages/<page_name>
	pageFolder := composer.NewFolder(PageData.Name)
	allPagesFolder.Add(pageFolder)
	staticFolder.Add(allPagesFolder)

	return staticFolder
}
