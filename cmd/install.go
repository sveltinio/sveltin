/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package cmd ...
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

var installCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"i", "init"},
	Short:   "Get all the dependencies from the `package.json` file",
	Long: resources.GetASCIIArt() + `
Initialize the Sveltin project getting all dependencies from the package.json file.

It wraps (npm|pnpm|yarn) install.
`,
	Run: RunInstallCmd,
}

// RunInstallCmd is the actual work function.
func RunInstallCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	cfg.log.Plain(markup.H1("Installing all dependencies"))

	pathToPkgFile := filepath.Join(cfg.pathMaker.GetRootFolder(), "package.json")
	npmClient, err := utils.RetrievePackageManagerFromPkgJSON(cfg.fs, pathToPkgFile)
	utils.ExitIfError(err)

	err = helpers.RunPMCommand(npmClient.Name, "install", "", nil, false)
	utils.ExitIfError(err)

	cfg.log.Success("Done")
}

func init() {
	rootCmd.AddCommand(installCmd)
}
