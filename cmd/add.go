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
	addCmdGroupId    = "add"
	addCmdGroupTitle = "Available subcommands:"
	addCmdShortMsg   = "Add content and metadata to an existing resource"
	addCmdLongMsg    = utils.MakeCmdLongMsg(`Command used to add content and metadata to an existing resources through its own subcommands.

Run 'sveltin add -h' for further details.`)
)

//=============================================================================

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:                   "add",
	Aliases:               []string{"a"},
	Short:                 addCmdShortMsg,
	Long:                  addCmdLongMsg,
	ValidArgs:             []string{"content", "metadata"},
	ArgAliases:            []string{"c", "m"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	DisableFlagsInUseLine: true,
}

func init() {
	addCmd.AddGroup(&cobra.Group{ID: addCmdGroupId, Title: addCmdGroupTitle})
	addCmd.SetHelpCommandGroupID(addCmdGroupId)
	addCmd.SetCompletionCommandGroupID(addCmdGroupId)
	rootCmd.AddCommand(addCmd)
}
