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
	newCmdGroupId    = "new"
	newCmdGroupTitle = "Available subcommands:"
	newCmdShortMsg   = "Create new resources and pages"
	newCmdLongMsg    = utils.MakeCmdLongMsg("Command used to creates SvelteKit routes in your project. A route in Sveltin is both a public page and a resource.")
)

//=============================================================================

var newCmd = &cobra.Command{
	Use:                   "new",
	Aliases:               []string{"n"},
	Short:                 newCmdShortMsg,
	Long:                  newCmdLongMsg,
	ValidArgs:             []string{"page", "resource"},
	ArgAliases:            []string{"p", "r"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	DisableFlagsInUseLine: true,
}

func init() {
	newCmd.AddGroup(&cobra.Group{ID: newCmdGroupId, Title: newCmdGroupTitle})
	newCmd.SetHelpCommandGroupID(newCmdGroupId)
	newCmd.SetCompletionCommandGroupID(newCmdGroupId)
	rootCmd.AddCommand(newCmd)
}
