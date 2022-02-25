/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
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
	"github.com/sveltinio/sveltin/resources"
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

// RunPreviewCmd is the actual work function.
func RunPreviewCmd(cmd *cobra.Command, args []string) {
	textLogger.Reset()
	textLogger.SetTitle("Preview your Sveltin project")

	pathToPkgFile := filepath.Join(pathMaker.GetRootFolder(), "package.json")
	npmClient, err := utils.RetrievePackageManagerFromPkgJson(AppFs, pathToPkgFile)
	utils.CheckIfError(err)

	// LOG TO STDOUT
	utils.PrettyPrinter(textLogger).Print()

	err = helpers.RunPMCommand(npmClient.Name, "preview", "", nil, false)
	utils.CheckIfError(err)
}

func init() {
	rootCmd.AddCommand(previewCmd)
}
