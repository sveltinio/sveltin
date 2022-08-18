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

//=============================================================================

var generateSitemapCmd = &cobra.Command{
	Use:   "sitemap",
	Short: "Generate a sitemap.xml file for your Sveltin project",
	Long: resources.GetASCIIArt() + `
It creates a sitemap file for your website.

It makes use of the .env.production file to reflect the base url for your website.
`,
	Run: RunGenerateSitemapCmd,
}

// RunGenerateSitemapCmd is the actual work function.
func RunGenerateSitemapCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	cfg.log.Plain(markup.H1("Generating the sitemap file"))

	cfg.log.Info("Getting list of existing public pages")
	pages := helpers.GetAllPublicPages(cfg.fs, cfg.pathMaker.GetPathToPublicPages())

	cfg.log.Info("Getting list of existing resources")
	existingResources := helpers.GetAllResources(cfg.fs, cfg.sveltin.GetContentPath())

	cfg.log.Info("Getting list of all resources contents")
	contents := helpers.GetResourceContentMap(cfg.fs, existingResources, cfg.sveltin.GetContentPath())

	cfg.log.Info("Getting list of all resources metadata")
	metadata := helpers.GetResourceMetadataMap(cfg.fs, existingResources, cfg.sveltin.GetRoutesPath())

	// GET FOLDER: static
	staticFolder := cfg.fsManager.GetFolder(StaticFolder)

	// NEW FILE: static/rss.xml
	cfg.log.Info("Saving the file to the static folder")
	sitemapFile := cfg.fsManager.NewNoPage("sitemap", &cfg.project, existingResources, contents, metadata, pages)
	staticFolder.Add(sitemapFile)

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
	generateCmd.AddCommand(generateSitemapCmd)
}
