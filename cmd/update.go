/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/internal/npmclient"
	"github.com/sveltinio/sveltin/tui/activehelps"
	"github.com/sveltinio/sveltin/tui/feedbacks"
	"github.com/sveltinio/sveltin/utils"
)

var (
	// Short description shown in the 'help' output.
	updateCmdShortMsg = "Update your project dependencies"
	// Long message shown in the 'help <this-command>' output.
	updateCmdLongMsg = utils.MakeCmdLongMsg(`Command used to update all dependencies from the 'package.json' file.

It wraps (npm|pnpm|yarn) update.`)
)

//=============================================================================

var updateCmd = &cobra.Command{
	Use:                   "update",
	Short:                 updateCmdShortMsg,
	Long:                  updateCmdLongMsg,
	ValidArgsFunction:     updateCmdValidArgs,
	Args:                  cobra.ExactArgs(0),
	DisableFlagsInUseLine: true,
	PreRun:                allExceptInitCmdPreRunHook,
	Run:                   RunUpdateCmd,
}

// RunUpdateCmd is the actual work function.
func RunUpdateCmd(cmd *cobra.Command, args []string) {
	cfg.log.Print(markup.H1("Updating all dependencies"))

	pathToPkgFile := filepath.Join(cfg.pathMaker.GetRootFolder(), "package.json")
	npmClientInfo, err := utils.RetrievePackageManagerFromPkgJSON(cfg.fs, pathToPkgFile)
	utils.ExitIfError(err)

	err = helpers.RunNPMCommand(npmClientInfo.Name, npmclient.UpdateCmd, "", nil)
	utils.ExitIfError(err)

	cfg.log.Print(feedbacks.Success())
}

// Command initialization.
func init() {
	rootCmd.AddCommand(updateCmd)
}

//=============================================================================

// Adding Active Help messages enhancing shell completions.
func updateCmdValidArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var comps []string
	comps = cobra.AppendActiveHelp(comps, activehelps.Hint("[WARN] This command does not take any argument or flag."))
	return comps, cobra.ShellCompDirectiveDefault
}
