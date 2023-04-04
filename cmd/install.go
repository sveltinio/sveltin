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
	installCmdShortMsg = "Install your project dependencies"
	// Long message shown in the 'help <this-command>' output.
	installCmdLongMsg = utils.MakeCmdLongMsg(`Command used to install all dependencies from the 'package.json' file.

It wraps (npm|pnpm|yarn) install.`)
)

//=============================================================================

var installCmd = &cobra.Command{
	Use:                   "install",
	Aliases:               []string{"i"},
	Short:                 installCmdShortMsg,
	Long:                  installCmdLongMsg,
	Args:                  cobra.ExactArgs(0),
	ValidArgsFunction:     installCmdValidArgs,
	DisableFlagsInUseLine: true,
	PreRun:                allExceptInitCmdPreRunHook,
	Run:                   RunInstallCmd,
}

// RunInstallCmd is the actual work function.
func RunInstallCmd(cmd *cobra.Command, args []string) {
	cfg.log.Print(markup.H1("Installing all dependencies"))

	pathToPkgFile := filepath.Join(cfg.pathMaker.GetRootFolder(), "package.json")
	npmClientInfo, err := utils.RetrievePackageManagerFromPkgJSON(cfg.fs, pathToPkgFile)
	utils.ExitIfError(err)

	err = helpers.RunNPMCommand(npmClientInfo.Name, npmclient.InstallCmd, "", nil)
	utils.ExitIfError(err)

	cfg.log.Print(feedbacks.Success())
}

// Command initialization.
func init() {
	rootCmd.AddCommand(installCmd)
}

//=============================================================================

// Adding Active Help messages enhancing shell completions.
func installCmdValidArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var comps []string
	comps = cobra.AppendActiveHelp(comps, activehelps.Hint("[WARN] This command does not take any argument or flag."))
	return comps, cobra.ShellCompDirectiveDefault
}
