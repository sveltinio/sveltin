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

var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Preview the production version locally",
	Long: resources.GetASCIIArt() + `
After you've built your app with sveltin build (or svelte-kit build),
you can start the production version locally with sveltin preview.

It wraps sveltekit-preview command.`,
	Run: RunPreviewCmd,
}

// RunPreviewCmd is the actual work function.
func RunPreviewCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	cfg.log.Plain(markup.H1("Preview your Sveltin project"))

	pathToPkgFile := filepath.Join(cfg.pathMaker.GetRootFolder(), "package.json")
	npmClient, err := utils.RetrievePackageManagerFromPkgJSON(cfg.fs, pathToPkgFile)
	utils.ExitIfError(err)

	err = helpers.RunPMCommand(npmClient.Name, "preview", "", nil, false)
	utils.ExitIfError(err)
}

func init() {
	rootCmd.AddCommand(previewCmd)
}
