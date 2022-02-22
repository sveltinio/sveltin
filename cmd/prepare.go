/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var prepareCmd = &cobra.Command{
	Use:     "prepare",
	Aliases: []string{"i", "install", "init"},
	Short:   "Get all the dependencies from the `package.json` file",
	Long: resources.GetAsciiArt() + `
Initialize the Sveltin project getting all dependencies from the package.json file.

It wraps (npm|pnpm|yarn) install.
`,
	Run: RunPrepareCmd,
}

func RunPrepareCmd(cmd *cobra.Command, args []string) {
	textLogger.Reset()
	textLogger.SetTitle("Prepare Sveltin project")
	textLogger.SetContent("* Getting dependencies")

	pathToPkgFile := filepath.Join(pathMaker.GetRootFolder(), "package.json")
	npmClient, err := utils.RetrievePackageManagerFromPkgJson(AppFs, pathToPkgFile)
	utils.CheckIfError(err)

	// LOG TO STDOUT
	utils.PrettyPrinter(textLogger).Print()

	err = helpers.RunPMCommand(npmClient.Name, "install", "", nil, false)
	utils.CheckIfError(err)
}

func init() {
	rootCmd.AddCommand(prepareCmd)
}
