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
It creates a sitemap file for your website.

It makes use of the .env.production file to reflect the base url for your website.
`,
	Run: RunGenerateSitemapCmd,
}

func RunGenerateSitemapCmd(cmd *cobra.Command, args []string) {
	textLogger.Reset()
	textLogger.SetTitle("A Sitemap will be created for your Sveltin project")

	listLogger.Reset()
	listLogger.AppendItem("Getting list of existing public pages")
	pages := helpers.GetAllPublicPages(AppFs, pathMaker.GetPathToPublicPages())

	listLogger.AppendItem("Getting list of existing resources")
	existingResources := helpers.GetAllResources(AppFs, conf.GetContentPath())

	listLogger.AppendItem("Getting list of all contents for the resources")
	contents := helpers.GetResourceContentMap(AppFs, existingResources, conf.GetContentPath())

	listLogger.AppendItem("Getting list of all metadata for the resources")
	metadata := helpers.GetResourceMetadataMap(AppFs, existingResources, conf.GetRoutesPath())

	// GET FOLDER: static
	staticFolder := fsManager.GetFolder(STATIC)

	// NEW FILE: static/rss.xml
	listLogger.AppendItem("Generating the sitemap.xml file")
	sitemapFile := fsManager.NewNoPage("sitemap", &projectConfig, existingResources, contents, metadata, pages)
	staticFolder.Add(sitemapFile)

	// SET FOLDER STRUCTURE
	projectFolder := fsManager.GetFolder(ROOT)
	projectFolder.Add(staticFolder)

	// GENERATE FOLDER STRUCTURE
	sfs := factory.NewNoPageArtifact(&resources.SveltinFS, AppFs)
	err := projectFolder.Create(sfs)
	utils.CheckIfError(err)

	// LOG TO STDOUT
	textLogger.SetContent(listLogger.Render())
	utils.PrettyPrinter(textLogger).Print()
}

func init() {
	generateCmd.AddCommand(generateSitemapCmd)
}
