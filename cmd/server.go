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
	"github.com/sveltinio/sveltin/tui/activehelps"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var (
	// Short description shown in the 'help' output.
	serverCmdShortMsg = "Run the development server (vite)"
	// Long message shown in the 'help <this-command>' output.
	serverCmdLongMsg = utils.MakeCmdLongMsg("It wraps vite dev to start a development server.")
)

// Adding Active Help messages enhancing shell completions.
var serverCmdValidArgsFunc = func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var comps []string
	comps = cobra.AppendActiveHelp(comps, activehelps.Hint("[WARN] This command does not take any argument or flag."))
	return comps, cobra.ShellCompDirectiveDefault
}

//=============================================================================

var serverCmd = &cobra.Command{
	Use:                   "server",
	Aliases:               []string{"s", "serve", "run", "dev"},
	Short:                 serverCmdShortMsg,
	Long:                  serverCmdLongMsg,
	Args:                  cobra.ExactArgs(0),
	ValidArgsFunction:     serverCmdValidArgsFunc,
	DisableFlagsInUseLine: true,
	PreRun:                preRunHook,
	Run:                   RunServerCmd,
}

// RunServerCmd is the actual work function.
func RunServerCmd(cmd *cobra.Command, args []string) {
	cfg.log.Plain(markup.H1("Running the Vite server"))

	pathToPkgFile := filepath.Join(cfg.pathMaker.GetRootFolder(), "package.json")
	npmClient, err := utils.RetrievePackageManagerFromPkgJSON(cfg.fs, pathToPkgFile)
	utils.ExitIfError(err)

	err = helpers.RunPMCommand(npmClient.Name, "dev", "", nil, false)
	utils.ExitIfError(err)
}

// Command initialization.
func init() {
	rootCmd.AddCommand(serverCmd)
}
