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
	"github.com/sveltinio/sveltin/tui/activehelps"
	"github.com/sveltinio/sveltin/tui/feedbacks"
	"github.com/sveltinio/sveltin/utils"
)

var (
	// Short description shown in the 'help' output.
	generateSitemapCmdShortMsg = "Generate the sitemap file for your Sveltin project"
	// Long message shown in the 'help <this-command>' output.
	generateSitemapCmdLongMsg = utils.MakeCmdLongMsg("Command used to generate the sitemap (sitemap.xml) file for your website.")
)

//=============================================================================

var generateSitemapCmd = &cobra.Command{
	Use:                   "sitemap",
	GroupID:               generateCmdGroupId,
	Short:                 generateSitemapCmdShortMsg,
	Long:                  generateSitemapCmdLongMsg,
	Args:                  cobra.ExactArgs(0),
	ValidArgsFunction:     generateSitemapCmdValidArgs,
	DisableFlagsInUseLine: true,
	Run:                   RunGenerateSitemapCmd,
}

// RunGenerateSitemapCmd is the actual work function.
func RunGenerateSitemapCmd(cmd *cobra.Command, args []string) {
	cfg.log.Print(markup.H1("Generating the sitemap file"))

	cfg.log.Info("Getting list of all resources contents")
	existingResources := helpers.GetAllResources(cfg.fs, cfg.settings.GetContentPath())
	contents := helpers.GetResourceContentMap(cfg.fs, existingResources, cfg.settings.GetContentPath())

	cfg.log.Info("Getting list of all routes")
	allRoutes := helpers.GetAllRoutes(cfg.fs, cfg.pathMaker.GetPathToRoutes())

	// GET FOLDER: static
	staticFolder := cfg.fsManager.GetFolder(StaticFolder)

	// NEW FILE: static/rss.xml
	cfg.log.Info("Saving the file to the static folder")
	sitemapFile := cfg.fsManager.NewNoPageFile("sitemap", &cfg.projectSettings, allRoutes, contents)
	staticFolder.Add(sitemapFile)

	// SET FOLDER STRUCTURE
	projectFolder := cfg.fsManager.GetFolder(RootFolder)
	projectFolder.Add(staticFolder)

	// GENERATE THE FOLDER TREE
	sfs := factory.NewNoPageArtifact(&resources.SveltinTemplatesFS, cfg.fs)
	err := projectFolder.Create(sfs)
	utils.ExitIfError(err)

	cfg.log.Print(feedbacks.Success())
}

// Command initialization.
func init() {
	generateCmd.AddCommand(generateSitemapCmd)
}

//=============================================================================

// Adding Active Help messages enhancing shell completions.
func generateSitemapCmdValidArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var comps []string
	comps = cobra.AppendActiveHelp(comps, activehelps.Hint("[WARN] This command does not take any argument or flag."))
	return comps, cobra.ShellCompDirectiveDefault
}
