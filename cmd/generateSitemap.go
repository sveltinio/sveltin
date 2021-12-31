/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
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
It creates a sitemap file for your Sveltin project.`,
	Run: RunGenerateSitemapCmd,
}

func RunGenerateSitemapCmd(cmd *cobra.Command, args []string) {
	logger.Reset()

	printer := utils.PrinterContent{
		Title: "A Sitemap will be created for your Sveltin project",
	}

	logger.AppendItem("Getting list of existing public pages")
	pages := helpers.GetAllPublicPages(AppFs, pathMaker.GetPathToPublicPages())

	logger.AppendItem("Getting list of existing resources")
	existingResources := helpers.GetAllResources(AppFs, conf.GetContentPath())

	logger.AppendItem("Getting list of all contents for the resources")
	contents := helpers.GetResourceContentMap(AppFs, existingResources, conf.GetContentPath())

	logger.AppendItem("Getting list of all metadata for the resources")
	metadata := helpers.GetResourceMetadataMap(AppFs, existingResources, conf.GetRoutesPath())

	// GET FOLDER: static
	staticFolder := fsManager.GetFolder(STATIC)

	// NEW FILE: static/rss.xml
	logger.AppendItem("Generating the sitemap.xml file")
	sitemapFile := fsManager.NewNoPage("sitemap", &siteConfig, existingResources, contents, metadata, pages)
	staticFolder.Add(sitemapFile)

	// SET FOLDER STRUCTURE
	projectFolder := fsManager.GetFolder(ROOT)
	projectFolder.Add(staticFolder)

	// GENERATE FOLDER STRUCTURE
	sfs := factory.NewNoPageArtifact(&resources.SveltinFS, AppFs)
	err := projectFolder.Create(sfs)
	utils.CheckIfError(err)

	// LOG TO STDOUT
	printer.SetContent(logger.Render())
	utils.PrettyPrinter(&printer).Print()
}

func init() {
	generateCmd.AddCommand(generateSitemapCmd)
}
