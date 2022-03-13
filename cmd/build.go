/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
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
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var buildCmd = &cobra.Command{
	Use:     "build",
	Aliases: []string{"b"},
	Short:   "Builds a production version of your static website",
	Long: resources.GetASCIIArt() + `
Builds a production version of your static website.

It wraps sveltekit-build command.

Ensure to edit env.production and .sveltin.toml files to reflect
your production environment
`,
	Run: RunBuildCmd,
}

// RunBuildCmd is the actual work function.
func RunBuildCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	log.Plain(utils.Underline("Building the Sveltin project"))

	pathToPkgFile := filepath.Join(pathMaker.GetRootFolder(), "package.json")
	npmClient, err := utils.RetrievePackageManagerFromPkgJSON(AppFs, pathToPkgFile)
	utils.ExitIfError(err)

	os.Setenv("VITE_PUBLIC_BASE_PATH", projectConfig.BaseURL)
	err = helpers.RunPMCommand(npmClient.Name, "build", "", nil, false)
	utils.ExitIfError(err)
	log.Success("Done")
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
