/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/sveltinlib/composer"
	"github.com/sveltinio/sveltin/utils"
)

var (
	withContentFlag bool
)

//=============================================================================

var generateMenuCmd = &cobra.Command{
	Use:   "menu",
	Short: "Generate the menu config file for your Sveltin project",
	Long: resources.GetAsciiArt() + `
It creates a 'menu.json' file into the 'config' folder to be used by Svelte components.`,
	Run: RunGenerateMenuCmd,
}

func RunGenerateMenuCmd(cmd *cobra.Command, args []string) {
	logger.Reset()

	printer := utils.PrinterContent{
		Title: "A 'menu.js' file will be created for your Sveltin project",
	}

	projectFolder := fsManager.GetFolder(ROOT)

	logger.AppendItem("Getting list of existing public pages")
	publicPages := helpers.GetAllPublicPages(AppFs, pathMaker.GetPathToPublicPages())

	logger.AppendItem("Getting list of existing resources")
	availableResources := helpers.GetAllResourcesWithContentName(AppFs, pathMaker.GetPathToExistingResources(), withContentFlag)

	// GET FOLDER: config
	configFolder := fsManager.GetFolder(CONFIG)

	// ADD FILE: config/menu.js
	logger.AppendItem("Generating menu.js file")
	menuFile := &composer.File{
		Name:       "menu.js",
		TemplateId: "menu",
		TemplateData: &config.TemplateData{
			Menu: &config.MenuConfig{
				Resources:   availableResources,
				Pages:       publicPages,
				WithContent: withContentFlag,
			},
		},
	}
	configFolder.Add(menuFile)

	// SET FOLDER STRUCTURE
	projectFolder.Add(configFolder)

	// GENERATE FOLDER STRUCTURE
	sfs := factory.NewMenuArtifact(&resources.SveltinFS, AppFs)
	err := projectFolder.Create(sfs)
	common.CheckIfError(err)

	// LOG TO STDOUT
	printer.SetContent(logger.Render())
	utils.PrettyPrinter(&printer).Print()
}

func menuCmdFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&withContentFlag, "full", "f", false, "Generate menu file including content names for all resources.")
}

func init() {
	generateCmd.AddCommand(generateMenuCmd)
	menuCmdFlags(generateMenuCmd)
}
