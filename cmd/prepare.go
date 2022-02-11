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
Initialize the Sveltin project getting all depencencies from the package.json file.

It wraps (npm|pnpm|yarn) install.
`,
	Run: RunPrepareCmd,
}

func RunPrepareCmd(cmd *cobra.Command, args []string) {
	pathToPkgFile := filepath.Join(pathMaker.GetRootFolder(), "package.json")
	npmClient := utils.RetrievePackageManagerFromPkgJson(AppFs, pathToPkgFile)

	// LOG TO STDOUT
	printer := utils.PrinterContent{
		Title: "Prepare Sveltin project",
	}
	printer.SetContent("* Getting dependencies")
	utils.PrettyPrinter(&printer).Print()

	err := helpers.RunPMCommand(npmClient.Name, "install", "", nil, false)
	utils.CheckIfError(err)
}

func init() {
	rootCmd.AddCommand(prepareCmd)
}
