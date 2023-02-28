/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var (
	generateCmdGroupId    = "generate"
	generateCmdGroupTitle = "Available subcommands:"
	generateCmdShortMsg   = "Generate static files (sitemap, rss, menu)"
	generateCmdLongMsg    = utils.MakeCmdLongMsg("Command used to generate static files through its own subcommands.")
)

//=============================================================================

var generateCmd = &cobra.Command{
	Use:                   "generate",
	Aliases:               []string{"g"},
	Short:                 generateCmdShortMsg,
	Long:                  generateCmdLongMsg,
	ValidArgs:             []string{"menu", "rss", "sitemap"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	DisableFlagsInUseLine: true,
}

func init() {
	generateCmd.AddGroup(&cobra.Group{ID: generateCmdGroupId, Title: generateCmdGroupTitle})
	generateCmd.SetHelpCommandGroupID(generateCmdGroupId)
	generateCmd.SetCompletionCommandGroupID(generateCmdGroupId)
	rootCmd.AddCommand(generateCmd)
}
