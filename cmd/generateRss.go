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
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/helpers/factory"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var generateRssCmd = &cobra.Command{
	Use:   "rss",
	Short: "Generate a rss.xml file for your Sveltin project",
	Long: resources.GetASCIIArt() + `
It creates an rss file for your website.

It makes use of the .env.production file to reflect the base url for your website.
`,
	Run: RunGenerateRSSCmd,
}

// RunGenerateRSSCmd is the actual work function.
func RunGenerateRSSCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	log.Plain(utils.Underline("The rss.xml feed file will be created"))

	log.Info("Getting all existing public pages")
	pages := helpers.GetAllPublicPages(AppFs, pathMaker.GetPathToPublicPages())

	log.Info("Getting all existing resources")
	existingResources := helpers.GetAllResources(AppFs, pathMaker.GetPathToExistingResources())

	log.Info("Getting all contents for the resources")
	contents := helpers.GetResourceContentMap(AppFs, existingResources, conf.GetContentPath())

	// GET FOLDER: static
	staticFolder := fsManager.GetFolder(Static)

	// NEW FILE: static/rss.xml
	log.Info("Generating the rss.xml file")
	rssFile := fsManager.NewNoPage("rss", &projectConfig, existingResources, contents, nil, pages)
	staticFolder.Add(rssFile)

	// SET FOLDER STRUCTURE
	projectFolder := fsManager.GetFolder(Root)
	projectFolder.Add(staticFolder)

	// GENERATE FOLDER STRUCTURE
	sfs := factory.NewNoPageArtifact(&resources.SveltinFS, AppFs)
	err := projectFolder.Create(sfs)
	utils.ExitIfError(err)

	log.Success("Done")
}

func init() {
	generateCmd.AddCommand(generateRssCmd)
}
