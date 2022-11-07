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
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"u"},
	Short:   "Update your project dependencies.",
	Long: resources.GetASCIIArt() + `
Command used to update all dependencies from the 'package.json' file.

It wraps (npm|pnpm|yarn) update.
`,
	Run: RunUpdateCmd,
}

// RunUpdateCmd is the actual work function.
func RunUpdateCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands either from a not valid directory or not latest sveltin version.
	isValidProject(true)

	cfg.log.Plain(markup.H1("Updating all dependencies"))

	pathToPkgFile := filepath.Join(cfg.pathMaker.GetRootFolder(), "package.json")
	npmClient, err := utils.RetrievePackageManagerFromPkgJSON(cfg.fs, pathToPkgFile)
	utils.ExitIfError(err)

	err = helpers.RunPMCommand(npmClient.Name, "update", "", nil, false)
	utils.ExitIfError(err)

	cfg.log.Success("Done\n")
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
