/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
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
	"github.com/sveltinio/sveltin/internal/composer"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
)

var (
	withContentFlag bool
)

//=============================================================================

var generateMenuCmd = &cobra.Command{
	Use:   "menu",
	Short: "Generate the menu config file for your Sveltin project",
	Long: resources.GetASCIIArt() + `
It creates a 'menu.json' file into the 'config' folder to be used by Svelte components.

By default it list all resources and public pages.

The --full flag will also includes content names for all resources.
`,
	Run: RunGenerateMenuCmd,
}

// RunGenerateMenuCmd is the actual work function.
func RunGenerateMenuCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	cfg.log.Plain(markup.H1("Generating the menu structure file"))

	projectFolder := cfg.fsManager.GetFolder(RootFolder)

	cfg.log.Info("Getting list of existing public pages")
	publicPages := helpers.GetAllPublicPages(cfg.fs, cfg.pathMaker.GetPathToPublicPages())

	cfg.log.Info("Getting list of existing resources")
	availableResources := helpers.GetAllResourcesWithContentName(cfg.fs, cfg.pathMaker.GetPathToExistingResources(), withContentFlag)

	// GET FOLDER: config
	configFolder := cfg.fsManager.GetFolder(ConfigFolder)

	// ADD FILE: config/menu.js
	cfg.log.Info("Saving the menu.js.ts file")
	menuFile := &composer.File{
		Name:       "menu.js.ts",
		TemplateID: "menu",
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

	// GENERATE THE FOLDER TREE
	sfs := factory.NewMenuArtifact(&resources.SveltinFS, cfg.fs)
	err := projectFolder.Create(sfs)
	utils.ExitIfError(err)

	cfg.log.Success("Done")
}

func menuCmdFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&withContentFlag, "full", "f", false, "Generate menu file including content names for all resources.")
}

func init() {
	generateCmd.AddCommand(generateMenuCmd)
	menuCmdFlags(generateMenuCmd)
}
