/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package cmd ...
package cmd

import (
	"github.com/spf13/cobra"
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
It creates a 'menu.json' file into the 'config' folder to be used by Svelte components.

By default it list all resources and public pages.

The --full flag will also includes content names for all resources.
`,
	Run: RunGenerateMenuCmd,
}

// RunGenerateMenuCmd is the actual work function.
func RunGenerateMenuCmd(cmd *cobra.Command, args []string) {
	textLogger.Reset()
	textLogger.SetTitle("The menu structure for your Sveltin project will be created")

	projectFolder := fsManager.GetFolder(ROOT)

	listLogger.Reset()
	listLogger.AppendItem("Getting list of existing public pages")
	publicPages := helpers.GetAllPublicPages(AppFs, pathMaker.GetPathToPublicPages())

	listLogger.AppendItem("Getting list of existing resources")
	availableResources := helpers.GetAllResourcesWithContentName(AppFs, pathMaker.GetPathToExistingResources(), withContentFlag)

	// GET FOLDER: config
	configFolder := fsManager.GetFolder(CONFIG)

	// ADD FILE: config/menu.js
	listLogger.AppendItem("Generating menu.js.ts file")
	menuFile := &composer.File{
		Name:       "menu.js.ts",
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
	utils.ExitIfError(err)

	// LOG TO STDOUT
	textLogger.SetContent(listLogger.Render())
	utils.PrettyPrinter(textLogger).Print()
}

func menuCmdFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&withContentFlag, "full", "f", false, "Generate menu file including content names for all resources.")
}

func init() {
	generateCmd.AddCommand(generateMenuCmd)
	menuCmdFlags(generateMenuCmd)
}
