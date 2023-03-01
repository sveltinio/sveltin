/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var (
	// How to use the command.
	generateMenuCmdExample = "sveltin generate menu --full"
	// Short description shown in the 'help' output.
	generateMenuCmdShortMsg = "Generate the menu file for your Sveltin project"
	// Long message shown in the 'help <this-command>' output.
	generateMenuCmdLongMsg = utils.MakeCmdLongMsg(`Command used to generate the menu (menu.js.ts) file into the 'config' folder to be used by Svelte components.

By default it list all resources and public pages.

The --full flag will includes content names for all resources too.`)
)

// Bind command flags.
var (
	withContentFlag bool
)

//=============================================================================

var generateMenuCmd = &cobra.Command{
	Use:     "menu",
	GroupID: generateCmdGroupId,
	Example: generateMenuCmdExample,
	Short:   generateMenuCmdShortMsg,
	Long:    generateMenuCmdLongMsg,
	Args:    cobra.ExactArgs(0),
	Run:     RunGenerateMenuCmd,
}

// RunGenerateMenuCmd is the actual work function.
func RunGenerateMenuCmd(cmd *cobra.Command, args []string) {
	cfg.log.Plain(markup.H1("Generating the menu structure file"))

	cfg.log.Info("Getting list of all resources contents")
	existingResources := helpers.GetAllResources(cfg.fs, cfg.settings.GetContentPath())
	contents := helpers.GetResourceContentMap(cfg.fs, existingResources, cfg.settings.GetContentPath())

	cfg.log.Info("Getting list of all routes")
	allRoutes := helpers.GetAllRoutes(cfg.fs, cfg.pathMaker.GetPathToRoutes())

	// GET FOLDER: config
	configFolder := cfg.fsManager.GetFolder(ConfigFolder)

	// ADD FILE: config/menu.js
	cfg.log.Info("Saving the menu.js.ts file")
	menuFile := cfg.fsManager.NewMenuFile("menu", allRoutes, contents, withContentFlag)
	configFolder.Add(menuFile)

	// SET FOLDER STRUCTURE
	projectFolder := cfg.fsManager.GetFolder(RootFolder)
	projectFolder.Add(configFolder)

	// GENERATE THE FOLDER TREE
	sfs := factory.NewMenuArtifact(&resources.SveltinTemplatesFS, cfg.fs)
	err := projectFolder.Create(sfs)
	utils.ExitIfError(err)

	cfg.log.Success("Done\n")
}

// Assign flags to the command.
func menuCmdFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&withContentFlag, "full", "f", false, "Generate menu file including content names for all resources")
}

// Command initialization.
func init() {
	generateCmd.AddCommand(generateMenuCmd)
	menuCmdFlags(generateMenuCmd)
}
