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

var serverCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s", "serve", "run", "dev"},
	Short:   "Run the development server",
	Long: resources.GetASCIIArt() + `
It wraps vite dev to start a development server`,
	Run: RunServerCmd,
}

// RunServerCmd is the actual work function.
func RunServerCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	cfg.log.Plain(markup.H1("Running the Vite server"))

	pathToPkgFile := filepath.Join(cfg.pathMaker.GetRootFolder(), "package.json")
	npmClient, err := utils.RetrievePackageManagerFromPkgJSON(cfg.fs, pathToPkgFile)
	utils.ExitIfError(err)

	err = helpers.RunPMCommand(npmClient.Name, "dev", "", nil, false)
	utils.ExitIfError(err)
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
