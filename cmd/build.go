/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
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
	Long: resources.GetAsciiArt() + `
Builds a production version of your static website.

It wraps sveltekit-build command.

Ensure to edit env.production and .sveltin.toml files to reflect
your production environment
`,
	Run: RunBuildCmd,
}

func RunBuildCmd(cmd *cobra.Command, args []string) {
	textLogger.Reset()
	textLogger.SetTitle("Building Sveltin project")

	pathToPkgFile := filepath.Join(pathMaker.GetRootFolder(), "package.json")
	npmClient, err := utils.RetrievePackageManagerFromPkgJson(AppFs, pathToPkgFile)
	utils.CheckIfError(err)

	os.Setenv("VITE_PUBLIC_BASE_PATH", projectConfig.BaseURL)
	err = helpers.RunPMCommand(npmClient.Name, "build", "", nil, false)
	utils.CheckIfError(err)

	// LOG TO STDOUT
	utils.PrettyPrinter(textLogger).Print()
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
