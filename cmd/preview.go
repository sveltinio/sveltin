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

var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Preview the production version locally",
	Long: resources.GetAsciiArt() + `
After you've built your app with sveltin build (or svelte-kit build),
you can start the production version locally with sveltin preview.

It wraps sveltekit-preview command.`,
	Run: RunPreviewCmd,
}

func RunPreviewCmd(cmd *cobra.Command, args []string) {
	pathToPkgFile := filepath.Join(pathMaker.GetRootFolder(), "package.json")
	pkgFileContent, err := afero.ReadFile(AppFs, pathToPkgFile)
	utils.CheckIfError(err)
	pkgParsed := packagejson.Parse(pkgFileContent)
	pmInfoString := pkgParsed.PackageManager
	npmClient := utils.GetNPMClient(pmInfoString)

	// LOG TO STDOUT
	printer := utils.PrinterContent{
		Title: "Preview your Sveltin project",
	}
	printer.SetContent("")
	utils.PrettyPrinter(&printer).Print()

	err = helpers.RunPMCommand(npmClient.Name, "preview", "", nil, false)
	utils.CheckIfError(err)
}

func init() {
	rootCmd.AddCommand(previewCmd)
}
