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
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var generateRssCmd = &cobra.Command{
	Use:   "rss",
	Short: "Generate the RSS feed for your Sveltin project.",
	Long: resources.GetASCIIArt() + `
Command used to generate the RSS feed (rss.xml) file for your website.

It makes use of the .env.production file to reflect the base url for your website.
`,
	Run: RunGenerateRSSCmd,
}

// RunGenerateRSSCmd is the actual work function.
func RunGenerateRSSCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	cfg.log.Plain(markup.H1("Generating the RSS feed file"))

	cfg.log.Info("Getting all existing resources")
	existingResources := helpers.GetAllResources(cfg.fs, cfg.pathMaker.GetPathToExistingResources())

	cfg.log.Info("Getting all resources contents")
	contents := helpers.GetResourceContentMap(cfg.fs, existingResources, cfg.sveltin.GetContentPath())

	cfg.log.Info("Getting all existing public pages")
	allRoutes := helpers.GetAllRoutes(cfg.fs, cfg.pathMaker.GetPathToRoutes())
	allRoutesExceptsResource := common.Difference(allRoutes, existingResources)
	// exclude api folder from the list
	pages := common.Difference(allRoutesExceptsResource, []string{ApiFolder})
	//pages := helpers.GetAllPublicPages(cfg.fs, cfg.pathMaker.GetPathToPublicPages())

	// GET FOLDER: static
	staticFolder := cfg.fsManager.GetFolder(StaticFolder)

	// NEW FILE: static/rss.xml
	cfg.log.Info("Saving the file to the static folder")
	rssFile := cfg.fsManager.NewNoPage("rss", &cfg.project, existingResources, contents, nil, pages)
	staticFolder.Add(rssFile)

	// SET FOLDER STRUCTURE
	projectFolder := cfg.fsManager.GetFolder(RootFolder)
	projectFolder.Add(staticFolder)

	// GENERATE THE FOLDER TREE
	sfs := factory.NewNoPageArtifact(&resources.SveltinFS, cfg.fs)
	err := projectFolder.Create(sfs)
	utils.ExitIfError(err)

	cfg.log.Success("Done")
}

func init() {
	generateCmd.AddCommand(generateRssCmd)
}
