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

var (
	// The group id under subcommands are grouped in the 'help' output.
	generateCmdGroupId = "generate"
	// The title for the group id.
	generateCmdGroupTitle = "Available subcommands:"
	// Short description shown in the 'help' output.
	generateCmdShortMsg = "Generate static files (sitemap, rss, menu)"
	// Long message shown in the 'help <this-command>' output.
	generateCmdLongMsg = utils.MakeCmdLongMsg("Command used to generate static files through its own subcommands.")
)

//=============================================================================

var generateCmd = &cobra.Command{
	Use:                   "generate",
	Aliases:               []string{"g"},
	Short:                 generateCmdShortMsg,
	Long:                  generateCmdLongMsg,
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	ValidArgs:             []string{"menu", "rss", "sitemap"},
	PersistentPreRun:      allExceptInitCmdPreRunHook,
	DisableFlagsInUseLine: true,
}

// Command initialization.
func init() {
	generateCmd.AddGroup(&cobra.Group{ID: generateCmdGroupId, Title: generateCmdGroupTitle})
	generateCmd.SetHelpCommandGroupID(generateCmdGroupId)
	generateCmd.SetCompletionCommandGroupID(generateCmdGroupId)
	rootCmd.AddCommand(generateCmd)
}
