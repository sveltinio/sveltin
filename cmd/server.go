/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package cmd

import (
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/helpers"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/sveltinlib/packagejson"
	"github.com/sveltinio/sveltin/utils"
)

//=============================================================================

var serverCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s", "serve"},
	Short:   "Run the server",
	Long: resources.GetAsciiArt() + `
It wraps svelte-kit defined commands to run the server`,
	Run: RunServerCmd,
}

func RunServerCmd(cmd *cobra.Command, args []string) {
	pathToPkgFile := filepath.Join(pathMaker.GetRootFolder(), "package.json")
	pkgFileContent, err := afero.ReadFile(AppFs, pathToPkgFile)
	utils.CheckIfError(err)
	pkgParsed := packagejson.Parse(pkgFileContent)
	pmInfoString := pkgParsed.PackageManager
	npmClient := utils.GetNPMClient(pmInfoString)

	// LOG TO STDOUT
	printer := utils.PrinterContent{
		Title: "Running Vite server",
	}
	printer.SetContent("")
	utils.PrettyPrinter(&printer).Print()

	err = helpers.RunPMCommand(npmClient.Name, "dev", "", nil, false)
	utils.CheckIfError(err)
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
