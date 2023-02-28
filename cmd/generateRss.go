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
	generateRssCmdShortMsg = "Generate the RSS feed for your Sveltin project"
	generateRssCmdLongMsg  = utils.MakeCmdLongMsg("Command used to generate the RSS feed (rss.xml) file for your website.")
)

//=============================================================================

var generateRssCmd = &cobra.Command{
	Use:                   "rss",
	GroupID:               generateCmdGroupId,
	Short:                 generateRssCmdShortMsg,
	Long:                  generateRssCmdLongMsg,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(0),
	Run:                   RunGenerateRSSCmd,
}

// RunGenerateRSSCmd is the actual work function.
func RunGenerateRSSCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands either from a not valid directory or not latest sveltin version.
	isValidProject(true)

	cfg.log.Plain(markup.H1("Generating the RSS feed file"))

	cfg.log.Info("Getting list of all resources contents")
	existingResources := helpers.GetAllResources(cfg.fs, cfg.pathMaker.GetPathToExistingResources())
	contents := helpers.GetResourceContentMap(cfg.fs, existingResources, cfg.settings.GetContentPath())

	cfg.log.Info("Getting list of all routes")
	allRoutes := helpers.GetAllRoutes(cfg.fs, cfg.pathMaker.GetPathToRoutes())

	// GET FOLDER: static
	staticFolder := cfg.fsManager.GetFolder(StaticFolder)

	// NEW FILE: static/rss.xml
	cfg.log.Info("Saving the file to the static folder")
	rssFile := cfg.fsManager.NewNoPageFile("rss", &cfg.projectSettings, allRoutes, contents)
	staticFolder.Add(rssFile)

	// SET FOLDER STRUCTURE
	projectFolder := cfg.fsManager.GetFolder(RootFolder)
	projectFolder.Add(staticFolder)

	// GENERATE THE FOLDER TREE
	sfs := factory.NewNoPageArtifact(&resources.SveltinTemplatesFS, cfg.fs)
	err := projectFolder.Create(sfs)
	utils.ExitIfError(err)

	cfg.log.Success("Done\n")
}

func init() {
	generateCmd.AddCommand(generateRssCmd)
}
