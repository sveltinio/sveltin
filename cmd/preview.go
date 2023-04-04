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

var (
	// Short description shown in the 'help' output.
	previewCmdShortMsg = "Preview the production version locally"
	// Long message shown in the 'help <this-command>' output.
	previewCmdLongMsg = utils.MakeCmdLongMsg(`Command used to start the production version locally.

Run after sveltin build (or vite build), you can start the production version locally with sveltin preview.

It wraps vite preview command.`)
)

//=============================================================================

var previewCmd = &cobra.Command{
	Use:                   "preview",
	Short:                 previewCmdShortMsg,
	Long:                  previewCmdLongMsg,
	Args:                  cobra.ExactArgs(0),
	ValidArgsFunction:     previewCmdValidArgs,
	DisableFlagsInUseLine: true,
	PreRun:                allExceptInitCmdPreRunHook,
	Run:                   RunPreviewCmd,
}

// RunPreviewCmd is the actual work function.
func RunPreviewCmd(cmd *cobra.Command, args []string) {
	cfg.log.Print(markup.H1("Preview your Sveltin project"))

	pathToPkgFile := filepath.Join(cfg.pathMaker.GetRootFolder(), "package.json")
	npmClientInfo, err := utils.RetrievePackageManagerFromPkgJSON(cfg.fs, pathToPkgFile)
	utils.ExitIfError(err)

	err = helpers.RunNPMCommand(npmClientInfo.Name, "preview", "", nil)
	utils.ExitIfError(err)
}

// Command initialization.
func init() {
	rootCmd.AddCommand(previewCmd)
}

//=============================================================================

// Adding Active Help messages enhancing shell completions.
func previewCmdValidArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var comps []string
	comps = cobra.AppendActiveHelp(comps, activehelps.Hint("[WARN] This command does not take any argument or flag."))
	return comps, cobra.ShellCompDirectiveDefault
}
