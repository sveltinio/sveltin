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
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
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
	Short: "Generate the menu file for your Sveltin project.",
	Long: resources.GetASCIIArt() + `
Command used to generate the menu (menu.js.ts) file into the 'config' folder to be used by Svelte components.

By default it list all resources and public pages.

The --full flag will includes content names for all resources too.
`,
	Run: RunGenerateMenuCmd,
}

// RunGenerateMenuCmd is the actual work function.
func RunGenerateMenuCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	cfg.log.Plain(markup.H1("Generating the menu structure file"))

	cfg.log.Info("Getting list of all resources contents")
	existingResources := helpers.GetAllResources(cfg.fs, cfg.sveltin.GetContentPath())
	contents := helpers.GetResourceContentMap(cfg.fs, existingResources, cfg.sveltin.GetContentPath())

	cfg.log.Info("Getting list of all routes")
	allRoutes := helpers.GetAllRoutes(cfg.fs, cfg.pathMaker.GetPathToRoutes())

	// GET FOLDER: config
	configFolder := cfg.fsManager.GetFolder(ConfigFolder)

	// ADD FILE: config/menu.js
	cfg.log.Info("Saving the menu.js.ts file")
	menuFile := cfg.fsManager.NewMenuFile("menu", &cfg.project, allRoutes, contents, withContentFlag)
	configFolder.Add(menuFile)

	// SET FOLDER STRUCTURE
	projectFolder := cfg.fsManager.GetFolder(RootFolder)
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
