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

var prepareCmd = &cobra.Command{
	Use:     "prepare",
	Aliases: []string{"sync"},
	Short:   "It wraps svelte-kit sync command.",
	Long: resources.GetASCIIArt() + `
It wraps svelte-kit sync command to ensure types are set up and correct before run typechecking.`,
	Run: RunPrepareCmd,
}

// RunPrepareCmd is the actual work function.
func RunPrepareCmd(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	log.Plain(utils.Underline("Running svelte-kit sync command"))

	pathToPkgFile := filepath.Join(pathMaker.GetRootFolder(), "package.json")
	npmClient, err := utils.RetrievePackageManagerFromPkgJSON(AppFs, pathToPkgFile)
	utils.ExitIfError(err)

	err = helpers.RunPMCommand(npmClient.Name, "prepare", "", nil, false)
	utils.ExitIfError(err)
	log.Success("Done")
}

func init() {
	rootCmd.AddCommand(prepareCmd)
}
