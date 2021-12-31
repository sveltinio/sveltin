/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var (
	pageType string
)

const (
	SVELTE   string = "svelte"
	MARKDOWN string = "markdown"
)

//=============================================================================

var newPageCmd = &cobra.Command{
	Use:   "page [name]",
	Short: "Create a new public page",
	Long: resources.GetAsciiArt() + `
Create a new "public" page.

Pages are Svelte components written in .svelte files. The filename determines the route.
A file called either src/routes/about.svelte or src/routes/about/index.svelte
would correspond to the /about route.

This command allows you to select between a svelte component page and a markdown page.`,
	Run: NewPageCmdRun,
}

func NewPageCmdRun(cmd *cobra.Command, args []string) {
	logger.Reset()

	pageName, err := getPageName(args)
	utils.CheckIfError(err)

	printer := utils.PrinterContent{
		Title: `New "` + pageName + `" page added to your Sveltin project`,
	}

	pageType, err := getPageType(pageType)
	utils.CheckIfError(err)

	// GET FOLDER: src/routes
	routesFolder := fsManager.GetFolder(ROUTES)

	// NEW FILE: src/routes/<page_name.svelte|svx>
	pageFile := fsManager.NewPublicPage(pageName, pageType)
	utils.CheckIfError(err)

	// ADD TO THE ROUTES FOLDER
	routesFolder.Add(pageFile)

	// SET FOLDER STRUCTURE
	projectFolder := fsManager.GetFolder(ROOT)
	projectFolder.Add(routesFolder)

	// GENERATE STRUCTURE
	sfs := factory.NewPageArtifact(&resources.SveltinFS, AppFs)
	err = projectFolder.Create(sfs)
	utils.CheckIfError(err)

	// LOG TO STDOUT
	printer.SetContent(logger.Render())
	utils.PrettyPrinter(&printer).Print()
}

func pageCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&pageType, "type", "t", "", "Sveltekit page as Svelte component or markdown via mdsvex. Possible values: svelte, markdown")
}

func init() {
	newCmd.AddCommand(newPageCmd)
	pageCmdFlags(newPageCmd)
}

//=============================================================================

func getPageName(inputs []string) (string, error) {
	var name string
	switch numOfArgs := len(inputs); {
	case numOfArgs < 1:
		pageNamePromptContent := config.PromptContent{
			ErrorMsg: "Please, provide a name for your page.",
			Label:    "What's the name of your page?",
		}
		name = common.PromptGetInput(pageNamePromptContent)
		return name, nil
	case numOfArgs == 1:
		name = inputs[0]
		return name, nil
	default:
		err := errors.New("something went wrong: value not valid")
		return "", sveltinerr.NewDefaultError(err)
	}
}

func getPageType(pageTypeFlag string) (string, error) {
	valid := []string{SVELTE, MARKDOWN}
	var page string
	switch nameLenght := len(pageTypeFlag); {
	case nameLenght == 0:
		pagePromptContent := config.PromptContent{
			ErrorMsg: "Please, select a type for your page",
			Label:    "What's the page type you prefer?",
		}
		page = common.PromptGetSelect(valid, pagePromptContent)
		return page, nil
	case nameLenght != 0:
		page = pageTypeFlag
		if !common.Contains(valid, page) {
			return "", sveltinerr.NewPageTypeNotValidError()
		} else {
			return page, nil
		}
	default:
		return "", sveltinerr.NewPageTypeNotValidError()
	}
}
