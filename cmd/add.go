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
	// The group id under subcommands are grouped in the 'help' output.
	addCmdGroupId = "add"
	// The title for the group id.
	addCmdGroupTitle = "Available subcommands:"
	// Short description shown in the 'help' output.
	addCmdShortMsg = "Add content and metadata to an existing resource"
	// Long message shown in the 'help <this-command>' output.
	addCmdLongMsg = utils.MakeCmdLongMsg("Command used to add content and metadata to an existing resources through its own subcommands.")
)

//=============================================================================

var addCmd = &cobra.Command{
	Use:                   "add",
	Aliases:               []string{"a"},
	Short:                 addCmdShortMsg,
	Long:                  addCmdLongMsg,
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	ArgAliases:            []string{"c", "m"},
	ValidArgs:             []string{"content", "metadata"},
	DisableFlagsInUseLine: true,
	PersistentPreRun:      preRunHook,
}

// Command initialization.
func init() {
	addCmd.AddGroup(&cobra.Group{ID: addCmdGroupId, Title: addCmdGroupTitle})
	addCmd.SetHelpCommandGroupID(addCmdGroupId)
	addCmd.SetCompletionCommandGroupID(addCmdGroupId)
	rootCmd.AddCommand(addCmd)
}
