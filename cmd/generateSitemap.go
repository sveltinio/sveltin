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

var generateSitemapCmd = &cobra.Command{
	Use:   "sitemap",
	Short: "Generate a sitemap.xml file for your Sveltin project",
	Long: resources.GetAsciiArt() + `
It creates a sitemap file for your website.

It makes use of the .env.production file to reflect the base url for your website.
`,
	Run: RunGenerateSitemapCmd,
}

// RunGenerateSitemapCmd is the actual work function.
func RunGenerateSitemapCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	log.Plain(utils.Underline("The sitemap.xml file will be created"))

	log.Info("Getting list of existing public pages")
	pages := helpers.GetAllPublicPages(AppFs, pathMaker.GetPathToPublicPages())

	log.Info("Getting list of existing resources")
	existingResources := helpers.GetAllResources(AppFs, conf.GetContentPath())

	log.Info("Getting list of all contents for the resources")
	contents := helpers.GetResourceContentMap(AppFs, existingResources, conf.GetContentPath())

	log.Info("Getting list of all metadata for the resources")
	metadata := helpers.GetResourceMetadataMap(AppFs, existingResources, conf.GetRoutesPath())

	// GET FOLDER: static
	staticFolder := fsManager.GetFolder(STATIC)

	// NEW FILE: static/rss.xml
	log.Info("Generating the sitemap.xml file")
	sitemapFile := fsManager.NewNoPage("sitemap", &projectConfig, existingResources, contents, metadata, pages)
	staticFolder.Add(sitemapFile)

	// SET FOLDER STRUCTURE
	projectFolder := fsManager.GetFolder(ROOT)
	projectFolder.Add(staticFolder)

	// GENERATE FOLDER STRUCTURE
	sfs := factory.NewNoPageArtifact(&resources.SveltinFS, AppFs)
	err := projectFolder.Create(sfs)
	utils.ExitIfError(err)

	log.Success("Done")
}

func init() {
	generateCmd.AddCommand(generateSitemapCmd)
}
