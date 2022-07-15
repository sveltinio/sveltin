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

	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/pkg/sveltinerr"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var (
	pageType string
)

const (
	// Svelte set svelte as the language used to scaffold a new page
	Svelte string = "svelte"
	// Markdown set markdown as the language used to scaffold a new page
	Markdown string = "markdown"
)

//=============================================================================

var newPageCmd = &cobra.Command{
	Use:     "page [name]",
	Aliases: []string{"p"},
	Short:   "Command to create a new public page",
	Long: resources.GetASCIIArt() + `
Create a new "public" page.

Pages are Svelte components written in .svelte files. The filename determines the route.
A file called either src/routes/about.svelte or src/routes/about/index.svelte
would correspond to the /about route.

This command allows you to select between a svelte component page and a markdown page.`,
	Run: NewPageCmdRun,
}

// NewPageCmdRun is the actual work function.
func NewPageCmdRun(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	pageName, err := promptPageName(args)
	utils.ExitIfError(err)

	pageType, err := promptPageType(pageType)
	utils.ExitIfError(err)

	cfg.log.Plain(utils.Underline(fmt.Sprintf("Creating '%s' as %s page", pageName, pageType)))

	// GET FOLDER: src/routes
	routesFolder := cfg.fsManager.GetFolder(RoutesFolder)

	// NEW FILE: src/routes/<page_name.svelte|svx>
	pageFile := cfg.fsManager.NewPublicPage(pageName, pageType)
	utils.ExitIfError(err)

	// ADD TO THE ROUTES FOLDER
	routesFolder.Add(pageFile)

	// SET FOLDER STRUCTURE
	projectFolder := cfg.fsManager.GetFolder(RootFolder)
	projectFolder.Add(routesFolder)

	// GENERATE THE FOLDER TREE
	sfs := factory.NewPageArtifact(&resources.SveltinFS, cfg.fs)
	err = projectFolder.Create(sfs)
	utils.ExitIfError(err)
	cfg.log.Success("Done")
}

func pageCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&pageType, "type", "t", "", "Sveltekit page as Svelte component or markdown via mdsvex. Possible values: svelte, markdown")
}

func init() {
	newCmd.AddCommand(newPageCmd)
	pageCmdFlags(newPageCmd)
}

//=============================================================================

func promptPageName(inputs []string) (string, error) {
	var name string
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		pageNamePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide a name for the page.",
			Label:    "What's the page name?",
		}
		name, err := common.PromptGetInput(pageNamePromptContent, nil, "")
		if err != nil {
			return "", err
		}
		return utils.ToSlug(name), nil
	case numOfArgs == 1:
		name = inputs[0]
		return utils.ToSlug(name), nil
	default:
		err := errors.New("something went wrong: value not valid")
		return "", sveltinerr.NewDefaultError(err)
	}
}

func promptPageType(pageTypeFlag string) (string, error) {
	promptObjects := []config.PromptObject{
		{ID: Svelte, Name: "Svelte"},
		{ID: Markdown, Name: "Markdown in Svelte"},
	}

	switch nameLenght := len(pageTypeFlag); {
	case nameLenght == 0:
		pagePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide a page type",
			Label:    "What's the page type?",
		}
		result, err := common.PromptGetSelect(pagePromptContent, promptObjects, true)
		if err != nil {
			return "", err
		}
		return result, nil
	case nameLenght != 0:
		valid := common.GetPromptObjectKeys(promptObjects)
		if !common.Contains(valid, pageTypeFlag) {
			return "", sveltinerr.NewPageTypeNotValidError()
		}
		return pageTypeFlag, nil
	default:
		return "", sveltinerr.NewPageTypeNotValidError()
	}
}
