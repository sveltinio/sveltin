/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package cmd ...
package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var buildCmd = &cobra.Command{
	Use:     "build",
	Aliases: []string{"b"},
	Short:   "Builds a production version of your static website.",
	Long: resources.GetASCIIArt() + `
Command used to build a production version of your static website.

It wraps vite build command.

Ensure to edit env.production and .sveltin.toml files to reflect
your production environment
`,
	Run: RunBuildCmd,
}

// RunBuildCmd is the actual work function.
func RunBuildCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	cfg.log.Plain(markup.H1("Building the Sveltin project"))

	pathToPkgFile := filepath.Join(cfg.pathMaker.GetRootFolder(), "package.json")
	npmClient, err := utils.RetrievePackageManagerFromPkgJSON(cfg.fs, pathToPkgFile)
	utils.ExitIfError(err)

	os.Setenv("VITE_PUBLIC_BASE_PATH", cfg.prodData.BaseURL)
	err = helpers.RunPMCommand(npmClient.Name, "build", "", nil, false)
	utils.ExitIfError(err)

	cfg.log.Success("Done\n")
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
