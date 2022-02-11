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

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"u"},
	Short:   "Update the dependencies from the `package.json` file",
	Long: resources.GetAsciiArt() + `
Update all depencencies from the package.json file.

It wraps (npm|pnpm|yarn) update.
`,
	Run: RunUpdateCmd,
}

func RunUpdateCmd(cmd *cobra.Command, args []string) {
	pathToPkgFile := filepath.Join(pathMaker.GetRootFolder(), "package.json")
	npmClient := utils.RetrievePackageManagerFromPkgJson(AppFs, pathToPkgFile)

	// LOG TO STDOUT
	printer := utils.PrinterContent{
		Title: "Update dependencies Sveltin project",
	}
	printer.SetContent("* Updating dependencies")
	utils.PrettyPrinter(&printer).Print()

	err := helpers.RunPMCommand(npmClient.Name, "update", "", nil, false)
	utils.CheckIfError(err)
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
