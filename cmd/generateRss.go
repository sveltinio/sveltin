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

var generateRssCmd = &cobra.Command{
	Use:   "rss",
	Short: "Generate a rss.xml file for your Sveltin project",
	Long: resources.GetAsciiArt() + `
Do you wish to have an RSS feed for your content?

	The "generate rss" command generates it for you.`,
	Run: RunGenerateRSSCmd,
}

func RunGenerateRSSCmd(cmd *cobra.Command, args []string) {
	logger.Reset()

	printer := utils.PrinterContent{
		Title: "An RSS feed file will be created for your Sveltin project",
	}

	logger.AppendItem("Getting all existing public pages")
	pages := helpers.GetAllPublicPages(AppFs, pathMaker.GetPathToPublicPages())

	logger.AppendItem("Getting all existing resources")
	existingResources := helpers.GetAllResources(AppFs, pathMaker.GetPathToExistingResources())

	logger.AppendItem("Getting all contents for the resources")
	contents := helpers.GetResourceContentMap(AppFs, existingResources, conf.GetContentPath())

	// GET FOLDER: static
	staticFolder := fsManager.GetFolder(STATIC)

	// NEW FILE: static/rss.xml
	logger.AppendItem("Generating the rss.xml file")
	rssFile := fsManager.NewNoPage("rss", &siteConfig, existingResources, contents, nil, pages)
	staticFolder.Add(rssFile)

	// SET FOLDER STRUCTURE
	projectFolder := fsManager.GetFolder(ROOT)
	projectFolder.Add(staticFolder)

	// GENERATE FOLDER STRUCTURE
	sfs := factory.NewNoPageArtifact(&resources.SveltinFS, AppFs)
	err := projectFolder.Create(sfs)
	utils.CheckIfError(err)

	// LOG TO STDOUT
	// LOG TO STDOUT
	printer.SetContent(logger.Render())
	utils.PrettyPrinter(&printer).Print()
}

func init() {
	generateCmd.AddCommand(generateRssCmd)
}
